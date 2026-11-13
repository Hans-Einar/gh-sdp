package main

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// Exercise the real gh binary-release install/upgrade implementation against an
// isolated TLS GitHub API fixture. This is not a local-extension symlink test.
func TestActualGHExtensionInstallUpgrade(t *testing.T) { testActualGH(t, false) }

func TestActualGHInterruptedUpdateRecovery(t *testing.T) { testActualGH(t, true) }

func testActualGH(t *testing.T, truncated bool) {
	gh, err := exec.LookPath("gh")
	if err != nil {
		t.Skip("gh required")
	}
	oldEngine, newEngine := os.Getenv("SDPTOOL_OLD_BINARY"), os.Getenv("GH_SDP_BINARY")
	if oldEngine == "" || newEngine == "" {
		t.Skip("set SDPTOOL_OLD_BINARY and GH_SDP_BINARY for canonical engine upgrade gate")
	}
	oldBytes, err := os.ReadFile(oldEngine)
	if err != nil {
		t.Fatal(err)
	}
	newBytes, err := os.ReadFile(newEngine)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(oldBytes, newBytes) {
		t.Fatal("upgrade needs distinct engines")
	}
	dir := t.TempDir()
	project := filepath.Join(dir, "project")
	if err = os.MkdirAll(filepath.Join(project, "SDP"), 0700); err != nil {
		t.Fatal(err)
	}
	original := []byte("project-owned instructions\n")
	if err = os.WriteFile(filepath.Join(project, "AGENTS.md"), original, 0600); err != nil {
		t.Fatal(err)
	}
	var latest atomic.Int32
	var failDownload atomic.Int32
	latest.Store(1)
	var server *httptest.Server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("gh API %s %s latest=%d", r.Method, r.URL.Path, latest.Load())
		switch {
		case strings.HasSuffix(r.URL.Path, "/repos/test/gh-sdp"):
			json.NewEncoder(w).Encode(map[string]any{"name": "gh-sdp", "full_name": "test/gh-sdp", "owner": map[string]any{"login": "test"}, "default_branch": "main"})
		case strings.Contains(r.URL.Path, "/releases/"):
			n := latest.Load()
			if strings.HasSuffix(r.URL.Path, "/tags/v2.2.1") {
				n = 1
			}
			assetBytes := oldBytes
			tag := "v2.2.1"
			if n == 2 {
				tag = "v2.3.0"
				assetBytes = newBytes
			}
			json.NewEncoder(w).Encode(map[string]any{"tag_name": tag, "assets": []any{map[string]any{"name": "gh-sdp-linux-amd64", "digest": "sha256:" + digest(assetBytes), "size": len(assetBytes), "url": fmt.Sprintf("%s/assets/%d", server.URL, n), "browser_download_url": fmt.Sprintf("%s/assets/%d", server.URL, n)}}})
		case r.URL.Path == "/assets/1":
			w.Write(oldBytes)
		case r.URL.Path == "/assets/2":
			if failDownload.Load() == 1 {
				http.Error(w, "injected release download failure", http.StatusServiceUnavailable)
				return
			}
			if failDownload.Load() == 2 {
				w.Header().Set("Content-Length", fmt.Sprint(len(newBytes)+1))
				w.Write(newBytes[:len(newBytes)/2])
				return
			}
			w.Write(newBytes)
		default:
			t.Logf("unhandled gh API: %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
		}
	})
	server = httptest.NewTLSServer(handler)
	defer server.Close()
	ca := filepath.Join(dir, "ca.pem")
	os.WriteFile(ca, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0600)
	// gh normalizes enterprise API hosts to port443. A disposable CONNECT
	// proxy routes its requests to our unprivileged TLS server.
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodConnect {
			http.Error(w, "CONNECT required", 405)
			return
		}
		upstream, e := net.Dial("tcp", strings.TrimPrefix(server.URL, "https://"))
		if e != nil {
			http.Error(w, e.Error(), 502)
			return
		}
		client, _, e := w.(http.Hijacker).Hijack()
		if e != nil {
			upstream.Close()
			return
		}
		fmt.Fprint(client, "HTTP/1.1 200 Connection Established\r\n\r\n")
		go func() { defer client.Close(); defer upstream.Close(); io.Copy(upstream, client) }()
		go func() { defer client.Close(); defer upstream.Close(); io.Copy(client, upstream) }()
	}))
	defer proxy.Close()
	host := "sdp.example.com"
	baseEnv := []string{"HOME=" + dir, "PATH=" + filepath.Join(dir, ".local", "bin") + ":" + os.Getenv("PATH"), "GH_CONFIG_DIR=" + filepath.Join(dir, "gh-config"), "XDG_DATA_HOME=" + filepath.Join(dir, "data"), "XDG_CACHE_HOME=" + filepath.Join(dir, "cache"), "GH_HOST=" + host, "GH_ENTERPRISE_TOKEN=fixture-only", "GH_TOKEN=fixture-only", "SSL_CERT_FILE=" + ca, "HTTPS_PROXY=" + proxy.URL, "GH_PROMPT_DISABLED=1", "GH_NO_UPDATE_NOTIFIER=1", "GH_NO_EXTENSION_UPDATE_NOTIFIER=1"}
	call := func(args ...string) []byte {
		t.Helper()
		cmd := exec.Command(gh, args...)
		cmd.Env = baseEnv
		cmd.Dir = project
		out, e := cmd.CombinedOutput()
		if e != nil {
			t.Fatalf("gh %v: %v\n%s", args, e, out)
		}
		return out
	}
	call("extension", "install", "https://"+host+"/test/gh-sdp")
	// Setup is the only initial registration operation. Existing gh binary
	// extension upgrades replace the same executable pathname.
	call("sdp", "setup", "--json")
	alias := filepath.Join(dir, ".local", "bin", "sdptool")
	target, err := os.Readlink(alias)
	if err != nil {
		t.Fatal(err)
	}
	latest.Store(2)
	// Expire gh's HTTP metadata cache to model a later update check.
	if err := os.RemoveAll(filepath.Join(dir, "cache", "gh")); err != nil {
		t.Fatal(err)
	}

	failDownload.Store(1)
	if truncated {
		failDownload.Store(2)
	}
	interrupted := exec.Command(gh, "extension", "upgrade", "sdp")
	interrupted.Env = baseEnv
	interrupted.Dir = project

	interruptionOutput, interruptionErr := interrupted.CombinedOutput()
	retained, e := os.ReadFile(alias)
	if e != nil {
		t.Fatal(e)
	}
	if truncated {
		if bytes.Equal(retained, newBytes) {
			t.Fatal("truncation fixture did not corrupt payload")
		}
		// gh2.97.0 currently reports success and commits truncated bytes, even when
		// asset metadata advertises digest/size. Retain this observable limitation.
		t.Logf("truncated gh update result: error=%v output=%s installedBytes=%d canonicalBytes=%d", interruptionErr, interruptionOutput, len(retained), len(newBytes))
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		broken := exec.CommandContext(ctx, alias, "--version", "--json")
		broken.Env = []string{"HOME=" + dir, "PATH=/nonexistent", "SDP_OFFLINE=true"}
		broken.Dir = project
		if out, e := broken.CombinedOutput(); e == nil {
			t.Fatalf("partial engine did not fail closed: %s", out)
		}
	} else {
		if interruptionErr == nil {
			t.Fatalf("HTTP503 download unexpectedly succeeded: %s", interruptionOutput)
		}
		if !bytes.Equal(retained, oldBytes) {
			t.Fatal("HTTP failure lost prior registered engine")
		}
	}
	failDownload.Store(0)
	if truncated {
		t.Logf("forced upgrade attempt: %s", call("extension", "upgrade", "sdp", "--force"))
		t.Logf("forced install attempt: %s", call("extension", "install", "https://"+host+"/test/gh-sdp", "--force"))
		call("extension", "remove", "sdp")
		if err := os.RemoveAll(filepath.Join(dir, "cache", "gh")); err != nil {
			t.Fatal(err)
		}
		call("extension", "install", "https://"+host+"/test/gh-sdp")
	} else {
		t.Logf("upgrade: %s", call("extension", "upgrade", "sdp"))
	}

	// Deliberately no gh sdp invocation between update and direct execution.
	installed, err := os.ReadFile(alias)
	if err != nil || !bytes.Equal(installed, newBytes) {
		t.Fatalf("registered alias did not see updated canonical bytes: %v", err)
	}
	next, err := os.Readlink(alias)
	if err != nil || next != target {
		t.Fatalf("registration target changed: %q -> %q (%v)", target, next, err)
	}
	// Remove auth, gh PATH and server/network availability for direct operations.
	server.Close()
	direct := exec.Command(alias, "--version", "--json")
	direct.Env = []string{"HOME=" + dir, "PATH=/nonexistent", "XDG_CACHE_HOME=" + filepath.Join(dir, "empty-cache"), "SDP_OFFLINE=true"}
	direct.Dir = project
	output, err := direct.CombinedOutput()
	if err != nil {
		t.Fatalf("offline direct version: %v %s", err, output)
	}
	reference := exec.Command(newEngine, "--version", "--json")
	expected, err := reference.Output()
	if err != nil || !bytes.Equal(output, expected) {
		t.Fatalf("offline identity differs: %v %s != %s", err, output, expected)
	}
	discover := exec.Command(alias, "discover", "--json")
	discover.Env = direct.Env
	discover.Dir = project
	if out, err := discover.CombinedOutput(); err != nil {
		t.Fatalf("offline direct discovery: %v %s", err, out)
	}
	after, err := os.ReadFile(filepath.Join(project, "AGENTS.md"))
	if err != nil || !bytes.Equal(after, original) {
		t.Fatal("tool update changed project instructions")
	}
	entries, err := os.ReadDir(filepath.Join(project, "SDP"))
	if err != nil || len(entries) != 0 {
		t.Fatal("tool update/discovery mutated project SDP")
	}

	if _, err := os.Stat(filepath.Join(dir, "empty-cache", "sdptool", "distribution")); !os.IsNotExist(err) {
		t.Fatal("runtime populated legacy engine cache")
	}
}

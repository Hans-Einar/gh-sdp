package main

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/Hans-Einar/SDP/SDPTool/bootstrap"
)

func digest(data []byte) string { h := sha256.Sum256(data); return hex.EncodeToString(h[:]) }
func writeFixture(t *testing.T, path string, data []byte, mode os.FileMode) {
	t.Helper()
	if err := os.WriteFile(path, data, mode); err != nil {
		t.Fatal(err)
	}
}
func jsonFixture(t *testing.T, path string, value any) []byte {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	data = append(data, '\n')
	writeFixture(t, path, data, 0600)
	return data
}

// Set both paths to opt into real packaged-candidate integration. The test never
// builds, downloads or substitutes its own engine and uses disposable roots only.
func TestPackagedCandidates(t *testing.T) {
	client, engine := os.Getenv("GH_SDP_BINARY"), os.Getenv("SDPTOOL_BINARY")
	if client == "" || engine == "" {
		t.Skip("set GH_SDP_BINARY and SDPTOOL_BINARY to verify packaged candidates")
	}
	client, _ = filepath.Abs(client)
	engine, _ = filepath.Abs(engine)
	binary, err := os.ReadFile(engine)
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	root := filepath.Join(dir, "project with spaces")
	if err = os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	original := []byte("Project-owned instructions\n")
	managed := []byte("Managed test instructions\n")
	writeFixture(t, filepath.Join(root, "AGENTS.md"), original, 0600)
	if err = os.Mkdir(filepath.Join(root, "SDP"), 0700); err != nil {
		t.Fatal(err)
	}
	artifact := filepath.Join(dir, "sdptool")
	writeFixture(t, artifact, binary, 0700)
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	key := filepath.Join(dir, "test-key.pub")
	writeFixture(t, key, []byte(base64.StdEncoding.EncodeToString(pub)+"\n"), 0600)
	descriptor := filepath.Join(dir, "release.json")
	release := map[string]any{"schemaVersion": "sdp-release-descriptor/1", "release": "dev-gip-client", "sourceCommit": strings.Repeat("a", 40), "protocol": bootstrap.Protocol, "processProfile": "sdp-five-phase/0.1", "managementProfile": "sdp-project-management/0.2", "capabilities": []string{bootstrap.Protocol}, "files": []any{map[string]any{"path": "AGENTS.md", "type": "file", "ownership": "managed", "sha256": digest(managed), "content": managed}}, "retired": []string{}, "upgradesFrom": []string{}, "binaries": []any{map[string]any{"platform": runtime.GOOS + "/" + runtime.GOARCH, "path": "sdptool", "sha256": digest(binary), "size": len(binary)}}}
	sign := func(path string, value any) []byte {
		data := jsonFixture(t, path, value)
		jsonFixture(t, path+".sig", map[string]any{"keyId": digest(pub), "signature": ed25519.Sign(priv, data)})
		return data
	}
	data := sign(descriptor, release)
	manifest := filepath.Join(dir, "adoption.json")
	jsonFixture(t, manifest, map[string]any{"schemaVersion": "sdp-adoption/1", "projectRoot": root, "baseline": "manual", "observedCommit": nil, "targetDigest": digest(data), "inventory": map[string]any{"SDP": map[string]any{"type": "directory", "sha256": nil}, "AGENTS.md": map[string]any{"type": "file", "sha256": digest(original)}}, "moves": []any{}, "refreshManaged": []string{"AGENTS.md"}, "allowReferenceWarnings": false})
	t.Setenv("SDP_RELEASE", descriptor)
	t.Setenv("SDP_TEST_KEY", key)
	t.Setenv("SDP_CACHE_DIR", filepath.Join(dir, "cache"))
	t.Setenv("SDP_OFFLINE", "false")
	// Optional real GitHub CLI routing uses an isolated extension configuration;
	// it does not install into or modify the user's extension directory.
	viaGH := os.Getenv("GH_SDP_VIA_GH") == "true"
	if viaGH {
		t.Setenv("GH_CONFIG_DIR", filepath.Join(dir, "gh-config"))
		extension := filepath.Join(dir, "gh-sdp")
		if err = os.Mkdir(extension, 0700); err != nil {
			t.Fatal(err)
		}
		clientBytes, readErr := os.ReadFile(client)
		if readErr != nil {
			t.Fatal(readErr)
		}
		name := "gh-sdp"
		if runtime.GOOS == "windows" {
			name += ".exe"
		}
		writeFixture(t, filepath.Join(extension, name), clientBytes, 0700)
		cmd := exec.Command("gh", "extension", "install", ".")
		cmd.Dir = extension
		if output, installErr := cmd.CombinedOutput(); installErr != nil {
			t.Fatalf("isolated gh install: %v %s", installErr, output)
		}
	}
	call := func(bin string, args ...string) ([]byte, []byte, int) {
		t.Helper()
		cmd := exec.Command(bin, args...)
		if viaGH && bin == client {
			cmd = exec.Command("gh", append([]string{"sdp"}, args...)...)
		}
		cmd.Dir = root
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		code := 0
		if err != nil {
			if e, ok := err.(*exec.ExitError); ok {
				code = e.ExitCode()
			} else {
				t.Fatal(err)
			}
		}
		return stdout.Bytes(), stderr.Bytes(), code
	}
	require := func(bin string, args ...string) []byte {
		t.Helper()
		out, diagnostic, code := call(bin, args...)
		if code != 0 {
			t.Fatalf("%s %v: exit %d\n%s\n%s", bin, args, code, out, diagnostic)
		}
		return out
	}
	clientPlan := filepath.Join(dir, "client-plan.json")
	directPlan := filepath.Join(dir, "direct-plan.json")
	wrapperOutput := require(client, "upgrade", "--manifest", manifest, "--plan-output", clientPlan, "--json")
	directOutput := require(engine, "upgrade", "--manifest", manifest, "--plan-output", directPlan, "--json")
	a, _ := os.ReadFile(clientPlan)
	b, _ := os.ReadFile(directPlan)
	if !bytes.Equal(a, b) || !bytes.Equal(wrapperOutput, directOutput) {
		t.Fatal("direct and delegated previews differ")
	}
	if b, _ := os.ReadFile(filepath.Join(root, "AGENTS.md")); !bytes.Equal(b, original) {
		t.Fatal("preview mutated project")
	}
	// Removing the source proves saved apply uses the verified cache. The plan
	// itself retains signed release facts and the engine performs saved-plan apply.
	if err = os.Remove(descriptor); err != nil {
		t.Fatal(err)
	}
	require(client, "upgrade", "--apply", clientPlan, "--json")
	if b, _ := os.ReadFile(filepath.Join(root, "AGENTS-project.md")); !bytes.Equal(b, original) {
		t.Fatal("delegated apply lost project instructions")
	}
	if b, _ := os.ReadFile(filepath.Join(root, "AGENTS.md")); !bytes.Equal(b, managed) {
		t.Fatal("delegated apply did not write managed instructions")
	}
	// The disposable root is restored exactly for applying the same root-bound
	// plan through the direct entry point. Administrative operation files are local.
	if err = os.RemoveAll(root); err != nil {
		t.Fatal(err)
	}
	if err = os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	writeFixture(t, filepath.Join(root, "AGENTS.md"), original, 0600)
	if err = os.Mkdir(filepath.Join(root, "SDP"), 0700); err != nil {
		t.Fatal(err)
	}
	require(engine, "upgrade", "--apply", directPlan, "--json")
	if b, _ := os.ReadFile(filepath.Join(root, "AGENTS-project.md")); !bytes.Equal(b, original) {
		t.Fatal("direct apply preservation differs")
	}
	if b, _ := os.ReadFile(filepath.Join(root, "AGENTS.md")); !bytes.Equal(b, managed) {
		t.Fatal("direct apply outcome differs")
	}
	writeFixture(t, descriptor, data, 0600)
	wrapperOutput = require(client, "upgrade", "--json")
	directOutput = require(engine, "upgrade", "--json")
	if !bytes.Equal(wrapperOutput, directOutput) {
		t.Fatal("repeat preview differs")
	}
	var envelope struct {
		Result struct {
			Actions  []json.RawMessage
			CanApply bool
		}
	}
	if err = json.Unmarshal(wrapperOutput, &envelope); err != nil || len(envelope.Result.Actions) != 0 {
		t.Fatalf("repeat not no-change: %s", wrapperOutput)
	}
	// A descriptor with an unsupported protocol must be rejected before the child.
	release["protocol"] = "sdp-install-command/999"
	bad := filepath.Join(dir, "incompatible.json")
	sign(bad, release)
	t.Setenv("SDP_RELEASE", bad)
	out, diagnostic, code := call(client, "upgrade", "--json")
	if code != 4 || len(out) != 0 || !bytes.Contains(diagnostic, []byte("protocol")) {
		t.Fatalf("incompatible: %d %s %s", code, out, diagnostic)
	}
	// A signed descriptor can name a digest-valid engine with the wrong probe
	// protocol. The client must reject that executable as well.
	incompatibleBinary := bytes.ReplaceAll(binary, []byte(bootstrap.Protocol), []byte("sdp-install-command/9"))
	if bytes.Equal(binary, incompatibleBinary) {
		t.Fatal("engine protocol marker absent")
	}
	writeFixture(t, filepath.Join(dir, "incompatible-engine"), incompatibleBinary, 0700)
	release["protocol"] = bootstrap.Protocol
	release["binaries"] = []any{map[string]any{"platform": runtime.GOOS + "/" + runtime.GOARCH, "path": "incompatible-engine", "sha256": digest(incompatibleBinary), "size": len(incompatibleBinary)}}
	badEngineDescriptor := filepath.Join(dir, "bad-engine.json")
	sign(badEngineDescriptor, release)
	t.Setenv("SDP_RELEASE", badEngineDescriptor)
	out, diagnostic, code = call(client, "--version")
	if code != 4 || len(out) != 0 || !bytes.Contains(diagnostic, []byte("lacks")) {
		t.Fatalf("incompatible binary: %d %s %s", code, out, diagnostic)
	}
	t.Setenv("SDP_RELEASE", descriptor)
	// Cached binary corruption must fail, including offline invocation.
	cacheBinary := filepath.Join(os.Getenv("SDP_CACHE_DIR"), digest(binary)+"-sdptool")
	if runtime.GOOS == "windows" {
		cacheBinary += ".exe"
	}
	writeFixture(t, cacheBinary, []byte("corrupt"), 0700)
	t.Setenv("SDP_OFFLINE", "true")
	out, diagnostic, code = call(client, "--version")
	if code != 4 || len(out) != 0 || !bytes.Contains(diagnostic, []byte("digest/size")) {
		t.Fatalf("corrupt: %d %s %s", code, out, diagnostic)
	}
	t.Logf("packaged SHA256 gh-sdp=%s sdptool=%s; preview equality, apply, preservation, repetition, incompatible and corrupt cache passed", fileDigest(t, client), digest(binary))
}
func fileDigest(t *testing.T, path string) string {
	t.Helper()
	b, e := os.ReadFile(path)
	if e != nil {
		t.Fatal(e)
	}
	return digest(b)
}

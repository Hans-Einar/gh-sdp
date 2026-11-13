// package is a build-time verifier/republisher, never an installed runtime.
package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/Hans-Einar/SDP/SDPTool/bootstrap"
)

func digest(b []byte) string { h := sha256.Sum256(b); return hex.EncodeToString(h[:]) }

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: package VERSION DESCRIPTOR OUTPUT DISTRIBUTION-COMMIT")
		os.Exit(2)
	}
	if err := pack(context.Background(), os.Args[1], os.Args[2], os.Args[3], os.Args[4], os.Getenv("SDP_TEST_KEY")); err != nil {
		fmt.Fprintln(os.Stderr, "package:", err)
		os.Exit(2)
	}
}

func pack(ctx context.Context, version, descriptor, output, commit, testKey string) error {
	if runtime.GOOS != "linux" || runtime.GOARCH != "amd64" {
		return fmt.Errorf("qualification supports Linux amd64 only")
	}
	if !regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`).MatchString(version) || !regexp.MustCompile(`^[a-f0-9]{40}$`).MatchString(commit) {
		return fmt.Errorf("invalid distribution identity")
	}
	cache, err := os.MkdirTemp("", "gh-sdp-package-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(cache)
	cfg := bootstrap.Config{Descriptor: descriptor, TestKey: testKey, CacheDir: cache}
	// Bootstrap is reused ONLY at build time: signature, bounded asset, digest and
	// protocol verification remain canonical. No launcher or cache is distributed.
	binary, err := bootstrap.Bootstrap(ctx, cfg)
	if err != nil {
		return err
	}
	cfg.Offline = true
	verified, err := bootstrap.Resolve(ctx, cfg)
	if err != nil {
		return err
	}
	var release struct {
		Release      string `json:"release"`
		SourceCommit string `json:"sourceCommit"`
	}
	if err = json.Unmarshal(verified.Bytes, &release); err != nil {
		return err
	}
	if release.Release != version || !regexp.MustCompile(`^[a-f0-9]{40}$`).MatchString(release.SourceCommit) {
		return fmt.Errorf("signed canonical identity does not match selected version")
	}
	probe, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	data, err := exec.CommandContext(probe, binary, "--version", "--json").Output()
	if err != nil {
		return fmt.Errorf("engine identity: %w", err)
	}
	var identity struct {
		Version  string `json:"version"`
		Revision string `json:"revision"`
	}
	if err = json.Unmarshal(data, &identity); err != nil {
		return err
	}
	if identity.Version != version || identity.Revision != release.SourceCommit {
		return fmt.Errorf("engine identity differs from signed canonical source")
	}
	bytes, err := os.ReadFile(binary)
	if err != nil {
		return err
	}
	const name = "gh-sdp-linux-amd64"
	manifest := map[string]any{
		"schemaVersion": "gh-sdp-package/2", "version": version, "platform": "linux/amd64", "file": name, "sha256": digest(bytes), "size": len(bytes),
		"distributionRepository": "https://github.com/Hans-Einar/gh-sdp", "distributionCommit": commit,
		"engineRepository": "https://github.com/Hans-Einar/SDP", "engineCommit": release.SourceCommit, "engineVersion": identity.Version,
		"descriptor": descriptor, "descriptorSHA256": verified.SHA256, "signingKeyId": verified.KeyID, "provenance": verified.Provenance,
	}
	mb, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}
	// Preserve the signed descriptor exactly and its verification envelope. A test
	// signer remains visibly test-signed and cannot become production provenance.
	verified.Path = "" // The ephemeral build cache is not an installed source.
	envelope, err := json.Marshal(verified)
	if err != nil {
		return err
	}
	signer := verified.KeyID
	if strings.HasPrefix(signer, "test:") {
		pub, e := base64.StdEncoding.DecodeString(strings.TrimPrefix(signer, "test:"))
		if e != nil {
			return e
		}
		signer = digest(pub)
	}
	signature, _ := json.Marshal(map[string]any{"keyId": signer, "signature": verified.Signature})
	if err = os.Mkdir(output, 0755); err != nil {
		return err
	}
	success := false
	defer func() {
		if !success {
			os.RemoveAll(output)
		}
	}()
	for n, b := range map[string][]byte{name: bytes, "gh-sdp.manifest.json": append(mb, '\n'), "checksums.txt": []byte(digest(bytes) + "  " + name + "\n"), "sdp-release.json": verified.Bytes, "sdp-release.json.sig": append(signature, '\n'), "sdp-release.verification.json": append(envelope, '\n')} {
		mode := os.FileMode(0644)
		if n == name {
			mode = 0755
		}
		if err = os.WriteFile(filepath.Join(output, n), b, mode); err != nil {
			return err
		}
	}
	success = true
	return nil
}

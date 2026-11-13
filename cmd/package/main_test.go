package main

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func fixture(t *testing.T) (string, string, []byte) {
	return fixtureIdentity(t, strings.Repeat("a", 40))
}

func fixtureIdentity(t *testing.T, source string) (string, string, []byte) {
	t.Helper()
	dir := t.TempDir()
	engine := []byte("#!/bin/sh\nprintf '%s\\n' '{\"version\":\"2.3.0\",\"revision\":\"" + strings.Repeat("a", 40) + "\",\"installationProtocol\":\"sdp-install-command/1\"}'\n")
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	key := filepath.Join(dir, "test.pub")
	os.WriteFile(key, []byte(base64.StdEncoding.EncodeToString(pub)), 0600)
	os.WriteFile(filepath.Join(dir, "sdptool"), engine, 0700)
	release := map[string]any{"schemaVersion": "sdp-release-descriptor/1", "protocol": "sdp-install-command/1", "release": "2.3.0", "sourceCommit": source, "binaries": []any{map[string]any{"platform": "linux/amd64", "path": "sdptool", "sha256": digest(engine), "size": len(engine)}}}
	data, _ := json.Marshal(release)
	path := filepath.Join(dir, "release.json")
	os.WriteFile(path, data, 0600)
	sig, _ := json.Marshal(map[string]any{"keyId": digest(pub), "signature": ed25519.Sign(priv, data)})
	os.WriteFile(path+".sig", sig, 0600)
	return path, key, engine
}

func TestRepublishesExactSignedEngine(t *testing.T) {
	descriptor, key, engine := fixture(t)
	output := filepath.Join(t.TempDir(), "package")
	if err := pack(context.Background(), "2.3.0", descriptor, output, strings.Repeat("b", 40), key); err != nil {
		t.Fatal(err)
	}
	actual, err := os.ReadFile(filepath.Join(output, "gh-sdp-linux-amd64"))
	if err != nil || !bytes.Equal(actual, engine) {
		t.Fatalf("changed engine bytes: %v", err)
	}
	data, _ := os.ReadFile(filepath.Join(output, "gh-sdp.manifest.json"))
	var m map[string]any
	json.Unmarshal(data, &m)
	if m["engineCommit"] != strings.Repeat("a", 40) || m["distributionCommit"] != strings.Repeat("b", 40) || m["provenance"] != "test-signed" || m["sha256"] != digest(engine) {
		t.Fatalf("incorrect provenance: %s", data)
	}
	signed, _ := os.ReadFile(descriptor)
	retained, _ := os.ReadFile(filepath.Join(output, "sdp-release.json"))
	if !bytes.Equal(signed, retained) {
		t.Fatal("changed signed descriptor")
	}
	if err = pack(context.Background(), "2.3.0", descriptor, output, strings.Repeat("b", 40), key); err == nil {
		t.Fatal("overwrote output")
	}
}

func TestRejectsUntrustedTamperedAndWrongIdentity(t *testing.T) {
	for _, kind := range []string{"untrusted", "descriptor", "asset", "version", "revision"} {
		t.Run(kind, func(t *testing.T) {
			descriptor, key, _ := fixture(t)
			version := "2.3.0"
			switch kind {
			case "untrusted":
				key = ""
			case "descriptor":
				f, _ := os.OpenFile(descriptor, os.O_APPEND|os.O_WRONLY, 0)
				f.WriteString(" ")
				f.Close()
			case "asset":
				os.WriteFile(filepath.Join(filepath.Dir(descriptor), "sdptool"), []byte("tampered"), 0700)
			case "version":
				version = "2.4.0"
			case "revision":
				descriptor, key, _ = fixtureIdentity(t, strings.Repeat("c", 40))
			}
			output := filepath.Join(t.TempDir(), "package")
			if err := pack(context.Background(), version, descriptor, output, strings.Repeat("b", 40), key); err == nil {
				t.Fatal("accepted invalid distribution")
			}
			if _, err := os.Stat(output); !os.IsNotExist(err) {
				t.Fatal("left a partial package")
			}
		})
	}
}

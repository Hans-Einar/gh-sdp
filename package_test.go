package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Validate the produced release directory, not a proposed filename. GitHub CLI's
// extension manager discovers Linux amd64 assets with HasSuffix("linux-amd64").
// See cli/cli pkg/cmd/extension/manager.go, installBin.
func TestPackagedAssetDiscovery(t *testing.T) {
	dir := os.Getenv("GH_SDP_PACKAGE_DIR")
	if dir == "" {
		t.Skip("set GH_SDP_PACKAGE_DIR to verify the actual release assets")
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	var selected []string
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), "linux-amd64") {
			selected = append(selected, entry.Name())
			info, err := entry.Info()
			if err != nil || !info.Mode().IsRegular() || info.Mode()&0111 == 0 {
				t.Fatalf("selected asset is not a regular executable: %s (%v)", entry.Name(), err)
			}
		}
	}
	if len(selected) != 1 {
		t.Fatalf("GitHub CLI must discover exactly one linux-amd64 asset; found %v", selected)
	}
	manifestBytes, err := os.ReadFile(filepath.Join(dir, "gh-sdp.manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var manifest struct {
		File, Platform, SHA256 string
		Size                   int
	}
	if err = json.Unmarshal(manifestBytes, &manifest); err != nil {
		t.Fatal(err)
	}
	binary, err := os.ReadFile(filepath.Join(dir, selected[0]))
	if err != nil {
		t.Fatal(err)
	}
	hash := digest(binary)
	if manifest.File != selected[0] || manifest.Platform != "linux/amd64" || manifest.SHA256 != hash || manifest.Size != len(binary) {
		t.Fatalf("selected asset does not match package manifest: %+v", manifest)
	}
	checksums, err := os.ReadFile(filepath.Join(dir, "checksums.txt"))
	if err != nil || string(checksums) != hash+"  "+selected[0]+"\n" {
		t.Fatalf("selected asset does not match checksums: %s (%v)", checksums, err)
	}
}

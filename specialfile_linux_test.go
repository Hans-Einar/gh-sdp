package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestPackagedClientRejectsFIFO(t *testing.T) {
	client := os.Getenv("GH_SDP_BINARY")
	if client == "" {
		t.Skip("set GH_SDP_BINARY for actual FIFO rejection")
	}
	client, err := filepath.Abs(client)
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	fifo := filepath.Join(dir, "descriptor.fifo")
	if err = syscall.Mkfifo(fifo, 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("SDP_RELEASE", fifo)
	t.Setenv("SDP_CACHE_DIR", filepath.Join(dir, "cache"))
	t.Setenv("SDP_OFFLINE", "false")
	t.Setenv("SDP_TEST_KEY", "")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	output, err := exec.CommandContext(ctx, client, "--version").CombinedOutput()
	if ctx.Err() != nil {
		t.Fatalf("unsupported FIFO blocked: %v", ctx.Err())
	}
	child, ok := err.(*exec.ExitError)
	if !ok || child.ExitCode() != 4 || !strings.Contains(string(output), "regular") {
		t.Fatalf("FIFO: %v %s", err, output)
	}
}

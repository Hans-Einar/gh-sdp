// gh-sdp is a thin process adapter for the canonical SDPTool executable.
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/Hans-Einar/SDP/SDPTool/bootstrap"
)

func main() {
	os.Exit(run(context.Background(), os.Args[1:], os.Getenv, os.Stdin, os.Stdout, os.Stderr, bootstrap.Bootstrap))
}

type resolveBinary func(context.Context, bootstrap.Config) (string, error)

func configuration(args []string, getenv func(string) string) (bootstrap.Config, error) {
	c := bootstrap.Config{Descriptor: getenv("SDP_RELEASE"), TestKey: getenv("SDP_TEST_KEY"), CacheDir: getenv("SDP_CACHE_DIR")}
	// Keep accepted environment spelling identical to the child's canonical
	// offline switch so bootstrap and preview cannot disagree about networking.
	switch getenv("SDP_OFFLINE") {
	case "true":
		c.Offline = true
	case "", "false":
	default:
		return c, fmt.Errorf("SDP_OFFLINE must be true or false")
	}
	// Saved operations never update the selected distribution during execution.
	for _, arg := range args {
		flag, _, _ := strings.Cut(arg, "=")
		if flag == "--apply" || flag == "-apply" || flag == "--resume" || flag == "-resume" {
			c.Offline = true
		}
	}
	return c, nil
}

func run(ctx context.Context, args []string, getenv func(string) string, stdin io.Reader, stdout, stderr io.Writer, resolve resolveBinary) int {
	c, err := configuration(args, getenv)
	if err != nil {
		fmt.Fprintln(stderr, "gh-sdp:", err)
		return 2
	}
	binary, err := resolve(ctx, c)
	if err != nil {
		fmt.Fprintln(stderr, "gh-sdp:", err)
		return 4
	}
	cmd := exec.CommandContext(ctx, binary, args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = stdin, stdout, stderr
	if err := cmd.Run(); err != nil {
		var child *exec.ExitError
		if errors.As(err, &child) {
			if code := child.ExitCode(); code >= 0 {
				return code
			}
		}
		fmt.Fprintln(stderr, "gh-sdp:", err)
		return 4
	}
	return 0
}

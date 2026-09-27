package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/Hans-Einar/SDP/SDPTool/bootstrap"
)

// The test executable provides a real child process for stream/argv tests.
func TestChildProcess(t *testing.T) {
	if os.Getenv("GH_SDP_CHILD_TEST") != "1" {
		return
	}
	cwd, _ := os.Getwd()
	input, _ := io.ReadAll(os.Stdin)
	data := struct {
		Args            []string
		Input, Cwd, Env string
	}{os.Args[3:], string(input), cwd, os.Getenv("GH_SDP_INHERITED")}
	_ = json.NewEncoder(os.Stdout).Encode(data)
	fmt.Fprint(os.Stderr, "child diagnostic\n")
	os.Exit(23)
}

func TestDelegatesExactArgumentsStreamsEnvironmentAndExit(t *testing.T) {
	t.Setenv("GH_SDP_CHILD_TEST", "1")
	t.Setenv("GH_SDP_INHERITED", "inherited value")
	executable, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-test.run=TestChildProcess", "--", "upgrade", "--manifest", "path with spaces;$(touch NO)", "", "--literal=*"}
	var stdout, stderr bytes.Buffer
	called := false
	code := run(context.Background(), args, func(string) string { return "" }, bytes.NewBufferString("stdin bytes\x00\n"), &stdout, &stderr,
		func(context.Context, bootstrap.Config) (string, error) { called = true; return executable, nil })
	if !called || code != 23 {
		t.Fatalf("called=%t exit=%d stderr=%s", called, code, &stderr)
	}
	var got struct {
		Args            []string
		Input, Cwd, Env string
	}
	if err := json.Unmarshal(stdout.Bytes(), &got); err != nil {
		t.Fatalf("%v: %s", err, &stdout)
	}
	cwd, _ := os.Getwd()
	if !reflect.DeepEqual(got.Args, args[2:]) || got.Input != "stdin bytes\x00\n" || got.Cwd != cwd || got.Env != "inherited value" {
		t.Fatalf("child observed %+v", got)
	}
	if stderr.String() != "child diagnostic\n" {
		t.Fatalf("stderr %q", &stderr)
	}
}

func TestConfigurationAndFailureBoundaries(t *testing.T) {
	for _, value := range []string{"1", "t", "TRUE"} {
		if _, err := configuration(nil, func(key string) string {
			if key == "SDP_OFFLINE" {
				return value
			}
			return ""
		}); err == nil {
			t.Fatalf("accepted offline spelling ignored by child: %q", value)
		}
	}
	c, err := configuration(nil, func(key string) string {
		if key == "SDP_OFFLINE" {
			return "true"
		}
		return ""
	})
	if err != nil || !c.Offline {
		t.Fatalf("offline true: %+v %v", c, err)
	}
	for _, args := range [][]string{{"upgrade", "--apply", "saved plan"}, {"upgrade", "--resume=operation"}, {"install", "-apply=x"}} {
		c, err := configuration(args, func(k string) string {
			return map[string]string{"SDP_RELEASE": "descriptor", "SDP_TEST_KEY": "test-key", "SDP_CACHE_DIR": "cache", "SDP_OFFLINE": "false"}[k]
		})
		if err != nil || !c.Offline || c.Descriptor != "descriptor" || c.TestKey != "test-key" || c.CacheDir != "cache" {
			t.Fatalf("config %+v: %v", c, err)
		}
	}
	for _, test := range []struct {
		name, offline string
		resolver      resolveBinary
		want          int
	}{
		{"invalid config", "garbage", func(context.Context, bootstrap.Config) (string, error) {
			t.Fatal("resolved after invalid configuration")
			return "", nil
		}, 2},
		{"verification failed", "", func(context.Context, bootstrap.Config) (string, error) { return "", errors.New("invalid signature") }, 4},
		{"start failed", "", func(context.Context, bootstrap.Config) (string, error) {
			return filepath.Join(t.TempDir(), "absent"), nil
		}, 4},
	} {
		t.Run(test.name, func(t *testing.T) {
			var out, diagnostic bytes.Buffer
			code := run(context.Background(), nil, func(k string) string {
				if k == "SDP_OFFLINE" {
					return test.offline
				}
				return ""
			}, nil, &out, &diagnostic, test.resolver)
			if code != test.want || out.Len() != 0 || diagnostic.Len() == 0 {
				t.Fatalf("exit=%d out=%q err=%q", code, &out, &diagnostic)
			}
		})
	}
}

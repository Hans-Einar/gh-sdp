package main

import (
	"encoding/base64"
	"github.com/Hans-Einar/SDP/SDPTool/bootstrap"
	"testing"
)

// This public non-release proof exercises the actual pinned dependency.
func TestBuildTimeLinuxPublisherTrust(t *testing.T) {
	data := []byte("SDP publisher trust-anchor probe: Linux build host, 2026-11-11. Not a release descriptor.\n")
	sig, err := base64.StdEncoding.DecodeString("adARfG5LbAdInSYkB7oC7n3W6xGVeib2sQlUyc6QeHYVRLkxMeAqk/4uqLrA5uMwXQ5UAIiK21cBJ627MExMAA==")
	if err != nil {
		t.Fatal(err)
	}
	const id = "2c947457cffe015f870168cb332242c49606fdea378a13d875579544c2989c03"
	if err := bootstrap.Verify(data, sig, id, "signed"); err != nil {
		t.Fatal(err)
	}
	if bootstrap.Verify(append(data, '!'), sig, id, "signed") == nil {
		t.Fatal("tampered proof accepted")
	}
}

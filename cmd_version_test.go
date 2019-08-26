package main

import (
	"bytes"
	"context"
	"runtime"
	"testing"
)

func TestVersion(t *testing.T) {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	//
	// Expected
	//
	expected := "unreleased\n"

	s := versionCmd{}
	s.Execute(context.TODO(), nil)
	if out.(*bytes.Buffer).String() != expected {
		t.Errorf("Expected '%s' received '%s'", expected, out)
	}
}

func TestVersionVerbose(t *testing.T) {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	//
	// Expected
	//
	expected := "unreleased\nBuilt with " + runtime.Version() + "\n"

	s := versionCmd{verbose: true}
	s.Execute(context.TODO(), nil)
	if out.(*bytes.Buffer).String() != expected {
		t.Errorf("Expected '%s' received '%s'", expected, out)
	}
}

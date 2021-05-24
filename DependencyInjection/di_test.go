package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Manuel")

	got := buffer.String()
	want := "Hello, Manuel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

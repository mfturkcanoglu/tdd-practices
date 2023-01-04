package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Mf")

	got := buffer.String()
	want := "Hello, Mf"

	if got != want {
		t.Errorf("Expected %q but got %q", want, got)
	}
}

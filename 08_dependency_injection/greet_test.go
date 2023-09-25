package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Randson")

	got := buffer.String()
	want := "Hello, Randson!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

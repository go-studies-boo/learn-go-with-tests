package injection

import (
	"bytes"
	"testing"
)

func TestGret(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
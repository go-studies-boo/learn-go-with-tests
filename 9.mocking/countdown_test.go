package mocking

import (
	"bytes"
	"testing"
)

// Spies are a kind of mock which can record how a dependency is used.
// They can record the arguments sent in, how many times it has been called, etc.
// In our case, we're keeping track of how many times Sleep() is called so we can check it in our test.
type SpySlepeer struct {
	Calls int
}

func (s *SpySlepeer) Sleep() {
	s.Calls++
}

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySlepeer{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}

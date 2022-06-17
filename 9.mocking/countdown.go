package mocking

import (
	"io"
	"fmt"
)

// - In main we will send to os.Stdout so our users see the countdown printed to the terminal.
// - In test we will send to bytes.Buffer so our tests can capture what data is being generated.

// out *bytes.Buffer works but, it would be better to use a general purpose interface instead
func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
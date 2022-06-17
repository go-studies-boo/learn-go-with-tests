package mocking

import (
	"io"
	"fmt"
)

const finalWord = "Go!"
const countdownStart = 3

// Let's define our dependency as an interface. 
// This lets us then use a real Sleeper in main and a spy sleeper in our tests.
type Sleeper interface {
	Sleep()
}

// - In main we will send to os.Stdout so our users see the countdown printed to the terminal.
// - In test we will send to bytes.Buffer so our tests can capture what data is being generated.

// out *bytes.Buffer works but, it would be better to use a general purpose interface instead
func Countdown(out io.Writer,  sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

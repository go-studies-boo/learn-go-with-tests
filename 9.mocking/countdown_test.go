package mocking

import (
	"time"
	"reflect"
	"bytes"
	"testing"
)

const write = "write"
const sleep = "sleep"

// Spies are a kind of mock which can record how a dependency is used.
// They can record the arguments sent in, how many times it has been called, etc.
// In our case, we're keeping track of how many times Sleep() is called so we can check it in our test.
type SpySlepeer struct {
	Calls int
}

type SpyCountDownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpySlepeer) Sleep() {
	s.Calls++
}

func (s *SpyTime) Sleep (duration time.Duration) {
	s.durationSlept = duration
}

func TestCountDown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
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
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountDownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

// You may have heard mocking is evil. Just like anything in software development it can be used for evil, just like DRY.
// People normally get in to a bad state when they don't listen to their tests and are not respecting the refactoring stage.
//
// If your mocking code is becoming complicated or you are having to mock out lots of things to test something, you should listen to that bad feeling and think about your code. Usually it is a sign of
//
// - The thing you are testing is having to do too many things (because it has too many dependencies to mock)
//     - Break the module apart so it does less
// - Its dependencies are too fine-grained
//     - Think about how you can consolidate some of these dependencies into one meaningful module
// - Your test is too concerned with implementation details
//     - Favour testing expected behaviour rather than the implementation
//
// What people see here is a weakness in TDD but it is actually a strength, more often than not poor test code is a 
// result of bad design or put more nicely, well-designed code is easy to test.

// ---
//
// Ever run into this situation?
//
// - You want to do some refactoring
// - To do this you end up changing lots of tests
// - You question TDD and make a post on Medium titled "Mocking considered harmful"

// This is usually a sign of you testing too much implementation detail. Try to make it so your tests 
// are testing useful behaviour unless the implementation is really important to how the system runs.

// --
// It is sometimes hard to know what level to test exactly but here are some thought processes and rules I try to follow:
//
// - The definition of refactoring is that the code changes but the behaviour stays the same. 
//   If you have decided to do some refactoring in theory you should be able to make the commit without any test changes. 
//   So when writing a test ask yourself
//     - Am I testing the behaviour I want, or the implementation details?
//     - If I were to refactor this code, would I have to make lots of changes to the tests?
//
// - Although Go lets you test private functions, I would avoid it as private functions are implementation detail to 
//   support public behaviour. Test the public behaviour. Sandi Metz describes private functions as being "less stable" 
//   and you don't want to couple your tests to them.
//
// - I feel like if a test is working with more than 3 mocks then it is a red flag - time for a rethink on the design
//
// - Use spies with caution. Spies let you see the insides of the algorithm you are writing which can be very useful 
//   but that means a tighter coupling between your test code and the implementation. Be sure you actually care about 
//   these details if you're going to spy on them
// ---

// # Can't I just use a mocking framework?
//
// Mocking requires no magic and is relatively simple; using a framework can make mocking seem more 
// complicated than it is. We don't use automocking in this chapter so that we get:
//
// - a better understanding of how to mock
// - practice implementing interfaces
//
// You should only use a mock generator that generates test doubles against an interface. 
// Any tool that overly dictates how tests are written, or that use lots of 'magic', can get in the sea.

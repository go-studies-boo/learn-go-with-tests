package mocking

import (
	"time"
	"os"
)

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
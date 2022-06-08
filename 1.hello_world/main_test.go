package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// t.Run is a subtests
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Iago", "")
		want := "Hello, Iago"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' whenan empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("iago", "spanish")
		want := "Hola, iago"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("iago", "french")
		want := "Bonjour, iago"

		assertCorrectMessage(t, got, want)
	})
}
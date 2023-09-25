package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Randson", "")
		want := "Hello, Randson!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to Oops!", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("possibility to add French as well", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Randson", "French"), "Bonjour, Randson!")
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

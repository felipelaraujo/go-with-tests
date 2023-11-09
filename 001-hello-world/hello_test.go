package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Felipe", "")
		want := "Hello, Felipe!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Felipe", "Portuguese")
		want := "Olá, Felipe!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Felipe", "French")
		want := "Bonjour, Felipe!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ivan")
		want := "Hello, Ivan!"
		assertCorrectMsg(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Worldd!"
		assertCorrectMsg(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

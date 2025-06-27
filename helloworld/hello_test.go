package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ivan", "")
		want := "Hello, Ivan!"
		assertCorrectMsg(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMsg(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty name, but lang is Spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World!"
		assertCorrectMsg(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty name, but lang is French", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World!"
		assertCorrectMsg(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMsg(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Marry", "French")
		want := "Bonjour, Marry!"
		assertCorrectMsg(t, got, want)
	})
	t.Run("in English", func(t *testing.T) {
		got := Hello("Marry", "English")
		want := "Hello, Marry!"
		assertCorrectMsg(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

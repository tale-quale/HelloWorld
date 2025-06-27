package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ivan")
	want := "Hello, Ivan!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

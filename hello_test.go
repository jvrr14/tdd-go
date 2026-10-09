package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Javokhir")
	want := "Hello, Javokhir!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Andrew")
	want := "Hello, Andrew"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

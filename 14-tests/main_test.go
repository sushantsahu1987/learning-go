package main

import "testing"

func TestAdd1(t *testing.T) {
	got := Add(1, 2)
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestAdd2(t *testing.T) {
	got := Add(3, 5)
	want := 3

	if got == want {
		t.Errorf("got %d != want %d", got, want)
	}
}

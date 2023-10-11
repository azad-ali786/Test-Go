package main

import "testing"

	// func TestHello(t *testing.T) {
	// 	got := Hello()
	// 	want := "Hello, world"

	// 	if got != want {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}
	// }

func TestHello(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		got1 := Hello("")
		want1 := "Hello, World"
	
		if got != want {
			t.Errorf("got %q want %q", got1, want1)
		}
}
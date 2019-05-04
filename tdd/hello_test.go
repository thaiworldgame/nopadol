package main

import (
	"testing"
	"fmt"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		fmt.Println(got)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

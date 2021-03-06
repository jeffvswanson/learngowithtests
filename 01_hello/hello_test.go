package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jeff", "")
		want := "Hello, Jeff!"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("Spanish greeting", func(t *testing.T) {
		got := Hello("Manuel", "Espanol")
		want := "Hola, Manuel!"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("French greeting", func(t *testing.T) {
		got := Hello("Pierre", "Francais")
		want := "Bonjour, Pierre!"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("German greeting", func(t *testing.T) {
		got := Hello("Hans", "Deutsch")
		want := "Guten Tag, Hans!"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
}

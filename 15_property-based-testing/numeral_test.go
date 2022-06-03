package main

import "testing"

func TestRomanNumerals(t *testing.T) {
	got := ConvertToRoman(i)
	want := "I"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

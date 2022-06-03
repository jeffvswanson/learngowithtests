package main

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 -> I", 1, "I"},
		{"2 -> II", 2, "II"},
	}
	for _, tc := range cases {
		t.Run(tc.Description, func(t *testing.t) {
			got := ConvertToRoman(tc.Arabic)
			if got != tc.Want {
				t.Errorf("got %q, want %q", got, tc.Want)
			}
		})
	}
}

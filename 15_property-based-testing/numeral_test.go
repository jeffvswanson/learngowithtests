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
		{"3 -> III", 3, "III"},
		{"4 -> IV", 4, "IV"},
		{"5 -> V", 5, "V"},
		{"6 -> VI", 6, "VI"},
		{"7 -> VII", 7, "VII"},
		{"8 -> VIII", 8, "VIII"},
		{"9 -> IX", 9, "IX"},
		{"10 -> X", 10, "X"},
		{"14 -> XIV", 14, "XIV"},
		{"20 -> XX", 20, "XX"},
		{"39 -> XXXIX", 39, "XXXIX"},
	}
	for _, tc := range cases {
		t.Run(tc.Description, func(t *testing.T) {
			got := ConvertToRoman(tc.Arabic)
			if got != tc.Want {
				t.Errorf("got %q, want %q", got, tc.Want)
			}
		})
	}
}

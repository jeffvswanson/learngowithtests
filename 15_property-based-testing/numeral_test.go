package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 798, Roman: "DCCXCVIII"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
}

func TestConvertToRoman(t *testing.T) {
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d -> %q", tc.Arabic, tc.Roman), func(t *testing.T) {
			got := ConvertToRoman(tc.Arabic)
			if got != tc.Roman {
				t.Errorf("got %q, want %q", got, tc.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%q -> %d", tc.Roman, tc.Arabic), func(t *testing.T) {
			got := ConvertToArabic(tc.Roman)
			if got != tc.Arabic {
				t.Errorf("got %d, want %d", got, tc.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}
	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}

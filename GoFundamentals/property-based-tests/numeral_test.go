package numeral

import (
	"fmt"
	"testing"
	"testing/quick"
)

var (
	cases = []struct {
		Arabic uint16
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 798, Roman: "DCCXCVIII"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
	}
)

func TestRomanNumerals(t *testing.T) {

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", tt.Arabic, tt.Roman), func(t *testing.T) {
			got := ConvertToRoman(tt.Arabic)

			if got != tt.Roman {
				t.Errorf("got %s, want %s", got, tt.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", tt.Roman, tt.Arabic), func(t *testing.T) {
			got := ConvertToArabic(tt.Roman)

			if got != tt.Arabic {
				t.Errorf("got %d, want %d", got, tt.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {

	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}

		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		if fromRoman != arabic {
			return false
		}

		var count uint8
		var lastCharacter rune
		for _, c := range roman {
			if c != lastCharacter {
				count = 1
			} else {
				count++
			}
			if count > 3 {
				break
			}
			lastCharacter = c
		}
		return count <= 3

	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

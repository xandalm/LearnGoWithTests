package numeral

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Arabic int
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

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", tt.Arabic, tt.Roman), func(t *testing.T) {
			got := ConvertToRoman(tt.Arabic)

			if got != tt.Roman {
				t.Errorf("got %s, want %s", got, tt.Roman)
			}
		})
	}
}

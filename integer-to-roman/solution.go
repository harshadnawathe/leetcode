package integertoroman

import "strings"

var digitToSymbol = map[int][]string{
	1:    {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	10:   {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	100:  {"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	1000: {"", "M", "MM", "MMM"},
}

func intToRoman(num int) string {
	var b strings.Builder

	for n := 1000; n >= 1; n /= 10 {
		d := num / n
		num = num % n
		b.WriteString(digitToSymbol[n][d])
	}
	return b.String()
}

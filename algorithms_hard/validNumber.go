package main

import "strings"

/*
Validate if a given string can be interpreted as a decimal number.

Some examples:
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
" -90e3   " => true
" 1e" => false
"e3" => false
" 6e-1" => true
" 99e2.5 " => false
"53.5e93" => true
" --6 " => false
"-+3" => false
"95a54e53" => false
*/

type validNumberTestCase struct {
	number   string
	isNumber bool
}

func main() {
	cases := []validNumberTestCase{
		{"0", true},
		{" 0.1 ", true},
		{"abc", false},
		{"1 a", false},
		{"2e10", true},
		{" -90e3   ", true},
		{" 1e", false},
		{"e3", false},
		{" 6e-1", true},
		{" 99e2.5 ", false},
		{"53.5e93", true},
		{" --6 ", false},
		{"-+3", false},
		{"95a54e53", false},
	}
	for _, cs := range cases {
		println(cs.number, "=>", isNumber(cs.number), "(", cs.isNumber, ")")
	}
}

func isNumeric(s string, integer bool) bool {
	if s == "" || (integer && strings.Contains(s, ".")) {
		return false
	}

	hasDot := false
	hasDigit := false
	start := 0
	if s[0] == '-' || s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] == '.' {
			if hasDot {
				return false
			}

			hasDot = true

			continue
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		hasDigit = true
	}

	return hasDigit
}

func isNumber(s string) bool {
	cleaned := strings.TrimSpace(s)
	parts := strings.Split(cleaned, "e")
	l := len(parts)
	if l < 1 || l > 2 {
		return false
	}

	isN := isNumeric(parts[0], false)
	if l == 2 {
		isN = isN && isNumeric(parts[1], true)
	}

	return isN
}

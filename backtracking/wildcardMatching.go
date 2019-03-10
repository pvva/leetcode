package main

type WildcardTestCase struct {
	input   string
	pattern string
	match   bool
}

func main() {
	testCases := []WildcardTestCase{
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"adceb", "*a*b", true},
		{"acdcb", "a*c?b", false},
		{"acdcb", "a*c*b", true},
		{"a", "a", true},
		{"a", "?", true},
		{"a", "?a", false},
		{"aa", "?a", true},
		{"abefcdgiescdfimde", "ab*cd?i*de", true},
		{"aaaa", "***a", true},
		{"aaaa", "***aa", true},
		{"aaaa", "***aaa", true},
		{"", "*", true},
		{"", "**", true},
		{"bcd", "??", false},
		{"bcd", "*?*", true},
		{"abc", "*ab", false},
		{"abc", "*bc", true},
		{"hi", "*?", true},
		{"hi", "*??", true},
		{"b", "?", true},
		{"b", "?*?", false},
		{"mississippi", "m*iss*iss*", true},
		{"mississippi", "m*si*", true},
	}

	for i, tc := range testCases {
		println(i, isMatch(tc.input, tc.pattern), tc.match, tc.input, tc.pattern)
	}
}

/*
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like ? or *.

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

Input:
s = "aa"
p = "*"
Output: true
Explanation: '*' matches any sequence.

Input:
s = "cb"
p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.

Input:
s = "adceb"
p = "*a*b"
Output: true
Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".
*/

func isMatch(s string, p string) bool {
	sIdx := 0
	pIdx := 0
	currentMatchIdx := 0
	skipIdx := -1

	for sIdx < len(s) {
		// if we have a match - advance pattern and string
		if pIdx < len(p) && (s[sIdx] == p[pIdx] || p[pIdx] == '?') {
			sIdx++
			pIdx++

			continue
		}

		// if we have * - compress all *'s into one
		// save position in pattern at which we found *
		// also save position in string at which we found * in pattern
		if pIdx < len(p) && p[pIdx] == '*' {
			skipIdx = pIdx
			currentMatchIdx = sIdx
			pIdx++

			continue
		}

		// if we had * in pattern then advance to next position in string and continue checks
		if skipIdx > -1 {
			pIdx = skipIdx + 1 // set this to start matching after * in pattern
			currentMatchIdx++
			sIdx = currentMatchIdx

			continue
		}

		// no symbol matches and no * in pattern
		return false
	}

	// if pattern ends on * - skip it
	for pIdx < len(p) && p[pIdx] == '*' {
		pIdx++
	}

	// if we reached end of pattern - we matched a string
	return pIdx == len(p)
}

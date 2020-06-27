package main

func main() {
	println(isMatch("aa", "a*"), true)
	println(isMatch("ab", ".*"), true)
	println(isMatch("aab", "c*a*b*"), true)
	println(isMatch("mississippi", "mis*is*p*."), false)
	println(isMatch("aaa", "a*a"), true)
	println(isMatch("aaaa", "a*a"), true)
	println(isMatch("a", "ab*"), true)
	println(isMatch("bbbba", ".*a*a"), true)
	println(isMatch("ab", ".*c"), false)
}

/*
https://leetcode.com/problems/regular-expression-matching/

Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.

Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

Example 2:

Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

Example 3:

Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".

Example 4:

Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".

Example 5:

Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
*/

func isMatch(s string, p string) bool {
	lp := len(p)
	ls := len(s)

	if lp == 0 {
		// end of pattern & string reached
		return ls == 0
	}

	if ls == 0 {
		// string end reached
		if lp > 1 && p[1] == '*' {
			// we have * - try to match in "zero chars matched" mode
			return isMatch(s, p[2:])
		}

		return false
	}

	if p[0] != '.' && s[0] != p[0] {
		// no matches
		if lp > 1 && p[1] == '*' {
			// we have * - try to match in "zero chars matched" mode
			return isMatch(s, p[2:])
		}

		return false
	}

	if lp > 1 && p[1] == '*' {
		// try to match next char or in "zero chars matched" mode
		return isMatch(s[1:], p) || isMatch(s, p[2:])
	}

	// chars matched, go on with next symbols in string and pattern
	return isMatch(s[1:], p[1:])
}

// this is faster and less memory expensive
func isMatchDP(s string, p string) bool {
	ls := len(s) + 1
	lp := len(p) + 1
	matches := make([][]bool, ls)

	for i := 0; i < ls; i++ {
		matches[i] = make([]bool, lp)
	}

	matches[0][0] = true

	for y := 1; y < lp; y++ {
		if p[y-1] == '*' {
			matches[0][y] = matches[0][y-2]
		}
	}

	for x := 1; x < ls; x++ {
		c := s[x-1]
		for y := 1; y < lp; y++ {
			t := p[y-1]
			var tPrev byte
			if y > 1 {
				tPrev = p[y-2]
			}

			if t == '.' || t == c {
				// take prev match
				matches[x][y] = matches[x-1][y-1]
			} else if t == '*' {
				matches[x][y] = matches[x][y-2]

				if tPrev == '.' || tPrev == c {
					matches[x][y] = matches[x][y] || matches[x-1][y]
				}
			}
		}
	}

	return matches[ls-1][lp-1]
}

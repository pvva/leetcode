package main

func main() {
	println(balancedString("WQWRQQQW"), 3)
	println(balancedString("QWER"), 0)
	println(balancedString("QQWQRRWRRQWWEWRQEREWQQWQREWRWQRQRWRQEQWE"), 4)
	println(balancedStringBetter("WQWRQQQW"), 3)
	println(balancedStringBetter("QWER"), 0)
	println(balancedStringBetter("QQWQRRWRRQWWEWRQEREWQQWQREWRWQRQRWRQEQWE"), 4)
}

/*
https://leetcode.com/problems/replace-the-substring-for-balanced-string/

You are given a string containing only 4 kinds of characters 'Q', 'W', 'E' and 'R'.
A string is said to be balanced if each of its characters appears n/4 times where n is the length of the string.
Return the minimum length of the substring that can be replaced with any other string of the same length to make
the original string s balanced.
Return 0 if the string is already balanced.

Example 1:

Input: s = "QWER"
Output: 0
Explanation: s is already balanced.

Example 2:

Input: s = "QQWE"
Output: 1
Explanation: We need to replace a 'Q' to 'R', so that "RQWE" (or "QRWE") is balanced.

Example 3:

Input: s = "QQQW"
Output: 2
Explanation: We can replace the first "QQ" to "ER".

Example 4:

Input: s = "QQQQ"
Output: 3
Explanation: We can replace the last 3 'Q' to make s = "QWER".
*/

var qwerIdx = map[byte]int{'Q': 0, 'W': 1, 'E': 2, 'R': 3}
var qwers = []byte{'Q', 'W', 'E', 'R'}

func balancedString(s string) int {
	limit := len(s) / 4

	qwer := [4]int{}
	for _, c := range s {
		qwer[qwerIdx[byte(c)]]++
	}

	extraChars := make(map[byte]int)
	for i := 0; i < 4; i++ {
		qwer[i] -= limit
		if qwer[i] > 0 {
			extraChars[qwers[i]] = qwer[i]
		}
	}
	if len(extraChars) > 0 {
		return findMinSubstrWithChars(s, extraChars)
	}

	return 0
}

func findMinSubstrWithChars(s string, chars map[byte]int) int {
	cChars := make(map[byte]int)

	min := len(s)

	si := 0
	for si < len(s) {
		if _, ex := chars[s[si]]; ex {
			break
		}
		si++
	}

	ei := si
	for ei < len(s) {
		cChars[s[ei]]++

		for si < ei && hasAllChars(cChars, chars) {
			t := ei + 1 - si
			if t < min {
				min = t
			}

			cChars[s[si]]--
			si++
		}
		ei++
	}
	for si < len(s) && hasAllChars(cChars, chars) {
		cChars[s[si]]--
		si++
	}
	t := ei + 1 - si
	if t < min {
		return t
	}

	return min
}

func hasAllChars(target, origin map[byte]int) bool {
	for k, v := range origin {
		if target[k] < v {
			return false
		}
	}

	return true
}

func balancedStringBetter(s string) int {
	limit := len(s) / 4

	cQ, cW, cE, cR := 0, 0, 0, 0
	for _, c := range s {
		switch c {
		case 'Q':
			cQ++
		case 'W':
			cW++
		case 'E':
			cE++
		case 'R':
			cR++
		}
	}
	cQ -= limit
	cW -= limit
	cE -= limit
	cR -= limit

	if cQ == 0 && cW == 0 && cE == 0 && cR == 0 {
		return 0
	}

	if cQ < 0 {
		cQ = 0
	}
	if cW < 0 {
		cW = 0
	}
	if cE < 0 {
		cE = 0
	}
	if cR < 0 {
		cR = 0
	}

	si := 0
	ei := 0
	min := len(s)
	rQ, rW, rE, rR := 0, 0, 0, 0

	for si <= ei {
		if rQ >= cQ && rW >= cW && rE >= cE && rR >= cR {
			if ei-si < min {
				min = ei - si
			}

			switch s[si] {
			case 'Q':
				rQ--
			case 'W':
				rW--
			case 'E':
				rE--
			case 'R':
				rR--
			}

			si++
		} else {
			if ei == len(s) {
				break
			}
			switch s[ei] {
			case 'Q':
				rQ++
			case 'W':
				rW++
			case 'E':
				rE++
			case 'R':
				rR++
			}
			ei++
		}
	}

	return min
}

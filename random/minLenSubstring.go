package main

func main() {
	println(minLengthSubstring("dcbefebce", "fd"), 5)
	println(minLengthSubstring("dcbefbce", "fb"), 2)
}

/*
You are given two strings s and t. You can select any substring of string s and rearrange the characters of the
selected substring. Determine the minimum length of the substring of s such that string t is a substring of the
selected substring.

Example
s = "dcbefebce"
t = "fd"
output = 5

Explanation:
Substring "dcbef" can be rearranged to "cfdeb", "cefdb", and so on. String t is a substring of "cfdeb". Thus, the
minimum length required is 5.
*/

func minLengthSubstring(s string, t string) int {
	target := make(map[byte]int)
	for _, c := range t {
		target[byte(c)]++
	}

	seenChars := make(map[byte]int)
	minSubstrLen := len(s)
	si := -1
	ei := 0

	for ei < len(s) {
		si = findNextStartIdx(s, target, seenChars, si+1)
		if si >= len(s) {
			break
		}
		if si >= ei {
			seenChars[s[si]]++
			ei = si + 1
		}
		ei = findNextEndIdx(s, target, seenChars, ei)
		if ei == -1 {
			break
		}
		if ei-si < minSubstrLen {
			minSubstrLen = ei - si
		}

		dropSeenChar(seenChars, s[si])
	}

	return minSubstrLen
}

func dropSeenChar(seenChars map[byte]int, c byte) {
	v, ex := seenChars[c]
	if !ex {
		return
	}
	if v <= 1 {
		delete(seenChars, c)
	} else {
		seenChars[c] = v - 1
	}
}

func findNextStartIdx(s string, target, seenChars map[byte]int, from int) int {
	for from < len(s) {
		c := s[from]
		if _, ex := target[c]; ex {
			break
		}
		dropSeenChar(seenChars, c)
		from++
	}
	return from
}

func findNextEndIdx(s string, target, seenChars map[byte]int, from int) int {
	for from < len(s) {
		seenChars[s[from]]++
		from++

		if sameSet(seenChars, target) {
			return from
		}
	}

	return -1
}

func sameSet(current, target map[byte]int) bool {
	for k, v := range target {
		if v2, ex := current[k]; !ex || v != v2 {
			return false
		}
	}

	return true
}

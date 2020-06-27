package main

import "math"

func main() {
	println(minWindow("ADOBECODEBANC", "ABC"))
	println(minWindow("bba", "ab"))
	println(minWindow("acbbaca", "aba"))
}

/*
https://leetcode.com/problems/minimum-window-substring/

Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
*/

func minWindow(s string, t string) string {
	mR := make(map[string]int)

	for i := 0; i < len(t); i++ {
		c := t[i : i+1]
		if v, ex := mR[c]; ex {
			mR[c] = v + 1
		} else {
			mR[c] = 1
		}
	}
	diff := len(t)
	minLen := math.MaxInt32
	start := 0
	head := 0
	idx := 0

	for idx < len(s) {
		c := s[idx : idx+1]

		if v, ex := mR[c]; ex && v > 0 {
			diff--
		}
		v, _ := mR[c]
		mR[c] = v - 1
		idx++

		for diff == 0 {
			if idx-start < minLen {
				minLen = idx - start
				head = start
			}

			c0 := s[start : start+1]
			if v0, ex := mR[c0]; ex && v0 == 0 {
				diff++
			}
			v0, _ := mR[c0]
			mR[c0] = v0 + 1
			start++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return s[head : head+minLen]
}

package main

func main() {
	println(3, numDistinct("rabbbit", "rabbit"))
	println(5, numDistinct("babgbag", "bag"))
}

/*
https://leetcode.com/problems/distinct-subsequences/

Given two strings s and t, return the number of distinct subsequences of s which equals t.

A string's subsequence is a new string formed from the original string by deleting some (can be none) of the characters
without disturbing the remaining characters' relative positions. (i.e., "ACE" is a subsequence of "ABCDE" while "AEC" is not).

It is guaranteed the answer fits on a 32-bit signed integer.

Example 1:

Input: s = "rabbbit", t = "rabbit"
Output: 3
Explanation:
As shown below, there are 3 ways you can generate "rabbit" from S.
rabbbit
rabbbit
rabbbit

Example 2:

Input: s = "babgbag", t = "bag"
Output: 5
Explanation:
As shown below, there are 5 ways you can generate "bag" from S.
babgbag
babgbag
babgbag
babgbag
babgbag

Constraints:

1 <= s.length, t.length <= 1000
s and t consist of English letters.
*/

func numDistinctRec(s string, sIdx int, t string, tIdx int, mem [][]int) int {
	if len(s)-sIdx < len(t)-tIdx {
		return 0
	}
	if tIdx == len(t) {
		return 1
	}
	if sIdx == len(s) {
		return 0
	}
	if mem[sIdx][tIdx] >= 0 {
		// we have seen this combination of positions before, so we can return it
		return mem[sIdx][tIdx]
	}

	if s[sIdx] == t[tIdx] {
		// at this point we either shift both strings forward or just first one as it can have more characters
		// that would match second string
		p1 := numDistinctRec(s, sIdx+1, t, tIdx+1, mem)
		p2 := numDistinctRec(s, sIdx+1, t, tIdx, mem)
		mem[sIdx][tIdx] = p1 + p2
	} else {
		// at this point portion of first string cannot be matched against second, so advance first one
		mem[sIdx][tIdx] = numDistinctRec(s, sIdx+1, t, tIdx, mem)
	}

	return mem[sIdx][tIdx]
}

func numDistinct(s string, t string) int {
	if len(t) > len(s) {
		return 0
	}
	freq := [int('z' - 'A' + 1)]int{}
	for i := range s {
		c := s[i] - 'A'
		freq[c]++
	}
	for i := range t {
		c := s[i] - 'A'
		freq[c]--
		if freq[c] < 0 {
			return 0
		}
	}

	mem := make([][]int, len(s))
	for i := range mem {
		mem[i] = make([]int, len(t))
		for j := 0; j < len(t); j++ {
			mem[i][j] = -1
		}
	}

	return numDistinctRec(s, 0, t, 0, mem)
}

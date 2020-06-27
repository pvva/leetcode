package main

func main() {
	println(isSubsequence("abc", "ahbgdc"))
	println(isSubsequence("axc", "ahbgdc"))
}

/*
https://leetcode.com/problems/is-subsequence/

Given a string s and a string t, check if s is subsequence of t.

A subsequence of a string is a new string which is formed from the original string by deleting some (can be none)
of the characters without disturbing the relative positions of the remaining characters. (ie, "ace" is a subsequence
of "abcde" while "aec" is not).

Follow up:
If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you want to check one by one to see if T has
its subsequence. In this scenario, how would you change your code?

Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false
*/

func isSubsequence(s string, t string) bool {
	ti := 0
	si := 0

	for si < len(s) && ti < len(t) {
		if s[si] == t[ti] {
			si++
		}
		ti++
	}

	return si == len(s)
}

// the answer on follow up is to use directed graph to store content of target string in the following way:
// t[i] is connected to t[i+1..n] where i < n
// then for every incoming string s it is a matter of finding the path in the graph if it exists

package main

func main() {
	println(longestPalindrome("babadada"), "adada")
}

/*
https://leetcode.com/problems/longest-palindromic-substring/

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:

Input: "cbbd"
Output: "bb"
*/

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	maxLen := 1
	maxStr := string(s[0])

	for l := 2; l <= 3; l++ {
		for i := 0; i <= len(s)-l; i++ {
			if isPalindrome(s[i : i+l]) {

				if maxLen < l {
					maxLen = l
					maxStr = s[i : i+l]
				}

				for j := 1; j <= min(i, len(s)-i-l) && s[i-j] == s[i+l-1+j]; j++ {
					cLen := l + 2*j - 1
					if maxLen < cLen {
						maxLen = cLen
						maxStr = s[i-j : i+l+j]
					}
				}
			}
		}
	}

	return maxStr
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)>>1; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

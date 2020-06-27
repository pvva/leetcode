package main

func main() {
	println(lengthOfLongestSubstring("aabaab!bb"), 3)
}

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/

Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	sIdx := 0
	maxSubstrLen := 0
	seenUnique := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		c := s[i]

		if seenUnique[c] {
			for sIdx < i {
				ic := s[sIdx]
				sIdx++
				delete(seenUnique, ic)
				if ic == c {
					break
				}
			}
		}
		seenUnique[c] = true

		if len(seenUnique) > maxSubstrLen {
			maxSubstrLen = len(seenUnique)
		}
	}

	return maxSubstrLen
}

package main

func main() {
	println(1, minCut("aab"))
	println(0, minCut("a"))
	println(1, minCut("ab"))
	println(1, minCut("aaab"))
	println(1, minCut("abaab"))
	println(1, minCut("abb"))
}

/*
https://leetcode.com/problems/palindrome-partitioning-ii/

Given a string s, partition s such that every substring of the partition is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s.

Example 1:

Input: s = "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.

Example 2:

Input: s = "a"
Output: 0
Example 3:

Input: s = "ab"
Output: 1

Constraints:

1 <= s.length <= 2000
s consists of lower-case English letters only.
*/

func findAllPalindromes(s string) [][]bool {
	res := make([][]bool, len(s))
	for i := range res {
		res[i] = make([]bool, len(s))
		// single letter palindrome
		res[i][i] = true
	}

	for palinLen := 2; palinLen <= len(s); palinLen++ {
		for startIdx := 0; startIdx <= len(s)-palinLen; startIdx++ {
			endIdx := startIdx + palinLen - 1

			res[startIdx][endIdx] = s[startIdx] == s[endIdx] && (palinLen == 2 || res[startIdx+1][endIdx-1])
		}
	}

	return res
}

func minCut(s string) int {
	allPalins := findAllPalindromes(s)
	cuts := make([]int, len(s))

	for endIdx := range s {
		if allPalins[0][endIdx] {
			// string part from 0 to endIdx (inclusive) is palindrome,
			// no cut at index endIdx is required
			cuts[endIdx] = 0
		} else {
			cuts[endIdx] = len(s)

			// start index of 0 is checked before
			for startIdx := 1; startIdx <= endIdx; startIdx++ {
				if allPalins[startIdx][endIdx] {
					// found palindrome, add another cut at startIdx
					newCuts := cuts[startIdx-1] + 1 // number of cuts up to startIdx + new cut at startIdx
					if newCuts < cuts[endIdx] {
						cuts[endIdx] = newCuts
					}
				}
			}
		}
	}

	return cuts[len(s)-1]
}

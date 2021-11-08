package main

func main() {
	println(3, minDistance("horse", "ros"))
	println(5, minDistance("intention", "execution"))
}

/*
https://leetcode.com/problems/edit-distance/

Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character

Example 1:

Input: word1 = "horse", word2 = "ros"

Output: 3

Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

Example 2:

Input: word1 = "intention", word2 = "execution"

Output: 5

Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')

Constraints:

0 <= word1.length, word2.length <= 500
word1 and word2 consist of lowercase English letters.
*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	} else if len(word2) == 0 {
		return len(word1)
	}
	currDistance := [2][]int{}
	currDistance[0] = make([]int, len(word2))
	currDistance[1] = make([]int, len(word2))
	for i := 0; i < len(word2); i++ {
		currDistance[0][i] = i + 1
	}

	currRowIdx := 1
	prevRowIdx := 0

	for i1 := 1; i1 <= len(word1); i1++ {
		s1 := word1[i1-1]
		for i2 := 0; i2 < len(word2); i2++ {
			s2 := word2[i2]
			minPrev := i1 - 1
			minCurr := i1
			if i2 > 0 {
				minPrev = currDistance[prevRowIdx][i2-1]
				minCurr = currDistance[currRowIdx][i2-1]
			}
			if s1 == s2 {
				currDistance[currRowIdx][i2] = minPrev
			} else {
				currDistance[currRowIdx][i2] = min(min(minCurr, currDistance[prevRowIdx][i2]), minPrev) + 1
			}
		}
		currRowIdx = 1 - currRowIdx
		prevRowIdx = 1 - prevRowIdx
	}
	// last "current" row idx is now prevRowIdx

	return currDistance[prevRowIdx][len(word2)-1]
}

package main

import "math"

/*
https://leetcode.com/problems/strong-password-checker

A password is considered strong if below conditions are all met:

It has at least 6 characters and at most 20 characters.
It must contain at least one lowercase letter, at least one uppercase letter, and at least one digit.
It must NOT contain three repeating characters in a row ("...aaa..." is weak, but "...aa...a..." is strong,
assuming other conditions are met).
Write a function strongPasswordChecker(s), that takes a string s as input, and return the MINIMUM change required
to make s a strong password. If s is already strong, return 0.

Insertion, deletion or replace of any one character are all considered as one change.
*/

type strongPasswordCheckerTestCase struct {
	password         string
	expectedDistance int
}

func main() {
	testCases := []strongPasswordCheckerTestCase{
		{"a", 5},
		{"aa", 4},
		{"aaa", 3},
		{"123456789aaa123456789", 2},
		{"aaaaaaaAAAAAA6666bbbbaaaaaaABBC", 13},
	}

	for _, tc := range testCases {
		println("Password:", tc.password, " expected distance:", tc.expectedDistance,
			" calculated distance:", strongPasswordChecker(tc.password))
	}
}

func strongPasswordChecker(s string) int {
	distance := 0

	l := len(s)
	lcCorr := 1
	ucCorr := 1
	dcCorr := 1
	totalReplacementsCount := 0
	replacements := []int{}
	currentLetter := byte(0)
	currentContinuousFragmentLen := 0
	for i := 0; i < l; i++ {
		lc := isLowerCase(s[i])
		uc := isUpperCase(s[i])
		dc := isDigit(s[i])
		if lc {
			lcCorr = 0
		}
		if uc {
			ucCorr = 0
		}
		if dc {
			dcCorr = 0
		}
		if s[i] != currentLetter {
			if currentContinuousFragmentLen > 2 {
				totalReplacementsCount += currentContinuousFragmentLen / 3
				replacements = append(replacements, currentContinuousFragmentLen)
			}
			currentLetter = s[i]
			currentContinuousFragmentLen = 1
		} else {
			currentContinuousFragmentLen++
		}
	}
	if currentContinuousFragmentLen > 2 {
		totalReplacementsCount += currentContinuousFragmentLen / 3
		replacements = append(replacements, currentContinuousFragmentLen)
	}
	corrections := 0

	if l > 20 {
		distance = l - 20
		if totalReplacementsCount > distance {
			corrections = totalReplacementsCount - distance
		} else {
			if len(replacements) > 0 {
				corrections = maxInt(lcCorr+ucCorr+dcCorr, minimizeReplacementsAndRemovals(replacements, distance))
			} else {
				corrections = maxInt(lcCorr+ucCorr+dcCorr, totalReplacementsCount)
			}
		}
	} else if l < 6 {
		distance = 6 - l
		tc := lcCorr + ucCorr + dcCorr
		additions := distance - tc
		if additions < 0 {
			corrections = -additions
		}
	} else {
		corrections = maxInt(lcCorr+ucCorr+dcCorr, totalReplacementsCount)
	}

	return distance + corrections
}

func minimizeReplacementsAndRemovals(replacements []int, removals int) int {
	l := len(replacements)
	dp := make([][]int, removals+1)
	for i := 0; i <= removals; i++ {
		dp[i] = make([]int, l+1)
		for j := 1; j <= l; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	for j := 1; j <= l; j++ {
		dp[0][j] = dp[0][j-1] + replacements[j-1]/3
	}
	dp[0][0] = 0

	for r := 1; r <= removals; r++ {
		for p := 1; p <= l; p++ {
			removeLimit := replacements[p-1] - 2
			if r < removeLimit {
				removeLimit = r
			}

			for k := 0; k <= removeLimit; k++ {
				dp[r][p] = minInt(dp[r][p], dp[k][p-1]+(replacements[p-1]-k)/3)
			}
		}
	}

	if dp[removals][l] != math.MaxInt32 {
		return dp[removals][l]
	}

	return 0
}

func isLowerCase(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isUpperCase(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

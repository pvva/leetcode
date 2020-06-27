package main

func main() {
	println(maxFreq("aababcaab", 2, 3, 4), 2)
	println(maxFreq("aaaa", 1, 3, 3), 2)
	println(maxFreq("abcde", 2, 3, 3), 0)
}

/*
https://leetcode.com/problems/maximum-number-of-occurrences-of-a-substring/

Given a string s, return the maximum number of occurrences of any substring under the following rules:

The number of unique characters in the substring must be less than or equal to maxLetters.
The substring size must be between minSize and maxSize inclusive.


Example 1:

Input: s = "aababcaab", maxLetters = 2, minSize = 3, maxSize = 4
Output: 2
Explanation: Substring "aab" has 2 ocurrences in the original string.
It satisfies the conditions, 2 unique letters and size 3 (between minSize and maxSize).
Example 2:

Input: s = "aaaa", maxLetters = 1, minSize = 3, maxSize = 3
Output: 2
Explanation: Substring "aaa" occur 2 times in the string. It can overlap.
*/

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	maxOccurrences := 0
	occurrences := make(map[string]int)

	for i := 0; i <= len(s)-minSize; i++ {
		subStr := s[i : i+minSize]
		if c, ex := occurrences[subStr]; ex {
			maxOccurrences = incrementOccurrenceAndGetMax(occurrences, c, maxOccurrences, subStr)
		} else {
			uniqueChars := countUniqueChars(subStr)
			if uniqueChars <= maxLetters {
				c, _ := occurrences[subStr]
				maxOccurrences = incrementOccurrenceAndGetMax(occurrences, c, maxOccurrences, subStr)
			}
		}
	}

	return maxOccurrences
}

func incrementOccurrenceAndGetMax(occurrences map[string]int, c, cMax int, s string) int {
	c++
	if cMax < c {
		cMax = c
	}
	occurrences[s] = c

	return cMax
}

func countUniqueChars(s string) int {
	mask := 0

	cnt := 0
	for _, c := range s {
		bit := 1 << (c - 'a')
		if bit&mask != bit {
			cnt++
			mask |= bit
		}
	}

	return cnt
}

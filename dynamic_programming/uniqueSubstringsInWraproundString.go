package main

func main() {
	println(findSubstringInWraproundString("a"), 1)
	println(findSubstringInWraproundString("cac"), 2)
	println(findSubstringInWraproundString("zab"), 6)
}

/*
https://leetcode.com/problems/unique-substrings-in-wraparound-string/

Consider the string s to be the infinite wraparound string of "abcdefghijklmnopqrstuvwxyz", so s will look
like this: "...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".

Now we have another string p. Your job is to find out how many unique non-empty substrings of p are present
in s. In particular, your input is the string p and you need to output the number of different non-empty
substrings of p in the string s.

Note: p consists of only lowercase English letters and the size of p might be over 10000.

Example 1:
Input: "a"
Output: 1

Explanation: Only the substring "a" of string "a" is in the string s.

Example 2:
Input: "cac"
Output: 2
Explanation: There are two substrings "a", "c" of string "cac" in the string s.

Example 3:
Input: "zab"
Output: 6
Explanation: There are six substrings "z", "a", "b", "za", "ab", "zab" of string "zab" in the string s.
*/

func findSubstringInWraproundString(p string) int {
	if p == "" {
		return 0
	}

	substrLengthsEndingWithLetter := [26]int{}
	substrLengthsEndingWithLetter[int(p[0]-'a')] = 1

	maxLen := 1
	for i := 1; i < len(p); i++ {
		idx := int(p[i] - 'a')
		if substrLengthsEndingWithLetter[idx] == 0 {
			substrLengthsEndingWithLetter[idx] = 1
		}

		if p[i]-p[i-1] == 1 || (p[i] == 'a' && p[i-1] == 'z') {
			maxLen++
			if substrLengthsEndingWithLetter[idx] < maxLen {
				substrLengthsEndingWithLetter[idx] = maxLen
			}
		} else {
			maxLen = 1
		}
	}

	result := 0
	for _, l := range substrLengthsEndingWithLetter {
		result += l
	}

	return result
}

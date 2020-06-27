package main

func main() {
	println(wordBreak("leetcode", []string{"leet", "code"}), true)
	println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), false)
}

/*
https://leetcode.com/problems/word-break/

Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be
segmented into a space-separated sequence of one or more dictionary words.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Example 1:

Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:

Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
*/

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) < len(s) {
		return wbN2(s, wordDict)
	}

	return wbNM(s, wordDict)
}

// shift over string one by one and check that we can match any of word dict entries that end at current position
func wbNM(s string, wordDict []string) bool {
	l := len(s)
	canBreak := make([]bool, l+1)
	canBreak[0] = true
	for i := 1; i <= l; i++ {
		for _, w := range wordDict {
			lw := len(w)
			if lw <= i && s[i-lw:i] == w {
				canBreak[i] = true
				break
			}
		}
	}

	return canBreak[l]
}

// extract parts of string and check their presence in word dict
func wbN2(s string, wordDict []string) bool {
	wd := make(map[string]bool)
	for _, w := range wordDict {
		wd[w] = true
	}

	l := len(s)
	canBreak := make([]bool, l+1)
	canBreak[0] = true
	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			if canBreak[j] && wd[s[j:i]] {
				canBreak[i] = true
				break
			}
		}
	}

	return canBreak[l]
}

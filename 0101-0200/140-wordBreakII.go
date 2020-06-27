package main

func main() {
	printArr(wordBreak2("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	//printArr(wordBreak2("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	//printArr(wordBreak2("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func printArr(arr []string) {
	if len(arr) == 0 {
		println("[]")
	} else {
		for _, s := range arr {
			println(s)
		}
	}
}

/*
https://leetcode.com/problems/word-break-ii/

Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to construct
a sentence where each word is a valid dictionary word. Return all such possible sentences.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
Output:
[
  "cats and dog",
  "cat sand dog"
]
Example 2:

Input:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
Output:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
Explanation: Note that you are allowed to reuse a dictionary word.

Example 3:

Input:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
Output:
[]
*/

func wordBreak2(s string, wordDict []string) []string {
	idxs := make(map[int][]string)

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

				idx, e := idxs[i]
				if !e {
					idx = []string{}
				}
				idx = append(idx, s[j:i])
				idxs[i] = idx
			}
		}
	}

	if canBreak[l] {
		return fillContent(l, idxs)
	}

	return []string{}
}

func fillContent(idx int, idxs map[int][]string) []string {
	ar := idxs[idx]
	content := []string{}

	for _, w := range ar {
		c := fillContent(idx-len(w), idxs)
		if len(c) > 0 {
			for _, v := range c {
				content = append(content, v+" "+w)
			}
		} else {
			content = append(content, w)
		}
	}

	return content
}

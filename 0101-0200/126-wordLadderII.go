package main

import "fmt"

func main() {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(findLadders("qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db",
		"mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra",
		"fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh",
		"sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co",
		"ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo",
		"na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb",
		"ni", "mr", "pa", "he", "lr", "sq", "ye"}))
}

/*
https://leetcode.com/problems/word-ladder-ii/

Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s)
from beginWord to endWord, such that:

Only one letter can be changed at a time
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
Note:

Return an empty list if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]
Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: []

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
*/

// This solution executes classic BFS algorithm.
// There's a solution that explicitly modifies the word by 1 letter, scans words dictionary and advances
// in case an entry is found. Such algorithm performs faster in this particular case, while provided
// implementation is generic.
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wl := make(map[string]bool)
	for _, w := range wordList {
		if w == beginWord {
			continue
		}
		wl[w] = true
	}

	stack := [][]string{{beginWord}}
	foundShortest := false
	for len(stack) > 0 && !foundShortest {
		totalNext := [][]string{}

		cLevel := make(map[string]bool)
		for _, cPath := range stack {
			from := cPath[len(cPath)-1]
			delete(wl, from)
			cLevel[from] = true
		}
		for _, cPath := range stack {
			from := cPath[len(cPath)-1]
			next, found := executeBFSWithResults(from, endWord, wl)

			if found {
				foundShortest = true
			}

			for _, step := range next {
				if (foundShortest && step != endWord) || cLevel[step] {
					continue
				}
				tPath := make([]string, len(cPath))
				copy(tPath, cPath)
				totalNext = append(totalNext, append(tPath, step))
			}
		}

		stack = totalNext
	}

	result := [][]string{}
	for _, path := range stack {
		if path[len(path)-1] == endWord {
			result = append(result, path)
		}
	}

	return result
}

func executeBFSWithResults(fromWord, endWord string, wl map[string]bool) ([]string, bool) {
	next := getNearestWords(fromWord, wl)

	for _, w := range next {
		if w == endWord {
			return next, true
		}
	}

	return next, false
}

func getNearestWords(word string, wl map[string]bool) []string {
	result := []string{}

	for w := range wl {
		if isWordNear(word, w) {
			result = append(result, w)
		}
	}

	return result
}

func isWordNear(word, target string) bool {
	diff := 0
	for i := 0; i < len(word); i++ {
		if word[i:i+1] != target[i:i+1] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}

	return diff == 1
}

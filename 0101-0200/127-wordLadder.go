package main

func main() {
	println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}), 5)
	println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}), 0)
}

/*
https://leetcode.com/problems/word-ladder/

Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest
transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list.
Note:

Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.

Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output: 5

Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
return its length 5.

Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: 0

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	counter := 1
	fromWords := []string{beginWord}
	found := false

	wordsMap := make(map[string]bool)
	for _, w := range wordList {
		wordsMap[w] = true
	}

	for !found && len(fromWords) > 0 {
		fromWords, found = findNextBfsLayer(fromWords, wordsMap, endWord)
		counter++
	}

	if found {
		return counter
	}
	return 0
}

func findNextBfsLayer(fromWords []string, words map[string]bool, targetWord string) ([]string, bool) {
	nextLayer := []string{}
	if len(words) == 0 {
		return nextLayer, false
	}
	for _, word := range fromWords {
		nextWords, found := findSingleLetterDifferentWords(word, words, targetWord)

		if found {
			return nil, true
		}

		if len(nextWords) == 0 {
			continue
		}

		nextLayer = append(nextLayer, nextWords...)
		if len(words) == 0 {
			break
		}
	}

	return nextLayer, false
}

func findSingleLetterDifferentWords(word string, words map[string]bool, endWord string) ([]string, bool) {
	result := []string{}
	dummy := []byte(word)
outer:
	for i, cw := range word {
		for c := 'a'; c <= 'z'; c++ {
			if c != cw {
				dummy[i] = byte(c)
				nw := string(dummy)
				if words[nw] {
					if nw == endWord {
						return nil, true
					}
					delete(words, nw)
					result = append(result, nw)

					if len(words) == 0 {
						break outer
					}
				}
			}
		}
		dummy[i] = byte(cw)
	}

	return result, false
}

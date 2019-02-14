package main

func main() {
	println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}

/*
Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
Note:

Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wl := make(map[string]bool)
	for _, w := range wordList {
		if w == beginWord {
			continue
		}
		wl[w] = true
	}

	stack := []string{beginWord}
	count := 1
	for len(stack) > 0 {
		count++

		totalNext := []string{}

		for _, cw := range stack {
			next, found := executeBFS(cw, endWord, wl)

			if found {
				return count
			}

			totalNext = append(totalNext, next...)
		}

		stack = totalNext
	}

	return 0
}

func executeBFS(word, endWord string, wl map[string]bool) ([]string, bool) {
	next := getNearest(word, wl)

	for _, w := range next {
		if w == endWord {
			return next, true
		}
	}

	return next, false
}

func getNearest(word string, wl map[string]bool) []string {
	result := []string{}

	for w := range wl {
		if isNear(word, w) {
			result = append(result, w)
			delete(wl, w)
		}
	}

	return result
}

func isNear(word, target string) bool {
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

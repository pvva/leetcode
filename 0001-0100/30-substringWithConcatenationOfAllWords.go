package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))

	fmt.Println(findSubstring2("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring2("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring2("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring2("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
	fmt.Println(findSubstring2("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

/*
https://leetcode.com/problems/substring-with-concatenation-of-all-words/

You are given a string s and an array of strings words. All the strings of words are of the same length.
A concatenated substring in s is a substring that contains all the strings of any permutation of words concatenated.

For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are
all concatenated strings. "acdbef" is not a concatenated substring because it is not the concatenation of any
permutation of words. Return the starting indices of all the concatenated substrings in s.
You can return the answer in any order.

Example 1:

Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0,9]
Explanation: Since words.length == 2 and words[i].length == 3, the concatenated substring has to be of length 6.
The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.
The output order does not matter. Returning [9,0] is fine too.
Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
Output: []
Explanation: Since words.length == 4 and words[i].length == 4, the concatenated substring has to be of length 16.
There is no substring of length 16 is s that is equal to the concatenation of any permutation of words.
We return an empty array.
Example 3:

Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
Output: [6,9,12]
Explanation: Since words.length == 3 and words[i].length == 3, the concatenated substring has to be of length 9.
The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"] which is a permutation of words.
The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"] which is a permutation of words.
The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"] which is a permutation of words.

Constraints:

1 <= s.length <= 104
1 <= words.length <= 5000
1 <= words[i].length <= 30
s and words[i] consist of lowercase English letters.
*/

func findSubstring(s string, words []string) []int {
	res := []int{}

	root := newTrie()
	trieSaves := []*trieSave{}
	minWordLen := len(s)

	for _, w := range words {
		trieNode := root.AddString(w)
		if trieNode.count == 1 {
			trieSaves = append(trieSaves, &trieSave{node: trieNode})
		}
		if len(w) < minWordLen {
			minWordLen = len(w)
		}
	}
	for _, save := range trieSaves {
		save.count = save.node.count
	}

	allWordsCount := len(words)

	for i := 0; i < len(s); i++ {
		nextI := i
		for nextI >= 0 {
			n := root.findNext(s, nextI)
			if n == -1 {
				//i += minWordLen - 1
				resetTrie(trieSaves)
				allWordsCount = len(words)

				break
			}
			nextI = n
			allWordsCount--
			if allWordsCount == 0 {
				// found all words
				resetTrie(trieSaves)
				allWordsCount = len(words)

				res = append(res, i)
				//i += minWordLen - 1
				break
			}
		}
	}

	return res
}

type trieSave struct {
	node  *trie
	count int
}

func resetTrie(saves []*trieSave) {
	for _, save := range saves {
		save.node.count = save.count
	}
}

type trie struct {
	count int
	next  map[byte]*trie
}

func newTrie() *trie {
	return &trie{
		next: make(map[byte]*trie),
	}
}

func (trie *trie) AddString(s string) *trie {
	ct := trie
	for _, r := range []byte(s) {
		nt, ex := ct.next[r]
		if !ex {
			nt = newTrie()
			ct.next[r] = nt
		}
		ct = nt
	}
	ct.count++

	return ct
}

func (trie *trie) findNext(s string, idx int) int {
	ct := trie
	for i := idx; i < len(s); i++ {
		nt, ex := ct.next[s[i]]
		if !ex {
			break
		}
		if nt.count > 0 {
			nt.count--

			return i + 1
		}
		ct = nt
	}

	return -1
}

// this is much faster
func findSubstring2(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	res := []int{}
	wLen := len(words[0])
	wordsMap := make(map[string]int)
	for _, w := range words {
		cnt, _ := wordsMap[w]
		cnt++
		wordsMap[w] = cnt
	}

	seenWords := make(map[string]int)

	for i := 0; i < len(s)-wLen*len(words)+1; i++ {
		word := s[i : i+wLen]

		if _, ex := wordsMap[word]; ex {
			seenWords[word] = 1

			j := 1
			for ; j < len(words); j++ {
				inSequence := false
				nextWord := s[i+j*wLen : i+(j+1)*wLen]

				if limit, ex := wordsMap[nextWord]; ex {
					cnt, _ := seenWords[nextWord]
					cnt++
					seenWords[nextWord] = cnt
					if cnt <= limit {
						inSequence = true
					}
				}

				if !inSequence {
					break
				}
			}
			if j == len(words) {
				res = append(res, i)
			}
			seenWords = make(map[string]int)
		}
	}

	return res
}

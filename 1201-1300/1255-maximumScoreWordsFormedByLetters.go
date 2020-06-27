package main

func main() {
	println(maxScoreWords([]string{"dog", "cat", "dad", "good"},
		[]byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'},
		[]int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
		23)
	println(maxScoreWords([]string{"xxxz", "ax", "bx", "cx"},
		[]byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'},
		[]int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}),
		27)
	println(maxScoreWords([]string{"ecddd", "eadc", "abded", "decbe", "abec", "ddbea"},
		[]byte{'a', 'b', 'b', 'b', 'b', 'c', 'c', 'c', 'c', 'd', 'd', 'd', 'e', 'e', 'e', 'e', 'e'},
		[]int{10, 10, 9, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
		60)
}

/*
https://leetcode.com/problems/maximum-score-words-formed-by-letters/

Given a list of words, list of  single letters (might be repeating) and score of every character.

Return the maximum score of any valid set of words formed by using the given letters (words[i] cannot
be used two or more times).

It is not necessary to use all characters in letters and each letter can only be used once. Score of letters
'a', 'b', 'c', ... ,'z' is given by score[0], score[1], ... , score[25] respectively.

Example 1:

Input: words = ["dog","cat","dad","good"], letters = ["a","a","c","d","d","d","g","o","o"],
score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
Output: 23
Explanation:
Score  a=1, c=9, d=5, g=3, o=2
Given letters, we can form the words "dad" (5+1+5) and "good" (3+2+2+5) with a score of 23.
Words "dad" and "dog" only get a score of 21.

Example 2:

Input: words = ["xxxz","ax","bx","cx"], letters = ["z","a","b","c","x","x","x"],
score = [4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10]
Output: 27
Explanation:
Score  a=4, b=4, c=4, x=5, z=10
Given letters, we can form the words "ax" (4+5), "bx" (4+5) and "cx" (4+5) with a score of 27.
Word "xxxz" only get a score of 25.

Example 3:

Input: words = ["leetcode"], letters = ["l","e","t","c","o","d"],
score = [0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0]
Output: 0
Explanation:
Letter "e" can only be used once.

Constraints:

1 <= words.length <= 14
1 <= words[i].length <= 15
1 <= letters.length <= 100
letters[i].length == 1
score.length == 26
0 <= score[i] <= 10
words[i], letters[i] contains only lower case English letters.
*/

func maxScoreWords(words []string, letters []byte, score []int) int {
	lettersMap := make(map[byte]int)
	for _, l := range letters {
		v, _ := lettersMap[l]
		v++
		lettersMap[l] = v
	}

	return maxScoreWordsBackTrack(words, lettersMap, score, 0)
}

func maxScoreWordsBackTrack(words []string, letters map[byte]int, score []int, currentScore int) int {
	newScore := currentScore
	for idx, w := range words {
		canMakeWord := true
		cLetters := make(map[byte]int)
		cScore := 0
		for i := 0; i < len(w); i++ {
			v, _ := cLetters[w[i]]
			v++
			cLetters[w[i]] = v

			vl, _ := letters[w[i]]
			if vl < v {
				canMakeWord = false
			}
			cScore += score[w[i]-'a']
		}

		if !canMakeWord {
			continue
		}

		modifyLetters(letters, cLetters, -1)

		possibleScore := maxScoreWordsBackTrack(words[idx+1:], letters, score, currentScore+cScore)
		if possibleScore > newScore {
			newScore = possibleScore
		}

		modifyLetters(letters, cLetters, 1)
	}

	return newScore
}

func modifyLetters(letters, cLetters map[byte]int, dir int) {
	for l, vl := range cLetters {
		v, _ := letters[l]
		v += dir * vl
		letters[l] = v
	}
}

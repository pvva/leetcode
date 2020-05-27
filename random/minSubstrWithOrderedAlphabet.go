package main

func main() {
	println(minSubstringWithOrderedAlphabet("aaaaaaaaab", []byte{'a', 'b'}), "ab")
	println(minSubstringWithOrderedAlphabet("bbbbbbbbac", []byte{'a', 'b', 'c'}), "<empty>")
	println(minSubstringWithOrderedAlphabet("aaaabacb", []byte{'a', 'b', 'c'}), "abac")
	println(minSubstringWithOrderedAlphabet("aaaabacbabc", []byte{'a', 'b', 'c'}), "abc")
	println(minSubstringWithOrderedAlphabet("acbacbacbacbacbabc", []byte{'a', 'b', 'c'}), "abc")
}

/*
Given text as string and an alphabet of characters find minimum substring where all characters of alphabet are ordered.
You can assume that alphabet that is given to the function is already ordered in a required way.
*/

// running time is O(n*l) where n is the size of the text and l is the size of the alphabet
func minSubstringWithOrderedAlphabet(text string, alphabet []byte) string {
	si := -1
	ei := 0

	minSi := -1
	minEi := len(text)

	for ei < len(text) {
		si = findNextStart(text, si+1, alphabet[0])
		if si == -1 {
			break
		}

		ei = findNextSubstringWithCompleteAlphabet(text, si, len(alphabet))
		if ei == -1 {
			break
		}

		if !isAlphabetInOrder(text[si:ei], alphabet) {
			continue
		}

		if minEi-minSi > ei-si {
			minSi = si
			minEi = ei
		}
	}

	if minSi != -1 {
		return text[minSi:minEi]
	}

	return ""
}

func findNextStart(s string, i int, firstChar byte) int {
	for i < len(s) {
		if s[i] == firstChar {
			for i < len(s)-1 && s[i+1] == firstChar {

				i++
			}
			break
		}
		i++
	}

	return i
}

func findNextSubstringWithCompleteAlphabet(s string, i int, alphabetSize int) int {
	seenLetters := make(map[byte]bool)

	for i < len(s) && len(seenLetters) < alphabetSize {
		seenLetters[s[i]] = true
		i++
	}

	if len(seenLetters) < alphabetSize {
		return -1
	}

	return i
}

func isAlphabetInOrder(s string, alphabet []byte) bool {
	cIdx := 0

	for i := range s {
		if s[i] == alphabet[cIdx] {
			cIdx++
		}

		i++
	}
	return cIdx == len(alphabet)
}

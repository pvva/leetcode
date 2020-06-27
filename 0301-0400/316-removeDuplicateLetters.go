package main

func main() {
	println(removeDuplicateLetters("bcabc"), "abc")
	println(removeDuplicateLetters("cbacdcbc"), "acdb")
	println(removeDuplicateLetters("abacb"), "abc")
	println(removeDuplicateLetters("bbcaac"), "bac")
	println(removeDuplicateLetters("ccacbaba"), "acb")
	println(removeDuplicateLetters("bbbacacca"), "bac")
}

/*
https://leetcode.com/problems/remove-duplicate-letters/

Given a string which contains only lowercase letters, remove duplicate letters so that every letter appears once and
only once. You must make sure your result is the smallest in lexicographical order among all possible results.

Example 1:

Input: "bcabc"
Output: "abc"

Example 2:

Input: "cbacdcbc"
Output: "acdb"
*/

// the idea is the following:
// - have number of letters amount in original string, since there can be only lowercase letters, we only need 26 positions
// - have "seen this letter" sign, again 26 positions
// - keep result in increasing order using monotonous queue approach
// - if the last letter from the result cannot be removed (it has current count of 1), then keep it
// - if seen already current letter - skip it
// - otherwise add current letter to the end of the result
// my original implementation was using map instead of [26] array, I improved it after checking other solutions :)
func removeDuplicateLetters(s string) string {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	seen := [26]bool{}

	result := make([]byte, 26)
	si := -1

	for _, r := range s {
		c := byte(r)
		if seen[c-'a'] {
			cnt[c-'a']--
			continue
		}

		for si >= 0 && c < result[si] && cnt[result[si]-'a'] > 1 {
			cnt[result[si]-'a']--
			seen[result[si]-'a'] = false
			si--
		}

		si++
		result[si] = c
		seen[c-'a'] = true
	}

	return string(result[:si+1])
}

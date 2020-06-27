package main

func main() {
	println(longestDupSubstring("banana"), "ana")
}

/*
https://leetcode.com/problems/longest-duplicate-substring/

Given a string S, consider all duplicated substrings: (contiguous) substrings of S that occur 2 or more times.
(The occurrences may overlap.)

Return any duplicated substring that has the longest possible length.  (If S does not have a duplicated substring,
the answer is "".)

Example 1:

Input: "banana"
Output: "ana"

Example 2:

Input: "abcd"
Output: ""

Note:

2 <= S.length <= 10^5
S consists of lowercase English letters.
*/

func longestDupSubstring(S string) string {
	l, r := 0, len(S)-1
	longest := ""

	for l <= r {
		m := (l + r) / 2
		res := findDuplicates(S, m)
		if res == "" {
			r = m - 1
		} else {
			if len(res) > len(longest) {
				longest = res
			}
			l = m + 1
		}
	}

	return longest
}

func findDuplicates(s string, length int) string {
	seenParts := make(map[int64][]int)

	l := length
	if l > 12 {
		l = 12
	}
	mask := int64(1<<(l*5) - 1)

	h := hash(s[0:length], mask)
	seenParts[h] = []int{0}

	for i := 1; i <= len(s)-length; i++ {
		h = ((h << 5) | int64(s[i+length-1]-'a')) & mask

		if _, ex := seenParts[h]; ex {
			t := s[i : i+length]
			for _, j := range seenParts[h] {
				if s[j:j+length] == t {
					return t
				}
			}
			seenParts[h] = append(seenParts[h], i)
		} else {
			seenParts[h] = []int{i}
		}
	}

	return ""
}

func hash(s string, mask int64) int64 {
	h := int64(0)

	for _, r := range s {
		h = ((h << 5) | int64(r-'a')) & mask
	}

	return h
}

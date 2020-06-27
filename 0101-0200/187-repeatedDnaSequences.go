package main

import (
	"fmt"
)

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"), []string{"AAAAACCCCC", "CCCCCAAAAA"})
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"), []string{"AAAAAAAAAA"})
}

/*
https://leetcode.com/problems/repeated-dna-sequences/

All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T, for example: "ACGAATTCCG".
When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.

Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

Example:

Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"

Output: ["AAAAACCCCC", "CCCCCAAAAA"]
*/

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}

	mask := 1<<20 - 1

	temp := make(map[int]bool)
	results := make(map[string]bool)

	hash := strHash(s[0:10])
	temp[hash] = true

	for i := 1; i < len(s)-9; i++ {
		hash = ((hash << 2) & mask) | singleCharHash(s[i+9])

		if temp[hash] {
			results[s[i:i+10]] = true
		}
		temp[hash] = true
	}
	out := make([]string, len(results))
	idx := 0
	for str := range results {
		out[idx] = str
		idx++
	}

	return out
}

func strHash(s string) int {
	h := singleCharHash(s[0])

	for i := 1; i < len(s); i++ {
		h <<= 2
		h |= singleCharHash(s[i])
	}

	return h
}

func singleCharHash(c byte) int {
	switch c {
	case 'C':
		return 1
	case 'G':
		return 2
	case 'T':
		return 3
	}

	return 0
}

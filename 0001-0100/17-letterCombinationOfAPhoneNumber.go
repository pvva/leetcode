package main

import "fmt"

func main() {
	fmt.Println([]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	fmt.Println([]string{}, letterCombinations(""))
	fmt.Println([]string{"a", "b", "c"}, letterCombinations("2"))
}

/*
https://leetcode.com/problems/letter-combinations-of-a-phone-number/

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could
represent. Return the answer in any order.

Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:

Input: digits = ""
Output: []

Example 3:

Input: digits = "2"
Output: ["a","b","c"]


Constraints:

0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].
*/

var numbersToLetters = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinationsRec(digits string, pos int, prefixes []string) []string {
	if pos == len(digits) {
		return prefixes
	}
	res := []string{}
	if pos == 0 {
		res = numbersToLetters[digits[pos:pos+1]]
	} else {
		for _, l := range numbersToLetters[digits[pos:pos+1]] {
			for _, p := range prefixes {
				res = append(res, p+l)
			}
		}
	}

	return letterCombinationsRec(digits, pos+1, res)
}

func letterCombinations(digits string) []string {
	return letterCombinationsRec(digits, 0, []string{})
}

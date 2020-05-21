package main

import "bytes"

func main() {
	println(reverseWords("the sky is blue"), "blue is sky the")
	println(reverseWords("  hello world!  "), "world! hello")
	println(reverseWords("a good   example"), "example good a")
}

/*
https://leetcode.com/problems/reverse-words-in-a-string/

Given an input string, reverse the string word by word.

Example 1:

Input: "the sky is blue"
Output: "blue is sky the"

Example 2:

Input: "  hello world!  "
Output: "world! hello"
Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:

Input: "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.


Note:

A word is defined as a sequence of non-space characters.
Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.
You need to reduce multiple spaces between two words to a single space in the reversed string.
*/

func reverseWords(s string) string {
	result := bytes.Buffer{}
	currentSpace := ""

	idx := len(s) - 1
	for idx >= 0 {
		c := s[idx]
		idx--

		if c == ' ' {
			continue
		} else {
			end := idx + 2
			for idx >= 0 && s[idx] != ' ' {
				idx--
			}

			result.WriteString(currentSpace)
			result.WriteString(s[idx+1 : end])
			currentSpace = " "
		}
	}

	return result.String()
}

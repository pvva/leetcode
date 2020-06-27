package main

import "fmt"

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("()((()"))
	fmt.Println(removeInvalidParentheses("(a)())()"))
	fmt.Println(removeInvalidParentheses("(()"))
	fmt.Println(removeInvalidParentheses(")d))"))

}

/*
https://leetcode.com/problems/remove-invalid-parentheses/

Remove the minimum number of invalid parentheses in order to make the input string valid. Return all possible results.

Note: The input string may contain letters other than the parentheses ( and ).

Example 1:

Input: "()())()"
Output: ["()()()", "(())()"]
Example 2:

Input: "(a)())()"
Output: ["(a)()()", "(a())()"]
Example 3:

Input: ")("
Output: [""]
*/

func removeInvalidParentheses(s string) []string {
	res := make(map[string]bool)
	extract(s, res, 0, []byte{'(', ')'})

	r := make([]string, len(res))
	c := 0
	for k := range res {
		r[c] = k
		c++
	}

	return r
}

func extract(s string, rs map[string]bool, start int, seq []byte) {
	stack := 0

	for i := start; i < len(s); i++ {
		if s[i] == seq[0] {
			stack++
		} else if s[i] == seq[1] {
			stack--
		}

		if stack >= 0 {
			continue
		}

		if i == 0 {
			extract(s[1:], rs, i, seq)
		} else {
			for j := 0; j <= i; j++ {
				if s[j] == seq[1] && (j == i || j+1 == len(s) || s[j+1] != seq[1]) {
					extract(s[:j]+s[j+1:], rs, i, seq)
				}
			}
		}

		return
	}

	r := reverse(s)
	if seq[0] == '(' {
		extract(r, rs, 0, []byte{')', '('})
	} else {
		rs[r] = true
	}
}

func reverse(s string) string {
	l := len(s)
	r := make([]byte, l, l)
	for i := 0; i < l; i++ {
		r[i] = s[l-i-1]
	}

	return string(r)
}

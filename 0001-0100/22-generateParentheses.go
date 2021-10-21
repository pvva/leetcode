package main

import "fmt"

func main() {
	fmt.Println([]string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	fmt.Println(generateParenthesis(3))
	fmt.Println([]string{"()"})
	fmt.Println(generateParenthesis(1))
	fmt.Println([]string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"})
	fmt.Println(generateParenthesis(4))
}

/*
https://leetcode.com/problems/generate-parentheses/

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:

Input: n = 1
Output: ["()"]

Constraints:

1 <= n <= 8
*/

func generateParenthesisRec(o, c int, state []string) []string {
	if o == 0 && c == 0 {
		return state
	}
	var t []string
	var res []string
	if o > 0 {
		if len(state) > 0 {
			t = make([]string, len(state))
			copy(t, state)
			for i := range t {
				t[i] += "("
			}
		} else {
			t = []string{"("}
		}
		res = generateParenthesisRec(o-1, c, t)
	}
	if c > o {
		t = make([]string, len(state))
		copy(t, state)
		for i := range t {
			t[i] += ")"
		}
		res = append(res, generateParenthesisRec(o, c-1, t)...)
	}

	return res
}

func generateParenthesis(n int) []string {
	return generateParenthesisRec(n, n, []string{})
}

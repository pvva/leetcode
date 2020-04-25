package main

func main() {
	println(longestValidParentheses("(()"), 2)
	println(longestValidParentheses(")()())"), 4)
	println(longestValidParentheses(")("), 0)
	println(longestValidParentheses("())"), 2)
	println(longestValidParentheses("()(()"), 2)
	println(longestValidParentheses("(()(((()"), 2)
	println(longestValidParentheses("))))((()(("), 2)
	println(longestValidParentheses("(()()"), 4)
}

/*
https://leetcode.com/problems/longest-valid-parentheses/

Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed)
parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
*/

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}

	stack := []int{}
	for i := 0; i < len(s); i++ {
		l := len(stack) - 1
		if s[i] == '(' {
			stack = append(stack, i)
		} else if l >= 0 {
			if s[stack[l]] == '(' {
				stack = stack[:l]
			} else {
				stack = append(stack, i)
			}
		} else {
			stack = append(stack, i)
		}
	}
	if len(stack) == 0 {
		return len(s)
	}

	p := 0
	c := -1
	for len(stack) > 0 {
		t := stack[0] - c
		if p < t-1 {
			p = t - 1
		}
		c = stack[0]
		stack = stack[1:]
	}
	t := len(s) - c
	if p < t-1 {
		p = t - 1
	}

	return p
}

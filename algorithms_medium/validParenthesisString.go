package main

func main() {
	println(checkValidString("()"), true)
	println(checkValidString("(*)"), true)
	println(checkValidString("(*))"), true)
}

/*
Given a string containing only three types of characters: '(', ')' and '*', write a function to check
whether this string is valid. We define the validity of a string by these rules:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
An empty string is also valid.

Example 1:
Input: "()"
Output: True

Example 2:
Input: "(*)"
Output: True

Example 3:
Input: "(*))"
Output: True
*/

func checkValidString(s string) bool {
	openCount := 0
	possibleClosingCount := 0

	for _, c := range s {
		if c == '(' {
			openCount++
		} else if c == ')' {
			openCount--
		} else if c == '*' {
			possibleClosingCount++
			if openCount > 0 {
				openCount--
				possibleClosingCount++
			}
		}

		if openCount < 0 {
			if possibleClosingCount > 0 {
				possibleClosingCount--
				openCount++
			} else {
				return false
			}
		}
	}

	return openCount == 0
}

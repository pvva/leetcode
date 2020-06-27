package main

import "math"

func main() {
	println(myAtoi("42"), 42)
	println(myAtoi("    -42"), -42)
	println(myAtoi("4193 with words"), 4193)
	println(myAtoi("words and 987"), 0)
	println(myAtoi("-91283472332"), -2147483648)
	println(myAtoi("-2147483648"), -2147483648)
	println(myAtoi("-2147483647"), -2147483647)
}

/*
https://leetcode.com/problems/string-to-integer-atoi/

Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found.
Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as
possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect
on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists
because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range:
[−2^31,  2^31 − 1]. If the numerical value is out of the range of representable values, INT_MAX (2^31 − 1) or
INT_MIN (−2^31) is returned.

Example 1:

Input: "42"
Output: 42

Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be performed.
Example 5:

Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−2^31) is returned.

*/

func myAtoi(str string) int {
	sign := int64(0)
	num := 0

	idx := 0
	for idx < len(str) {
		c := str[idx]
		idx++

		isDigit := c >= '0' && c <= '9'

		if sign == 0 {
			if c == ' ' {
				continue
			} else if c == '-' {
				sign = -1
			} else if c == '+' {
				sign = 1
			} else if isDigit {
				sign = 1
				num = int(c - '0')
			} else {
				break
			}
		} else if isDigit {
			t := int64(num*10) + sign*int64(c-'0')
			if t <= math.MinInt32 {
				num = math.MinInt32
				break
			} else if t >= math.MaxInt32 {
				num = math.MaxInt32
				break
			} else {
				num = int(t)
			}
		} else {
			break
		}
	}

	return num
}

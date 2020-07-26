package main

func main() {
	println(addDigits(38), 2)
	println(addDigitsCT(38), 2)
}

/*
https://leetcode.com/problems/add-digits/

Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

Example:

Input: 38
Output: 2
Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
             Since 2 has only one digit, return it.
Follow up:
Could you do it without any loop/recursion in O(1) runtime?
*/

func addDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}

	if sum < 10 {
		return sum
	}

	return addDigits(sum)
}

func addDigitsCT(num int) int {
	if num < 10 {
		return num
	}

	return num % 9
}

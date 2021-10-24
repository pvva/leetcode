package main

func main() {
	println("6", multiply("2", "3"))
	println("56088", multiply("123", "456"))
	println("121932631112635269", multiply("123456789", "987654321"))
	println("891", multiply("9", "99"))
}

/*
https://leetcode.com/problems/multiply-strings/

Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2,
also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.

Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"

Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"

Constraints:

1 <= num1.length, num2.length <= 200
num1 and num2 consist of digits only.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
*/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n1 := make([]byte, len(num1))
	n2 := make([]byte, len(num2))

	s := 0
	for i := len(num1) - 1; i >= 0; i-- {
		n1[s] = num1[i] - '0'
		s++
	}
	s = 0
	for i := len(num2) - 1; i >= 0; i-- {
		n2[s] = num2[i] - '0'
		s++
	}
	l1 := len(n1)
	l2 := len(n2)
	res := make([]byte, l1+l2)
	c := byte(0)
	for i := 0; i < l2; i++ {
		for j := 0; j < l1; j++ {
			r := res[i+j] + n1[j]*n2[i] + c
			c = r / 10
			r %= 10
			res[i+j] = r
		}
		for s = i; c > 0; s++ {
			r := res[s+l1] + c
			c = r / 10
			r %= 10
			res[s+l1] = r
		}
	}
	res[l1+l2-1] += c
	l1 = len(res) - 1
	for res[l1] == 0 {
		res = res[:l1]
		l1--
	}
	l1 = 0
	l2 = len(res) - 1
	for l1 < l2 {
		res[l1], res[l2] = res[l2], res[l1]
		res[l1] += '0'
		res[l2] += '0'
		l1++
		l2--
	}
	if l1 == l2 {
		res[l1] += '0'
	}

	return string(res)
}

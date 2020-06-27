package main

import "math"

func main() {
	println(nextGreaterElement(10698762), 10726689)
	println(nextGreaterElement(12), 21)
	println(nextGreaterElement(21), -1)
}

/*
https://leetcode.com/problems/next-greater-element-iii/

Given a positive 32-bit integer n, you need to find the smallest 32-bit integer which has exactly the same digits
existing in the integer n and is greater in value than n. If no such positive 32-bit integer exists,
you need to return -1.

Example 1:

Input: 12
Output: 21

Example 2:

Input: 21
Output: -1
*/

func nextGreaterElement(n int) int {
	digits := getDigits(n)

	idx := -1
	for i := 1; i < len(digits); i++ {
		if digits[i] < digits[i-1] {
			idx = i
			break
		}
	}
	if idx == -1 {
		return -1
	}
	reverse(digits[:idx])
	for i := idx - 1; i >= 0; i-- {
		if digits[i] > digits[idx] {
			digits[i], digits[idx] = digits[idx], digits[i]
			break
		}
	}

	return toNumber(digits)
}

func getDigits(n int) []byte {
	digits := []byte{}
	for n > 0 {
		t := n / 10
		digits = append(digits, byte(n-t*10))
		n = t
	}

	return digits
}

func toNumber(digits []byte) int {
	n := int64(0)
	for i := len(digits) - 1; i >= 0; i-- {
		n = 10*n + int64(digits[i])
	}
	if n > math.MaxInt32 {
		return -1
	}

	return int(n)
}

func reverse(arr []byte) {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

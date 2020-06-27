package main

func main() {
	println(removeKdigits("1432219", 3), "1219")
	println(removeKdigits("1234567890", 9), "0")
	println(removeKdigits("10200", 1), "200")
	println(removeKdigits("10", 2), "0")
}

/*
https://leetcode.com/problems/remove-k-digits/

Given a non-negative integer num represented as a string, remove k digits from the number so that the new number
is the smallest possible.

Note:
The length of num is less than 10002 and will be ≥ k.
The given num does not contain any leading zero.

Example 1:

Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.

Example 2:

Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
Example 3:

Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.
*/

// main idea here is to use stack
// push first digit to the stack, then for every next digit do the following:
// - if next digit is less then the one on stack and we can drop digits - pop from stack
// - if stack is empty and we have digit 0 - skip it
// - otherwise push digit onto stack
// in the end if we still have digits to drop, pop them from stack
// neat idea for the implementation is to use array of digits as stack, keeping just an index that points to stack head
func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	if k == 0 {
		return num
	}

	result := make([]byte, len(num))
	rIdx := 1 // next place for digit
	result[0] = num[0]

	for i := 1; i < len(num); i++ {
		for k > 0 && rIdx > 0 && result[rIdx-1] > num[i] {
			rIdx--
			k--
		}

		if rIdx == 0 && num[i] == '0' {
			continue
		}

		result[rIdx] = num[i]
		rIdx++
	}

	for k > 0 {
		rIdx--
		k--
	}

	if rIdx == 0 {
		return "0"
	}

	return string(result[:rIdx])
}

package main

func main() {
	println(addBinary("11", "1"), "100")
	println(addBinary("1010", "1011"), "10101")
	println(addBinary("110010", "10111"), "1001001")
	println(addBinary("1111", "1111"), "11110")
	println(addBinary("101", "111"), "1100")
	println(addBinary("111", "111"), "1110")
}

/*
https://leetcode.com/problems/add-binary/

Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"


Constraints:

Each string consists only of '0' or '1' characters.
1 <= a.length, b.length <= 10^4
Each string is either "0" or doesn't contain any leading zero.
*/

func addBinary(a string, b string) string {
	var t *string
	var res []byte
	if len(a) > len(b) {
		res = []byte(a)
		t = &b
	} else {
		res = []byte(b)
		t = &a
	}

	c := byte(0)
	for i, l := len(res)-1, len(*t)-1; l >= 0; i-- {
		rv := res[i] - '0'
		v := (*t)[l] - '0'
		l--
		r := rv + v + c
		c = (r & 2) >> 1
		res[i] = r&1 + '0'
	}

	for i := len(res) - len(*t) - 1; i >= 0 && c > 0; i-- {
		rv := res[i] - '0'
		res[i] = (rv ^ c) + '0'
		c &= rv
	}

	if c > 0 {
		return "1" + string(res)
	}

	return string(res)
}

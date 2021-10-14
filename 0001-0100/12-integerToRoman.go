package main

func main() {
	for _, n := range []int{3, 4, 8, 9, 12, 18, 27, 49, 449, 949, 999, 1998, 2021} {
		println(n, intToRoman(n), intToRoman2(n))
	}
}

/*
https://leetcode.com/problems/integer-to-roman/

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply
X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.
Instead, the number four is written as IV. Because the one is before the five we subtract it making four.
The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral.
*/

var intToRomanValues = []struct {
	v int
	r string
}{
	{v: 1000, r: "M"},
	{v: 500, r: "D"},
	{v: 100, r: "C"},
	{v: 50, r: "L"},
	{v: 10, r: "X"},
	{v: 5, r: "V"},
	{v: 1, r: "I"},
}

func getSpecialRomanNumber(v int) (string, int) {
	r := ""
	rv := v
	if v == 4 {
		r += "IV"
		rv = 0
	} else if v == 9 {
		r += "IX"
		rv = 0
	} else if v >= 40 && v <= 49 {
		r += "XL"
		rv = v - 40
	} else if v >= 90 && v <= 99 {
		r += "XC"
		rv = v - 90
	} else if v >= 400 && v <= 499 {
		r += "CD"
		rv = v - 400
	} else if v >= 900 && v <= 999 {
		r += "CM"
		rv = v - 900
	}
	return r, rv
}

func intToRoman(num int) string {
	r := ""

	for _, s := range intToRomanValues {
		if num == 0 {
			break
		}
		for {
			sn, n := getSpecialRomanNumber(num)
			if sn == "" {
				break
			}
			r += sn
			num = n
		}

		i := num / s.v
		if i == 0 {
			continue
		}
		d := i * s.v
		num -= d
		for ; i > 0; i-- {
			r += s.r
		}
	}

	return r
}

var intToRomanValues2 = []struct {
	v int
	r string
}{
	{v: 1000, r: "M"},
	{v: 900, r: "CM"},
	{v: 500, r: "D"},
	{v: 400, r: "CD"},
	{v: 100, r: "C"},
	{v: 90, r: "XC"},
	{v: 50, r: "L"},
	{v: 40, r: "XL"},
	{v: 10, r: "X"},
	{v: 9, r: "IX"},
	{v: 5, r: "V"},
	{v: 4, r: "IV"},
	{v: 1, r: "I"},
}

func intToRoman2(num int) string {
	r := ""

	for _, s := range intToRomanValues2 {
		for num >= s.v {
			num -= s.v
			r += s.r
		}
	}

	return r
}

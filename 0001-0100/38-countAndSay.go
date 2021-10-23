package main

import "strings"

func main() {
	println("1", countAndSay(1))
	println("1211", countAndSay(4))
}

/*
https://leetcode.com/problems/count-and-say/

The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

countAndSay(1) = "1"
countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into
a different digit string.
To determine how you "say" a digit string, split it into the minimal number of groups so that each group is a
contiguous section all of the same character. Then for each group, say the number of characters, then say
the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.

Given a positive integer n, return the nth term of the count-and-say sequence.

Example 1:

Input: n = 1
Output: "1"
Explanation: This is the base case.

Example 2:

Input: n = 4
Output: "1211"
Explanation:
countAndSay(1) = "1"
countAndSay(2) = say "1" = one 1 = "11"
countAndSay(3) = say "11" = two 1's = "21"
countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"
*/

func countAndSayRec(s string, n int) string {
	if n == 0 {
		return s
	}
	ns := strings.Builder{}
	for i := 0; i < len(s); {
		c := s[i]
		j := i + 1
		for ; j < len(s) && s[j] == c; j++ {
		}
		j -= i
		i += j
		ns.WriteByte(byte(j + '0'))
		ns.WriteByte(c)
	}

	return countAndSayRec(ns.String(), n-1)
}

func countAndSay(n int) string {
	return countAndSayRec("1", n-1)
}

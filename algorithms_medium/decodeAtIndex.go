package main

func main() {
	println(decodeAtIndex("leet2code3", 10), "o")
	println(decodeAtIndex("ha22", 5), "h")
	println(decodeAtIndex("a2345678999999999999999", 1), "a")
	println(decodeAtIndex("czjkk9elaqwiz7s6kgvl4gjixan3ky7jfdg3kyop3husw3fm289thisef8blt7a7zr5v5lhxqpntenvxnmlq7l34ay3jaayikjps", 768077956), "c")
}

/*
https://leetcode.com/problems/decoded-string-at-index/

An encoded string S is given.  To find and write the decoded string to a tape, the encoded string is read one
character at a time and the following steps are taken:

If the character read is a letter, that letter is written onto the tape.
If the character read is a digit (say d), the entire current tape is repeatedly written d-1 more times in total.
Now for some encoded string S, and an index K, find and return the K-th letter (1 indexed) in the decoded string.

Example 1:

Input: S = "leet2code3", K = 10
Output: "o"
Explanation:
The decoded string is "leetleetcodeleetleetcodeleetleetcode".
The 10th letter in the string is "o".

Example 2:

Input: S = "ha22", K = 5
Output: "h"
Explanation:
The decoded string is "hahahaha".  The 5th letter is "h".

Example 3:

Input: S = "a2345678999999999999999", K = 1
Output: "a"
Explanation:
The decoded string is "a" repeated 8301530446056247680 times.  The 1st letter is "a".
*/

func decodeAtIndex(S string, K int) string {
	pos := 0

	for i, c := range S {
		if c >= 'a' && c <= 'z' {
			pos++
			if pos == K {
				return string(c)
			}
		} else {
			t := int(c - '0')
			if pos*t >= K {
				tPos := 0
				for x := 1; x <= t; x++ {
					tPos += pos
					if tPos >= K {
						tPos -= pos
						break
					}
				}
				return decodeAtIndex(S[:i], K-tPos)
			}
			pos *= t
		}
	}

	return ""
}

package main

func main() {
	println(findNthDigit(11), 0)
	println(findNthDigit(10), 1)
	println(findNthDigit(15), 2)
	println(findNthDigit(20), 1)
	println(findNthDigit(23), 6)
}

/*
https://leetcode.com/problems/nth-digit/

Find the nth digit of the infinite integer sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...

Note:
n is positive and will fit within the range of a 32-bit signed integer (n < 231).

Example 1:

Input:
3

Output:
3

Example 2:

Input:
11

Output:
0

Explanation:
The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a 0, which is part of the number 10.
*/

func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}

	seqStart := 1
	seqN := 1

	t := 9 * seqStart * seqN
	for n > t {
		n -= t
		seqStart *= 10
		seqN++

		t = 9 * seqStart * seqN
	}

	target := seqStart + n/seqN - 1 // minus one because we start indexing from zero, and n/seqN is a index shift
	pos := n % seqN
	if pos != 0 {
		// take (seqN-pos)-th digit from target+1
		target++

		for i := 0; i < seqN-pos; i++ {
			target /= 10
		}
	}

	// take last digit
	return target - int(target/10)*10
}

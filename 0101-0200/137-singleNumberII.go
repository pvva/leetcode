package main

func main() {
	println(singleNumber([]int{2, 2, 3, 2}), 3)
	println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}), 99)

	println(singleNumberBitwise([]int{2, 2, 3, 2}), 3)
	println(singleNumberBitwise([]int{0, 1, 0, 1, 0, 1, 99}), 99)
}

/*
https://leetcode.com/problems/single-number-ii/

Given a non-empty array of integers, every element appears three times except for one, which appears exactly once.
Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,3,2]
Output: 3

Example 2:

Input: [0,1,0,1,0,1,99]
Output: 99
*/

func singleNumber(nums []int) int {
	t := make(map[int]int)

	for _, v := range nums {
		t[v]++
	}

	for v, c := range t {
		if c == 1 {
			return v
		}
	}

	return 0
}

// first time the value will appear in singles
// second time encountered it will appear in doubles and will be weed out of singles
// third time encountered it will appear in singles but will be weed out by being in doubles
func singleNumberBitwise(nums []int) int {
	singles := 0
	doubles := 0

	for _, v := range nums {
		// add current value into singles if it is not in doubles
		// exclude from singles if it's there already
		singles = (^doubles) & (singles ^ v)
		// add current value into doubles if it is not in singles
		// exclude from double if it's there already
		doubles = (^singles) & (doubles ^ v)
	}

	return singles
}

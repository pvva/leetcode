package main

func main() {
	println(maxCoins([]int{9}))
	println(maxCoins([]int{3, 1, 5, 8}))
	println(maxCoins([]int{3, 1, 5, 8, 2}))
	println(maxCoins([]int{35, 16, 83, 87, 84, 59, 48, 41}))
	println(maxCoins([]int{9, 76, 64, 21, 97}))
	println(maxCoins([]int{35, 16, 83, 87, 84, 59, 48, 41, 20, 54}))
}

/*
https://leetcode.com/problems/burst-balloons/

Given n balloons, indexed from 0 to n-1. Each balloon is painted with a number on it represented by array nums.
You are asked to burst all the balloons. If you burst balloon i you will get nums[left] * nums[i] * nums[right] coins.
Here left and right are adjacent indices of i. After the burst, the left and right then becomes adjacent.

Find the maximum coins you can collect by bursting the balloons wisely.

Note:

You may imagine nums[-1] = nums[n] = 1. They are not real therefore you can not burst them.
0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
Example:

Input: [3,1,5,8]
Output: 167
Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
             coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
*/

func maxCoins(nums []int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	}

	maxSeq := make([][]int, ln)
	for i := 0; i < ln; i++ {
		maxSeq[i] = make([]int, ln)
	}

	for l := 1; l <= ln; l++ {
		for startIdx := 0; startIdx <= ln-l; startIdx++ {

			endIdx := startIdx + l - 1
			for i := startIdx; i <= endIdx; i++ {
				// get values if anything between start and end indices is burst out except for i'th index
				values := getValues(startIdx-1, i, endIdx+1, nums)
				cm := values[0] * values[1] * values[2]
				if i > 0 {
					// add max of what was burst out between start and i
					cm += maxSeq[startIdx][i-1]
				}
				if i+1 < ln {
					// add max of what was burst out between i and end
					cm += maxSeq[i+1][endIdx]
				}
				maxSeq[startIdx][endIdx] = max(maxSeq[startIdx][endIdx], cm)
			}
		}
	}

	return maxSeq[0][ln-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getValues(sa, sb, sc int, nums []int) [3]int {
	if len(nums) == 1 {
		return [3]int{1, nums[0], 1}
	}

	a := 1
	if sa >= 0 {
		a = nums[sa]
	}
	b := nums[sb]
	c := 1
	if sc < len(nums) {
		c = nums[sc]
	}

	return [3]int{a, b, c}
}

package main

func main() {
	println(maxProduct([]int{2, 3, -2, 4}))
	println(maxProduct([]int{-2, 0, -1}))
}

/*
https://leetcode.com/problems/maximum-product-subarray/

Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

func maxProduct(nums []int) int {
	max := nums[0]
	cMax := nums[0]
	cMin := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] >= 0 {
			cMax = maxI(nums[i], cMax*nums[i])
			cMin = minI(nums[i], cMin*nums[i])
		} else {
			t := cMax
			cMax = maxI(nums[i], cMin*nums[i])
			cMin = minI(nums[i], t*nums[i])
		}

		max = maxI(max, cMax)
	}

	return max
}

func maxI(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minI(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxProductN2(nums []int) int {
	l := len(nums)
	prod := make([]int, l)

	prod[l-1] = nums[l-1]
	max := prod[l-1]

	for i := l - 2; i >= 0; i-- {
		prod[i] = nums[i]
		if max < prod[i] {
			max = prod[i]
		}

		for j := i + 1; j < l; j++ {
			prod[j] *= prod[i]
			if max < prod[j] {
				max = prod[j]
			}
		}
	}

	return max
}

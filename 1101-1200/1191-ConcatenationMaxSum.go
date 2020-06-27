package main

func main() {
	println(kConcatenationMaxSum([]int{1, 2}, 3), 9)
	println(kConcatenationMaxSum([]int{1, -2, 1}, 5), 2)
	println(kConcatenationMaxSum([]int{-1, -2}, 7), 0)
	println(kConcatenationMaxSum([]int{-5, 4, -4, -3, 5, -3}, 3), 5)
	println(kConcatenationMaxSum([]int{-5, -2, 0, 0, 3, 9, -2, -5, 4}, 5), 20)
	println(kConcatenationMaxSum([]int{2, -5, 1, 0, -2, -2, 2}, 2), 4)
}

/*
https://leetcode.com/problems/k-concatenation-maximum-sum/

Given an integer array arr and an integer k, modify the array by repeating it k times.

For example, if arr = [1, 2] and k = 3 then the modified array will be [1, 2, 1, 2, 1, 2].

Return the maximum sub-array sum in the modified array. Note that the length of the sub-array can be 0 and its
sum in that case is 0.

As the answer can be very large, return the answer modulo 10^9 + 7.

Example 1:

Input: arr = [1,2], k = 3
Output: 9

Example 2:

Input: arr = [1,-2,1], k = 5
Output: 2

Example 3:

Input: arr = [-1,-2], k = 7
Output: 0
*/

const mod = 1000000007

func kConcatenationMaxSum(arr []int, k int) int {
	maxSum, allSum := maxSubArray(&arr, 1)
	if k == 1 {
		return maxSum % mod
	}
	if allSum <= 0 {
		max2Sum, _ := maxSubArray(&arr, 2)

		if max2Sum > maxSum {
			return max2Sum % mod
		}
		return maxSum % mod
	}

	return int((int64(allSum)*int64(k-1) + int64(maxSum)) % mod)
}

func maxSubArray(arr *[]int, k int) (int, int) {
	currentSum := 0
	maxSum := 0
	allSum := 0

	for i := 0; i < k; i++ {
		allSum = 0
		for _, v := range *arr {
			allSum += v
			if currentSum < 0 {
				currentSum = v
			} else {
				currentSum += v
			}
			if maxSum < currentSum {
				maxSum = currentSum
			}
		}
	}

	return maxSum, allSum
}

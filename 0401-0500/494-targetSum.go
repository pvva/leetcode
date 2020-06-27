package main

func main() {
	println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

/*
https://leetcode.com/problems/target-sum/

You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols + and -.
For each integer, you should choose one from + and - as its new symbol.

Find out how many ways to assign symbols to make sum of integers equal to target S.

Example 1:
Input: nums is [1, 1, 1, 1, 1], S is 3.
Output: 5
Explanation:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

There are 5 ways to assign symbols to make the sum of nums be target 3.
Note:
The length of the given array is positive and will not exceed 20.
The sum of elements in the given array will not exceed 1000.
Your output answer is guaranteed to be fitted in a 32-bit integer.
*/

func findTargetSumWays(nums []int, S int) int {
	return findTargetSumWaysWithMemo(nums, len(nums)-1, 0, S, make([]map[int]int, len(nums)))
}

func findTargetSumWaysWithMemo(nums []int, idx, currentSum, targetSum int, memo []map[int]int) int {
	if idx < 0 {
		if currentSum == targetSum {
			return 1
		}

		return 0
	}

	if memo[idx] != nil {
		if v, ex := memo[idx][currentSum]; ex {
			return v
		}
	} else {
		memo[idx] = make(map[int]int)
	}

	positives := findTargetSumWaysWithMemo(nums, idx-1, currentSum+nums[idx], targetSum, memo)
	negatives := findTargetSumWaysWithMemo(nums, idx-1, currentSum-nums[idx], targetSum, memo)

	memo[idx][currentSum] = positives + negatives

	return positives + negatives
}

package main

func main() {
	println(rob([]int{2, 3, 2}), 3)
	println(rob([]int{1, 2, 3, 1}), 4)
}

/*
https://leetcode.com/problems/house-robber-ii/

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed.
All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one.
Meanwhile, adjacent houses have security system connected and it will automatically contact the police if two adjacent
houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount
of money you can rob tonight without alerting the police.

Example 1:

Input: [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
             because they are adjacent houses.

Example 2:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
*/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	// run calculation 2 time to exclude adjacent houses (first and last)
	p1 := robLimited(nums, 0, len(nums)-1)
	p2 := robLimited(nums, 1, len(nums))

	if p1 > p2 {
		return p1
	}
	return p2
}

// the idea for this iteration is the following:
// - we have sum at previous step (p1) and pre previous step (p2)
// - calculate num[i] + p2, as we cannot take adjacent houses
// - prepare p1 and p2 for next iteration, p1 will become p2 (so p2 <- p1),
//   and p1 will be max(p1, num[i] + p2)
func robLimited(nums []int, from, to int) int {
	prev2StepsSum := 0
	prev1StepSum := 0

	for i := from; i < to; i++ {
		t := prev2StepsSum + nums[i]
		prev2StepsSum = prev1StepSum
		if t > prev1StepSum {
			prev1StepSum = t
		}
	}

	return prev1StepSum
}

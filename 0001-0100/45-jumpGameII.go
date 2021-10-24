package main

func main() {
	println(2, jump([]int{2, 3, 1, 1, 4}))
	println(2, jump([]int{2, 3, 0, 1, 4}))
	println(1, jump([]int{4, 1, 2, 3, 1}))
	println(1, jump([]int{1, 2}))
}

/*
https://leetcode.com/problems/jump-game-ii/

Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Your goal is to reach the last index in the minimum number of jumps.
You can assume that you can always reach the last index.

Example 1:

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps
to the last index.

Example 2:

Input: nums = [2,3,0,1,4]
Output: 2

Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 1000
*/

// dynamic programming
func jump(nums []int) int {
	t := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		t[i] = 1<<31 - 1
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			if t[i+j] > t[i]+1 {
				t[i+j] = t[i] + 1
			}
		}
	}

	return t[len(nums)-1]
}

// greedy
func jump2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	jumps := 1
	farthest := nums[0]
	steps := farthest
	idx := 0

	for steps >= 0 {
		if idx == len(nums)-1 {
			break
		}
		next := idx + nums[idx]
		if farthest < next {
			farthest = next
		}

		if steps == 0 {
			steps = farthest - idx
			jumps++
		} else {
			steps--
			idx++
		}
	}

	return jumps
}

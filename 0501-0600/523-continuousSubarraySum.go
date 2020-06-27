package main

func main() {
	println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6), true)
	println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 6), true)

	println(checkSubarraySumBetter([]int{1, 1}, 2), true)
}

/*
https://leetcode.com/problems/continuous-subarray-sum/

Given a list of non-negative numbers and a target integer k, write a function to check if the array has a continuous
subarray of size at least 2 that sums up to a multiple of k, that is, sums up to n*k where n is also an integer.

Example 1:

Input: [23, 2, 4, 6, 7],  k=6
Output: True
Explanation: Because [2, 4] is a continuous subarray of size 2 and sums up to 6.

Example 2:

Input: [23, 2, 6, 4, 7],  k=6
Output: True
Explanation: Because [23, 2, 6, 4, 7] is an continuous subarray of size 5 and sums up to 42.
*/

// O(n^2)
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := nums[0] + nums[1]
	if sum%k == 0 {
		return true
	}

	partialSums := make([]int, len(nums)-2)
	for i := 2; i < len(nums); i++ {
		sum += nums[i]
		if sum%k == 0 {
			return true
		}
		partialSums[i-2] = sum
	}

	for i := range partialSums {
		for j := i; j < len(partialSums); j++ {
			partialSums[j] -= nums[i]
			if partialSums[j]%k == 0 {
				return true
			}
		}
	}

	return false
}

func checkSubarraySumBetter(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	if k < 0 {
		k = -k
	}
	if k == 1 {
		return true
	}
	mods := make(map[int]int)
	sum := 0
	for i, v := range nums {
		if v == 0 {
			if i > 0 && nums[i-1] == 0 {
				return true
			}
		}
		sum += v
		mod := sum
		if k > 0 {
			mod = mod % k
		}
		if mod == 0 && sum > 0 && i >= 1 {
			return true
		}
		if sameModIdx, ex := mods[mod]; ex && i-sameModIdx >= 2 {
			return true
		}
		mods[mod] = i
	}

	return false
}

package main

func main() {
	println(findNumberOfLIS([]int{1, 3, 5, 4, 7}), 2)
	println(findNumberOfLIS([]int{2, 2, 2, 2, 2}), 5)
}

/*
https://leetcode.com/problems/number-of-longest-increasing-subsequence/

Given an unsorted array of integers, find the number of longest increasing subsequence.

Example 1:
Input: [1,3,5,4,7]
Output: 2
Explanation: The two longest increasing subsequence are [1, 3, 4, 7] and [1, 3, 5, 7].

Example 2:
Input: [2,2,2,2,2]
Output: 5
Explanation: The length of longest continuous increasing subsequence is 1, and there are 5 subsequences' length is 1, so output 5.
Note: Length of the given array will be not exceed 2000 and the answer is guaranteed to be fit in 32-bit signed int.
*/

func findNumberOfLIS(nums []int) int {
	lisSizesCounts := make([][2]int, len(nums))
	maxLIS := 1

	for i := range lisSizesCounts {
		lisSizesCounts[i] = [2]int{1, 1}
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				currentLIS := lisSizesCounts[j][0] + 1

				if lisSizesCounts[i][0] < currentLIS {
					lisSizesCounts[i][0] = currentLIS
					lisSizesCounts[i][1] = lisSizesCounts[j][1]
				} else if lisSizesCounts[i][0] == currentLIS {
					lisSizesCounts[i][1] += lisSizesCounts[j][1]
				}

				if maxLIS < currentLIS {
					maxLIS = currentLIS
				}
			}
		}
	}

	total := 0

	for _, sc := range lisSizesCounts {
		if sc[0] == maxLIS {
			total += sc[1]
		}
	}

	return total
}

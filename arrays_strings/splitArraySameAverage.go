package main

func main() {
	println(splitArraySameAverage([]int{1, 2, 3, 4, 5, 6, 7, 8}), true)
	println(splitArraySameAverage([]int{1, 1, 2, 2, 2}), false)
	println(splitArraySameAverage([]int{3, 1, 2}), true)
	println(splitArraySameAverage([]int{60, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30}), false)
}

/*
https://leetcode.com/problems/split-array-with-same-average/
In a given integer array A, we must move every element of A to either list B or list C. (B and C initially start empty.)

Return true if and only if after such a move, it is possible that the average value of B is equal to the average value of C,
and B and C are both non-empty.

Example :
Input:
[1,2,3,4,5,6,7,8]
Output: true
Explanation: We can split the array into [1,4,5,8] and [2,3,6,7], and both of them have the average of 4.5.
Note:

The length of A will be in the range [1, 30].
A[i] will be in the range of [0, 10000].\\
*/
func splitArraySameAverage(A []int) bool {
	if len(A) < 2 {
		return false
	}

	sum := 0
	for _, n := range A {
		sum += n
	}
	avg := float32(sum) / float32(len(A))

	for i := 1; i < len(A)>>1+1; i++ {
		targetSum := avg * float32(i)
		if targetSum-float32(int(targetSum)) > 1e-4 {
			continue
		}
		if findCombination(A, 0, i, targetSum) {
			return true
		}
	}

	return false
}

func findCombination(nums []int, start, targetSize int, targetSum float32) bool {
	if targetSum < 0 {
		return false
	}
	if targetSize == 0 {
		return targetSum <= 1e-4
	}

	for i := start; i < len(nums); i++ {
		if findCombination(nums, i+1, targetSize-1, targetSum-float32(nums[i])) {
			return true
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return false
}

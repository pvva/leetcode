package main

func main() {
	println(searchInsert([]int{1, 3, 5, 6}, 5), 2)
	println(searchInsert([]int{1, 3, 5, 6}, 2), 1)
	println(searchInsert([]int{1, 3, 5, 6}, 7), 4)
	println(searchInsert([]int{1, 3, 5, 6}, 0), 0)
}

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the
index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0
*/

func searchInsert(nums []int, target int) int {
	s, e := 0, len(nums)

	for e-s > 0 {
		h := s + (e-s)/2
		if nums[h] == target {
			return h
		} else if nums[h] > target {
			e = h
		} else {
			s = h + 1
		}
	}
	return e
}

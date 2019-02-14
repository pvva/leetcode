package main

import (
	"sort"
)

func main() {
	println(firstMissingPositiveBF([]int{1, 2, 0}))
	println(firstMissingPositiveBF([]int{3, 4, -1, 1}))
	println(firstMissingPositiveBF([]int{7, 8, 9, 11, 12}))

	println(firstMissingPositive([]int{1, 2, 0}))
	println(firstMissingPositive([]int{3, 4, -1, 1}))
	println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	println(firstMissingPositive([]int{1}))
}

//Given an unsorted integer array, find the smallest missing positive integer.
func firstMissingPositiveBF(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	dv := 1
	for _, v := range nums {
		if v <= 0 {
			continue
		}

		if v == dv {
			dv++
		} else if v > dv {
			break
		}
	}

	return dv
}

func firstMissingPositive(nums []int) int {
	l := len(nums)

	for i, v := range nums {
		if v <= 0 {
			nums[i] = l + 1
		}
	}
	for _, v := range nums {
		av := v
		if av < 0 {
			av = -av
		}

		if av <= l && nums[av-1] > 0 {
			nums[av-1] *= -1
		}
	}
	for i, v := range nums {
		if v < 0 {
			continue
		}

		return i + 1
	}

	return l + 1
}

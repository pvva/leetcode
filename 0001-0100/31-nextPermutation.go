package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums, []int{1, 3, 2})
}

/*
https://leetcode.com/problems/next-permutation/

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.
Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

// lexicographically ordered means increasing sequence, so the idea comes out of this fact
// starting from the right we are tracing increasing sequence, until we find a fall
// if fall is not found - just reverse the array, and it will be sorted in ascending order
// if fall is found, then we don't care about left part, because we will get correct order in the right part, which
// will give us next greater permutation at first reverse the right part, so that when we are looking from the left,
// it's increasing then find the number which is larger than the fall and swap them, this way the condition is met.
func nextPermutation(nums []int) {
	idx := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i+1] > nums[i] {
			idx = i
			break
		}
	}
	if idx == -1 {
		reverse(nums)
	} else {
		reverse(nums[idx+1:])
		for i := idx + 1; i < len(nums); i++ {
			if nums[i] > nums[idx] {
				nums[i], nums[idx] = nums[idx], nums[i]
				break
			}
		}
	}
}

func reverse(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

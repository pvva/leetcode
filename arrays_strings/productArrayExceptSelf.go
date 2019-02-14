package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{2, 4, 6, 8}))
}

// Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
// Note: Please solve it without division and in O(n).

func productExceptSelf(nums []int) []int {
	l := len(nums)

	out := make([]int, l)

	// all products except for last element
	out[0] = 1
	for i := 1; i < l; i++ {
		out[i] = out[i-1] * nums[i-1]
	}

	// all products backward with skipping current element
	product := nums[l-1]
	for i := l - 2; i >= 0; i-- {
		out[i] *= product
		product *= nums[i]
	}

	return out
}

package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(countSmallerBF([]int{5, 2, 6, 1}))
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
	//fmt.Println(countSmaller([]int{-1, -1, -1}))
}

// You are given an integer array nums and you have to return a new counts array.
// The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].
func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	// use binary index tree, but instead of storing sum of array, store how many bigger numbers in array are
	// to the right of a given element
	// this means that binary index tree must be built with larger elements closer to root
	// also in order to properly work with binary index tree we need to make sure that only positive numbers are in array
	res := make([]int, len(nums))

	// max is needed for binary index tree array representation size
	max := math.MinInt32
	min := math.MaxInt32
	for _, v := range nums {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	max = max - min + 1
	// in place replace
	for i, v := range nums {
		// +1 to achieve positive numbers (min - min = 0)
		nums[i] = v - min + 1
	}
	bit := make([]int, max+1)
	for i := len(nums) - 1; i >= 0; i-- {
		// get all counts starting from preceding int, because of task condition
		// it it would have been "number of smaller or equal elements", then get would be from nums[i]
		res[i] = getBitValue(nums[i]-1, bit)
		updateBitValue(nums[i], bit)
	}

	return res
}

func getBitValue(v int, bit []int) int {
	s := 0
	for v > 0 {
		s += bit[v]
		v -= v & (-v)
	}

	return s
}

func updateBitValue(v int, bit []int) {
	for v < len(bit) {
		bit[v]++
		v += v & (-v)
	}
}

func countSmallerBF(nums []int) []int {
	l := len(nums)
	res := make([]int, l)

	for i := l - 2; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			if nums[i] > nums[j] {
				res[i]++
			}
		}
	}

	return res
}

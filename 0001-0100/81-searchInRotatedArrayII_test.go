package main

import (
	"math/rand"
	"sort"
	"testing"
)

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return true
		} else if nums[m] < nums[l] && nums[m] < nums[r] {
			// pivot is between l and m
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else if nums[m] > nums[l] && nums[m] > nums[r] {
			// pivot is between m and r
			if target < nums[m] && target >= nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] > nums[l] && nums[m] < nums[r] {
			// no pivot
			if target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[l] == target {
				return true
			}
			l++
		}
	}

	return false
}

func search2(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return true
		}
		if nums[l] < nums[m] || nums[m] > nums[r] {
			if target > nums[m] || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else if nums[m] < nums[r] || nums[m] < nums[l] {
			if target > nums[r] || target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			r--
		}
	}
	return false
}

func generateRandomSortedAndPivotedArray(size int) []int {
	//rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(size / 2)
	}
	sort.Ints(arr)

	pivot := rand.Intn(size)
	t := make([]int, pivot)
	copy(t, arr[len(arr)-pivot:])
	copy(arr[pivot:], arr[:len(arr)-pivot])
	copy(arr, t)

	return arr
}

var result bool

func BenchmarkV1(b *testing.B) {
	arr := generateRandomSortedAndPivotedArray(1e6)
	for i := 0; i < b.N; i++ {
		t := rand.Intn(len(arr))
		result = search(arr, t)
	}
}

func BenchmarkV2(b *testing.B) {
	arr := generateRandomSortedAndPivotedArray(1e6)
	for i := 0; i < b.N; i++ {
		t := rand.Intn(len(arr))
		result = search2(arr, t)
	}
}

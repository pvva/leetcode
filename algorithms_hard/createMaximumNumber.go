package main

import "fmt"

func main() {
	fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5), []int{9, 8, 6, 5, 3})
	fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5), []int{6, 7, 6, 0, 4})
	fmt.Println(maxNumber([]int{3, 9}, []int{8, 9}, 3), []int{9, 8, 9})
	fmt.Println(maxNumber([]int{}, []int{1}, 1), []int{1})
	fmt.Println(maxNumber([]int{5, 7, 3}, []int{4, 2, 3}, 3), []int{7, 4, 3})
	fmt.Println(maxNumber([]int{2, 5, 6, 4, 4, 0}, []int{7, 3, 8, 0, 6, 5, 7, 6, 2}, 15),
		[]int{7, 3, 8, 2, 5, 6, 4, 4, 0, 6, 5, 7, 6, 2, 0})
}

/*
https://leetcode.com/problems/create-maximum-number/

Given two arrays of length m and n with digits 0-9 representing two numbers. Create the maximum number of
length k <= m + n from digits of the two. The relative order of the digits from the same array must be preserved.
Return an array of the k digits.

Note: You should try to optimize your time and space complexity.

Example 1:

Input:
nums1 = [3, 4, 6, 5]
nums2 = [9, 1, 2, 5, 8, 3]
k = 5
Output:
[9, 8, 6, 5, 3]

Example 2:

Input:
nums1 = [6, 7]
nums2 = [6, 0, 4]
k = 5
Output:
[6, 7, 6, 0, 4]

Example 3:

Input:
nums1 = [3, 9]
nums2 = [8, 9]
k = 3
Output:
[9, 8, 9]
*/

// time complexity is O(k * len(num1) * len(nums2) * k^2) => O(len(num1) * len(nums2) * k^3)
// explanation: for k iterations find the largest subsequence from both nums1 and nums2,
//   this gives k * len(num1) * len(nums2)
// k^2 comes from merge function, which is executed for every iteration
// if it is possible to find merge function linear to k, then TC would be O(len(num1) * len(nums2) * k^2)
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	maxNum := make([]int, k)
	for i := 0; i <= k; i++ {
		if i > len(nums1) || k-i > len(nums2) {
			continue
		}
		arr1 := findLargestSequence(nums1, i)
		arr2 := findLargestSequence(nums2, k-i)

		arr := merge(arr1, arr2)
		if isLarger(arr, maxNum) {
			maxNum = arr
		}
	}

	return maxNum
}

func findLargestSequence(arr []int, k int) []int {
	result := make([]int, k)
	ri := -1

	for i, v := range arr {
		for ri >= 0 && result[ri] < v && len(arr)-i >= k-ri {
			ri--
		}
		if ri < k-1 {
			ri++
			result[ri] = v
		}
	}

	return result
}

func merge(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	ci1 := 0
	ci2 := 0
	i := 0

	for ; ci1 < len(arr1) && ci2 < len(arr2); i++ {
		if isLarger(arr1[ci1:], arr2[ci2:]) {
			result[i] = arr1[ci1]
			ci1++
		} else {
			result[i] = arr2[ci2]
			ci2++
		}
	}

	if ci1 < len(arr1) {
		copy(result[i:], arr1[ci1:])
	} else {
		copy(result[i:], arr2[ci2:])
	}

	return result
}

func isLarger(arr1, arr2 []int) bool {
	i1 := 0
	i2 := 0
	for i1 < len(arr1) && i2 < len(arr2) && arr1[i1] == arr2[i1] {
		i1++
		i2++
	}
	return i2 == len(arr2) || (i1 < len(arr1) && arr1[i1] > arr2[i2])
}

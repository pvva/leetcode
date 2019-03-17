package main

func main() {
	println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	println(findMedianSortedArrays([]int{}, []int{2, 3}))
}

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2

	res := make([]int, l)

	i1 := 0
	i2 := 0
	ir := 0

	// merge
	for ir < l {
		if i2 == l2 || (i1 < l1 && nums1[i1] < nums2[i2]) {
			res[ir] = nums1[i1]
			i1++
		} else {
			res[ir] = nums2[i2]
			i2++
		}
		ir++
	}

	// median
	if l&1 == 1 {
		return float64(res[l>>1])
	}
	ir = l >> 1

	return (float64(res[ir]) + float64(res[ir-1])) / 2.0
}

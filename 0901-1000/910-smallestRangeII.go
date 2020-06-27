package main

import "sort"

func main() {
	println(smallestRangeII([]int{1}, 0), 1)
	println(smallestRangeII([]int{0, 10}, 2), 6)
	println(smallestRangeII([]int{1, 3, 6}, 3), 3)
}

/*
https://leetcode.com/problems/smallest-range-ii/

Given an array A of integers, for each integer A[i] we need to choose either x = -K or x = K, and add x to A[i] (only once).
After this process, we have some array B.
Return the smallest possible difference between the maximum value of B and the minimum value of B.

Example 1:

Input: A = [1], K = 0
Output: 0
Explanation: B = [1]

Example 2:

Input: A = [0,10], K = 2
Output: 6
Explanation: B = [2,8]

Example 3:

Input: A = [1,3,6], K = 3
Output: 3
Explanation: B = [4,6,3]
*/

// suppose the array is sorted, in this case first X elements will be increased by K, and other elements will be decreased by K
// the task is to find X so that after all changes the difference between max and min in the array is smallest
// at i-th element except for the last one we are trying to find whether increase of i-th element and decrease of next element
//    produce the desired minimum diff of max and min after transformation
// in this case we need to get new possible minimum, which is min(a[0] + k, a[i+1] - k),
//    and new possible maximum, which is max(a[i] + k, a[n] - k)
// we are increasing a[0] because first X elements are increased, and decreasing a[n] because other elements are decreased
// last thing to do is to compare original result, which is a[n] - a[0] and new result which is newMax - newMin
// this algorithm is impossible if array is not sorted
func smallestRangeII(A []int, K int) int {
	sort.Ints(A)

	n := len(A) - 1
	result := A[n] - A[0]

	for i := 0; i < n; i++ {
		incr := A[i] + K
		decr := A[i+1] - K

		min := A[0] + K
		if min > decr {
			min = decr
		}
		max := A[n] - K
		if max < incr {
			max = incr
		}

		r := max - min
		if result > r {
			result = r
		}
	}

	return result
}

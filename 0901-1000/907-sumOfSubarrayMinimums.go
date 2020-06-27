package main

func main() {
	println(sumSubarrayMins([]int{3, 1, 2, 4}), 17)
}

/*
https://leetcode.com/problems/sum-of-subarray-minimums/

Given an array of integers A, find the sum of min(B), where B ranges over every (contiguous) subarray of A.
Since the answer may be large, return the answer modulo 10^9 + 7.

Example 1:

Input: [3,1,2,4]
Output: 17
Explanation: Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].
Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.  Sum is 17.
*/

// let's say we know for each position what is the next lesser element (NLE) on the left and on the right
// then for each position we have (i - left NLE index) sub arrays that end on index i and
//   (right NLE index - i) sub arrays that start on index i
// each sub array that starts before i ends also at some element to the right of i, so to calculate the number of
//   sub arrays where element at index i participates we need to multiply those numbers: (i - left) * (right - i)
// since left and right are indices of LESSER elements, we know that element at index i is the minimum for all
//   sub arrays that contain element at index i, which means that the part of sum would be (i - left) * (right - i) * A[i]
// the idea is to use monotonous increasing stack to find NLE from the left as we go along the array and
//   calculate partial result when we drop elements from the stack
// stack stores indices
// at the end if stack is not empty we need to pop it completely calculating remaining results
func sumSubarrayMins(A []int) int {
	sum := int64(0)

	stack := make([]int, len(A))
	si := -1

	for i, v := range A {
		for si >= 0 && v < A[stack[si]] {
			li := -1
			if si > 0 {
				li = stack[si-1]
			}
			sum += int64((stack[si] - li) * (i - stack[si]) * A[stack[si]])
			si--
		}
		si++
		stack[si] = i
	}
	for si >= 0 {
		li := -1
		if si > 0 {
			li = stack[si-1]
		}
		sum += int64((stack[si] - li) * (len(A) - stack[si]) * A[stack[si]])
		si--
	}

	return int(sum % 1000000007)
}

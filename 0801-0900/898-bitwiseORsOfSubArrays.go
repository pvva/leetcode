package main

func main() {
	println(subarrayBitwiseORs([]int{0}), 1)
	println(subarrayBitwiseORs([]int{1, 1, 2}), 3)
	println(subarrayBitwiseORs([]int{1, 2, 4}), 6)
	println(subarrayBitwiseORs([]int{10, 12}), 3)
	println(subarrayBitwiseORs([]int{13, 4, 2}), 5)
}

/*
https://leetcode.com/problems/bitwise-ors-of-subarrays/

We have an array A of non-negative integers.

For every (contiguous) subarray B = [A[i], A[i+1], ..., A[j]] (with i <= j), we take the bitwise OR of all the elements in B,
obtaining a result A[i] | A[i+1] | ... | A[j].

Return the number of possible results.  (Results that occur more than once are only counted once in the final answer.)

Example 1:

Input: [0]
Output: 1
Explanation:
There is only one possible result: 0.

Example 2:

Input: [1,1,2]
Output: 3
Explanation:
The possible subarrays are [1], [1], [2], [1, 1], [1, 2], [1, 1, 2].
These yield the results 1, 1, 2, 1, 3, 3.
There are 3 unique values, so the answer is 3.
*/

func subarrayBitwiseORs(A []int) int {
	results := make(map[int]bool)

	tSet := make(map[int]bool)
	for i := 0; i < len(A); i++ {
		cSet := make(map[int]bool)

		cSet[A[i]] = true

		for elem := range tSet {
			cSet[elem|A[i]] = true
		}

		for elem := range cSet {
			results[elem] = true
		}

		tSet = cSet
	}

	return len(results)
}

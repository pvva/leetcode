package main

func main() {
	println(shortestSubarray([]int{1}, 1), 1)
	println(shortestSubarray([]int{1, 2}, 4), -1)
	println(shortestSubarray([]int{2, -1, 2}, 3), 3)
	println(shortestSubarray([]int{48, 99, 37, 4, -31}, 140), 2)
	println(shortestSubarray([]int{17, 85, 93, -45, -21}, 150), 2)
	println(shortestSubarray([]int{84, -37, 32, 40, 95}, 167), 3)
	println(shortestSubarray([]int{-28, 81, -20, 28, -29}, 89), 3)
	println(shortestSubarray([]int{-10, 36, 13, 93, 41, -10, 78, 91, 34, -47, -17, 37, 41, 70, 44, 23, 23, 42, 70, 8}, 207), 5)
	println(shortestSubarray([]int{-36, 10, -28, -42, 17, 83, 87, 79, 51, -26, 33, 53, 63, 61, 76, 34, 57, 68, 1, -30}, 484), 9)
}

/*
https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/

Return the length of the shortest, non-empty, contiguous subarray of A with sum at least K.

If there is no non-empty subarray with sum at least K, return -1.

Example 1:

Input: A = [1], K = 1
Output: 1

Example 2:

Input: A = [1,2], K = 4
Output: -1

Example 3:

Input: A = [2,-1,2], K = 3
Output: 3

Note:

1 <= A.length <= 50000
-10 ^ 5 <= A[i] <= 10 ^ 5
1 <= K <= 10 ^ 9
*/

func shortestSubarray(A []int, K int) int {
	prefixSums := make([]int, len(A)+1)
	for i := 0; i < len(A); i++ {
		prefixSums[i+1] = prefixSums[i] + A[i]
	}

	// deque with the rule that prefixSums[j] > prefixSums[i] for j > i
	d := []int{}
	n := len(A) + 1

	for i := 0; i < len(prefixSums); i++ {
		// skip from start if condition is still satisfied
		for len(d) > 0 && prefixSums[i]-prefixSums[d[0]] >= K {
			if n > i-d[0] {
				n = i - d[0]
			}

			d = d[1:]
		}
		// skip from end those sums which are less then current, due to observed negative values
		// these values contribute to increasing the size of corresponding subarray
		l := len(d)
		for l > 0 && prefixSums[i] <= prefixSums[d[l-1]] {
			l--
			d = d[:l]
		}

		d = append(d, i)
	}

	if n == len(A)+1 {
		return -1
	}

	return n
}

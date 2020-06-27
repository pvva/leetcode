package main

func main() {
	a := []int{0, 1, -1}
	b := []int{-1, 1, 0}
	c := []int{0, 0, 1}
	d := []int{-1, 1, 1}

	println(fourSumCount(a, b, c, d))
}

/*
https://leetcode.com/problems/4sum-ii/

Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that
A[i] + B[j] + C[k] + D[l] is zero.

To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500. All integers are in the
range of -228 to 228 - 1 and the result is guaranteed to be at most 231 - 1.

Example:

Input:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

Output:
2

Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/

func fourSumCount(A []int, B []int, C []int, D []int) int {
	result := 0
	hsum := make(map[int]int)

	for aIdx := 0; aIdx < len(A); aIdx++ {
		for bIdx := 0; bIdx < len(B); bIdx++ {
			hs := A[aIdx] + B[bIdx]
			if t, ex := hsum[hs]; ex {
				hsum[hs] = t + 1
			} else {
				hsum[hs] = 1
			}
		}
	}
	for cIdx := 0; cIdx < len(C); cIdx++ {
		for dIdx := 0; dIdx < len(D); dIdx++ {
			hs := -C[cIdx] - D[dIdx]
			if t, ex := hsum[hs]; ex {
				result += t
			}
		}
	}

	return result
}

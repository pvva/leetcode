package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
	fmt.Println(combine(4, 3))
}

/*
https://leetcode.com/problems/combinations/

Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].

You may return the answer in any order.

Example 1:

Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
Example 2:

Input: n = 1, k = 1
Output: [[1]]

Constraints:

1 <= n <= 20
1 <= k <= n
*/

func combineRec(arr []int, k int) [][]int {
	if k == 0 {
		return nil
	}
	res := [][]int{}
	if k == 1 {
		for _, v := range arr {
			res = append(res, []int{v})
		}
		return res
	}

	for i := 0; i < len(arr)-1; i++ {
		t := combineRec(arr[i+1:], k-1)
		for r := range t {
			res = append(res, append(t[r], arr[i]))
		}
	}

	return res
}

func combine(n int, k int) [][]int {
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}

	return combineRec(arr, k)
}

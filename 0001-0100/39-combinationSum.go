package main

import "fmt"

func main() {
	fmt.Println([][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println([][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum([]int{2, 3, 5}, 8))
	fmt.Println([][]int{}, combinationSum([]int{2}, 1))
	fmt.Println([][]int{{1}}, combinationSum([]int{1}, 1))
}

/*
https://leetcode.com/problems/combination-sum/

Given an array of distinct integers candidates and a target integer target, return a list of all
unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if
the frequency of at least one of the chosen numbers is different.

It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations
for the given input.

Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:

Input: candidates = [2], target = 1
Output: []

Example 4:

Input: candidates = [1], target = 1
Output: [[1]]

Example 5:

Input: candidates = [1], target = 2
Output: [[1,1]]

Constraints:

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
All elements of candidates are distinct.
1 <= target <= 500
*/

func combinationSumRec(candidates []int, target, startFrom int) [][]int {
	if target == 0 {
		return nil
	}
	res := [][]int{}
	for i := startFrom; i < len(candidates); i++ {
		v := candidates[i]
		nv := target - v
		if nv < 0 {
			continue
		}
		if nv == 0 {
			res = append(res, []int{v})
			continue
		}
		t := combinationSumRec(candidates, nv, i)
		if t != nil {
			for _, combo := range t {
				combo = append(combo, v)
				res = append(res, combo)
			}
		}
	}

	return res
}

func combinationSum(candidates []int, target int) [][]int {
	return combinationSumRec(candidates, target, 0)
}

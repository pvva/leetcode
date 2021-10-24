package main

import "fmt"

func main() {
	fmt.Println([][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println([][]int{{1, 2, 2}, {5}}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))
	fmt.Println([][]int{{1}}, combinationSum2([]int{1, 1}, 1))
}

/*
https://leetcode.com/problems/combination-sum-ii/

Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in
candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

Example 2:

Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

Constraints:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/

func combinationSum2Rec(candidates [][]int, target, startFrom int, currentSet []int, res *[][]int) {
	if target == 0 {
		t := make([]int, len(currentSet))
		copy(t, currentSet)
		*res = append(*res, t)
		return
	}
	for i := startFrom; i < len(candidates); i++ {
		vFreq := candidates[i]
		if vFreq[1] == 0 {
			continue
		}
		nv := target - vFreq[0]
		if nv < 0 {
			continue
		}

		currentSet = append(currentSet, vFreq[0])
		candidates[i][1]--
		combinationSum2Rec(candidates, nv, i, currentSet, res)
		candidates[i][1]++
		currentSet = currentSet[:len(currentSet)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	candMap := make(map[int]int)
	for _, v := range candidates {
		c, _ := candMap[v]
		candMap[v] = c + 1
	}
	candFreq := make([][]int, len(candMap))
	i := 0
	for v, freq := range candMap {
		candFreq[i] = []int{v, freq}
		i++
	}
	res := [][]int{}
	combinationSum2Rec(candFreq, target, 0, []int{}, &res)

	return res
}

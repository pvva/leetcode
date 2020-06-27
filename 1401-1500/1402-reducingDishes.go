package main

import "sort"

func main() {
	println(maxSatisfaction([]int{-1, -8, 0, 5, -9}), 14)
	println(maxSatisfaction([]int{4, 3, 2}), 20)
	println(maxSatisfaction([]int{-1, -4, -5}), 0)
	println(maxSatisfaction([]int{-2, 5, -1, 0, 3, -3}), 35)
}

/*
https://leetcode.com/problems/reducing-dishes/

A chef has collected data on the satisfaction level of his n dishes. Chef can cook any dish in 1 unit of time.
Like-time coefficient of a dish is defined as the time taken to cook that dish including previous dishes multiplied
by its satisfaction level  i.e.  time[i]*satisfaction[i]
Return the maximum sum of Like-time coefficient that the chef can obtain after dishes preparation.
Dishes can be prepared in any order and the chef can discard some dishes to get this maximum value.

Example 1:

Input: satisfaction = [-1,-8,0,5,-9]
Output: 14
Explanation: After Removing the second and last dish, the maximum total Like-time coefficient will be equal to (-1*1 + 0*2 + 5*3 = 14). Each dish is prepared in one unit of time.
Example 2:

Input: satisfaction = [4,3,2]
Output: 20
Explanation: Dishes can be prepared in any order, (2*1 + 3*2 + 4*3 = 20)
Example 3:

Input: satisfaction = [-1,-4,-5]
Output: 0
Explanation: People don't like the dishes. No dish is prepared.
Example 4:

Input: satisfaction = [-2,5,-1,0,3,-3]
Output: 35
*/

func maxSatisfaction(satisfaction []int) int {
	if len(satisfaction) == 0 {
		return 0
	}

	sort.Ints(satisfaction)

	maxSum := 0

	currentSum := 0
	nextPartialSum := 0
	allNeg := true

	for i, v := range satisfaction {
		currentSum += (i + 1) * v
		nextPartialSum += v
		allNeg = allNeg && v <= 0
	}
	if allNeg {
		return 0
	}
	previousSum := currentSum

	for idx := 0; idx < len(satisfaction) && currentSum >= previousSum; idx++ {
		if currentSum > maxSum {
			maxSum = currentSum
		}

		previousSum = currentSum
		currentSum -= nextPartialSum
		nextPartialSum -= satisfaction[idx]
	}

	return maxSum
}

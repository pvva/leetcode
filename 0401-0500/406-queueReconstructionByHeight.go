package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
	fmt.Println([][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}})
}

/*
https://leetcode.com/problems/queue-reconstruction-by-height/

Suppose you have a random list of people standing in a queue. Each person is described by a pair of integers (h, k),
where h is the height of the person and k is the number of people in front of this person who have a height greater
than or equal to h. Write an algorithm to reconstruct the queue.

Note:
The number of people is less than 1,100.

Example

Input:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

Output:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
*/

func reconstructQueue(people [][]int) [][]int {
	result := [][]int{}

	sort.SliceStable(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	for _, person := range people {
		if person[1] == 0 {
			result = append([][]int{person}, result...)
		} else {
			t := make([][]int, person[1])
			copy(t, result[:person[1]])
			result = append(append(t, person), result[person[1]:]...)
		}
	}

	return result
}

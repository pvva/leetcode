package main

func main() {
	println(5, candy([]int{1, 0, 2}))
	println(4, candy([]int{1, 2, 2}))
	println(17, candy([]int{1, 0, 1, 2, 4, 0, 1, 1, 1}))
	println(7, candy([]int{1, 3, 2, 2, 1}))
	println(6, candy([]int{3, 2, 1}))
	println(18, candy([]int{1, 6, 10, 8, 7, 3, 2}))
	println(11, candy([]int{1, 3, 4, 5, 2}))
	println(13, candy([]int{0, 1, 2, 3, 2, 1}))
}

/*
https://leetcode.com/problems/candy/

There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

You are giving candies to these children subjected to the following requirements:

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
Return the minimum number of candies you need to have to distribute the candies to the children.

Example 1:

Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

Example 2:

Input: ratings = [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
The third child gets 1 candy because it satisfies the above two conditions.

Constraints:

n == ratings.length
1 <= n <= 2 * 10^4
0 <= ratings[i] <= 2 * 10^4
*/

func candy(ratings []int) int {
	res := 1 // first is skipped, keep it initially

	sumUp := 0
	ups := 0
	sumDown := 0
	downs := 0

	prevDir := 0
	currDir := 0
	for i := 1; i < len(ratings); i++ {
		currDir = ratings[i] - ratings[i-1]

		if currDir == 0 || (currDir > 0 && prevDir < 0) {
			// flat surface or valley, handle previously accumulated sums
			res += sumUp + sumDown

			// whichever side count is larger - contributes to total sum
			// this excludes calculating peaks twice and moving peaks only to the longest side
			if ups > downs {
				res += ups
			} else {
				res += downs
			}
			ups = 0
			sumUp = 0
			downs = 0
			sumDown = 0
		}

		if currDir > 0 {
			// increasing
			ups++
			sumUp += ups
		} else if currDir < 0 {
			// decreasing
			downs++
			sumDown += downs
		} else {
			// flat
			res++
		}
		prevDir = currDir
	}
	// take into account sums accumulated to the end
	res += sumUp + sumDown
	if ups > downs {
		res += ups
	} else {
		res += downs
	}

	return res
}

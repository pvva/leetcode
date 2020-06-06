package main

import "sort"

func main() {
	println(twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}), 110)
}

/*
There are 2N people a company is planning to interview. The cost of flying the i-th person to city A is costs[i][0],
and the cost of flying the i-th person to city B is costs[i][1].

Return the minimum cost to fly every person to a city such that exactly N people arrive in each city.

Example 1:

Input: [[10,20],[30,200],[400,50],[30,20]]
Output: 110
Explanation:
The first person goes to city A for a cost of 10.
The second person goes to city A for a cost of 30.
The third person goes to city B for a cost of 50.
The fourth person goes to city B for a cost of 20.

The total minimum cost is 10 + 30 + 50 + 20 = 110 to have half the people interviewing in each city.

Note:

1 <= costs.length <= 100
It is guaranteed that costs.length is even.
1 <= costs[i][0], costs[i][1] <= 1000
*/

// extremely neat idea:
// - put all persons to city A, sum result
// - calculate diff between B and A costs
// - sort diffs ascending
// - add n/2 first diffs to result
// If you want to know who goes where, then you need to do DP (n x n matrix, A goes by x axes and B by y)
func twoCitySchedCost(costs [][]int) int {
	diffBA := make([]int, len(costs))
	sum := 0

	for i, abCosts := range costs {
		sum += abCosts[0]
		diffBA[i] = abCosts[1] - abCosts[0]
	}

	sort.Ints(diffBA)

	for i := 0; i < len(costs)/2; i++ {
		sum += diffBA[i]
	}

	return sum
}

package main

func main() {
	println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1), 200)
	println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0), 500)
	println(findCheapestPrice(4, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1), 6)
	println(findCheapestPrice(5, [][]int{{1, 2, 10}, {2, 0, 7}, {1, 3, 8}, {4, 0, 10}, {3, 4, 2},
		{4, 2, 10}, {0, 3, 3}, {3, 1, 6}, {2, 4, 5}}, 0, 4, 1), 5)
	println(findCheapestPrice(3, [][]int{{0, 1, 2}, {1, 2, 1}, {2, 0, 10}}, 1, 2, 1), 1)

	println(findCheapestPrice2(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1), 200)
	println(findCheapestPrice2(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0), 500)
	println(findCheapestPrice2(4, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1), 6)
	println(findCheapestPrice2(5, [][]int{{1, 2, 10}, {2, 0, 7}, {1, 3, 8}, {4, 0, 10}, {3, 4, 2},
		{4, 2, 10}, {0, 3, 3}, {3, 1, 6}, {2, 4, 5}}, 0, 4, 1), 5)
	println(findCheapestPrice2(3, [][]int{{0, 1, 2}, {1, 2, 1}, {2, 0, 10}}, 1, 2, 1), 1)
}

/*
https://leetcode.com/problems/cheapest-flights-within-k-stops/

There are n cities connected by m flights. Each flight starts from city u and arrives at v with a price w.

Now given all the cities and flights, together with starting city src and the destination dst, your task is to
find the cheapest price from src to dst with up to k stops. If there is no such route, output -1.

Example 1:
Input:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
Output: 200

The cheapest price from city 0 to city 2 with at most 1 stop costs 200, as marked red in the picture.
Example 2:
Input:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 0
Output: 500
*/

// DP, O(n^2), completely my idea
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	paths := make(map[int]map[int]int)
	for _, flight := range flights {
		from := paths[flight[0]]
		if from == nil {
			from = make(map[int]int)
			paths[flight[0]] = from
		}
		from[flight[1]] = flight[2]
	}

	costs := make([][]int, n+1)
	for i := range costs {
		costs[i] = make([]int, n)
	}

	queue := map[int]bool{src: true}
	step := 1
	for step <= K+1 {
		newQueue := make(map[int]bool)

		for source := range queue {
			if from, ex := paths[source]; ex {
				for dest, price := range from {
					newPrice := costs[step-1][source] + price
					if costs[step][dest] == 0 || costs[step][dest] > newPrice {
						costs[step][dest] = newPrice
					}
					newQueue[dest] = true
				}
			}
		}

		queue = newQueue
		step++
	}

	cost := -1
	for K >= 0 {
		if costs[K+1][dst] > 0 && (cost == -1 || cost > costs[K+1][dst]) {
			cost = costs[K+1][dst]
		}
		K--
	}

	return cost
}

// DV v2, O(n^2), less memory, not my idea, 4 times faster
// this is natural optimisation of the 1st version, when you don't keep entire n+1 x n table, but
//   calculate every row of it in place, thus you only need linear array of n elements, min selection happens
//   at calculation
func findCheapestPrice2(n int, flights [][]int, src int, dst int, K int) int {
	costs := make([]int, n)
	for i := range costs {
		costs[i] = -1
	}
	costs[src] = 0

	for step := 0; step <= K; step++ {
		newCosts := make([]int, n)
		copy(newCosts, costs)

		for _, flight := range flights {
			from := flight[0]
			to := flight[1]
			price := flight[2]

			if costs[from] == -1 {
				continue
			}

			newPrice := costs[from] + price // take new price from previous step calculation
			if newCosts[to] == -1 || newCosts[to] > newPrice {
				newCosts[to] = newPrice
			}
		}

		copy(costs, newCosts)
	}

	return costs[dst]
}

package main

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, nil}), [][]int{{0, 1, 3}, {0, 2, 3}})
}

/*
https://leetcode.com/problems/all-paths-from-source-to-target/

Given a directed, acyclic graph of N nodes.  Find all possible paths from node 0 to node N-1, and return them in any order.

The graph is given as follows:  the nodes are 0, 1, ..., graph.length - 1.  graph[i] is a list of all nodes j
for which the edge (i, j) exists.

Example:
Input: [[1,2], [3], [3], []]
Output: [[0,1,3],[0,2,3]]
Explanation: The graph looks like this:
0--->1
|    |
v    v
2--->3
There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
Note:

The number of nodes in the graph will be in the range [2, 15].
You can print different paths in any order, but you should keep the order of nodes inside one path.
*/

func allPathsSourceTarget(graph [][]int) [][]int {
	result := [][]int{}
	nextStepDfs(graph, 0, []int{}, &result)

	return result
}

func nextStepDfs(graph [][]int, point int, currentPath []int, result *[][]int) {
	if point == len(graph)-1 {
		// found path
		*result = append(*result, append(currentPath, point))
		return
	}

	path := make([]int, len(currentPath)+1)
	copy(path, currentPath)
	path[len(currentPath)] = point

	for _, target := range graph[point] {
		nextStepDfs(graph, target, path, result)
	}
}

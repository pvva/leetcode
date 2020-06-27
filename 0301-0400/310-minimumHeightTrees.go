package main

/*
https://leetcode.com/problems/minimum-height-trees/

For an undirected graph with tree characteristics, we can choose any node as the root. The result graph is then
a rooted tree. Among all possible rooted trees, those with minimum height are called minimum height trees (MHTs).
Given such a graph, write a function to find all the MHTs and return a list of their root labels.

Format
The graph contains n nodes which are labeled from 0 to n - 1. You will be given the number n and a list of
undirected edges (each edge is a pair of labels).
You can assume that no duplicate edges will appear in edges. Since all edges are undirected, [0, 1] is
the same as [1, 0] and thus will not appear together in edges.

Example 1 :

Input: n = 4, edges = [[1, 0], [1, 2], [1, 3]]

        0
        |
        1
       / \
      2   3

Output: [1]

Example 2 :

Input: n = 6, edges = [[0, 3], [1, 3], [2, 3], [4, 3], [5, 4]]

     0  1  2
      \ | /
        3
        |
        4
        |
        5

Output: [3, 4]
Note:
*/

// the idea is to remove nodes with one edge at each iteration until 2 or 1 nodes remain, those nodes is the answer
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 0 {
		return nil
	} else if n == 1 {
		return []int{0}
	} else if n == 2 {
		return []int{0, 1}
	}

	nodesEdges := make([]map[int]bool, n)

	for _, edge := range edges {
		f := edge[0]
		t := edge[1]

		if nodesEdges[f] == nil {
			nodesEdges[f] = make(map[int]bool)
		}
		if nodesEdges[t] == nil {
			nodesEdges[t] = make(map[int]bool)
		}
		nodesEdges[f][t] = true
		nodesEdges[t][f] = true
	}

	queue := []int{}

	for i, ne := range nodesEdges {
		if len(ne) == 1 {
			queue = append(queue, i)
		}
	}

	for n > 2 {
		newQueue := []int{}

		for _, node := range queue {
			for tn := range nodesEdges[node] {
				delete(nodesEdges[tn], node)
				if len(nodesEdges[tn]) == 1 {
					newQueue = append(newQueue, tn)
				}
			}
			n--
		}

		queue = newQueue
	}

	return queue
}

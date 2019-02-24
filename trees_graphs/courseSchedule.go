package main

func main() {
	println(canFinish(2, [][]int{{1, 0}}))
	println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	println(canFinish(3, [][]int{{1, 0}, {2, 1}}))
	println(canFinish(3, [][]int{{0, 2}, {1, 2}, {2, 0}}))
	println(canFinish(2, [][]int{{0, 1}}))
	println(canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
}

/*
here are a total of n courses you have to take, labeled from 0 to n-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1,
which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?

Input: 2, [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0. So it is possible.

Input: 2, [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0, and to take course 0 you should
             also have finished course 1. So it is impossible.
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	// if there are no cycles in graph, then can finish
	links := make(map[int][]int)
	for _, prereq := range prerequisites {
		if l, ex := links[prereq[1]]; ex {
			links[prereq[1]] = append(l, prereq[0])
		} else {
			links[prereq[1]] = []int{prereq[0]}
		}
	}
	seen := make([]bool, numCourses)
	stack := make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		if !seen[i] && dfsDetectCycle(i, seen, stack, links) {
			return false
		}
	}

	return true
}

func dfsDetectCycle(c int, seen, stack []bool, m map[int][]int) bool {
	// Mark the current node as visited and a part of recursion stack
	seen[c] = true
	stack[c] = true
	if l, ex := m[c]; ex {
		for _, v := range l {
			if stack[v] {
				return true
			}
			if !seen[v] && dfsDetectCycle(v, seen, stack, m) {
				return true
			}
		}
	}
	stack[c] = false

	return false
}

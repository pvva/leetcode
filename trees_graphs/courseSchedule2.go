package main

import "fmt"

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

/*
There are a total of n courses you have to take, labeled from 0 to n-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1,
which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, return the ordering of courses you
should take to finish all courses.

There may be multiple correct orders, you just need to return one of them. If it is impossible to finish
all courses, return an empty array.

Input: 2, [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished
             course 0. So the correct course order is [0,1] .

Input: 4, [[1,0],[2,0],[3,1],[3,2]]
Output: [0,1,2,3] or [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both
             courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
             So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3] .
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	// do a topological sort
	links := make(map[int][]int)
	for _, prereq := range prerequisites {
		if l, ex := links[prereq[1]]; ex {
			links[prereq[1]] = append(l, prereq[0])
		} else {
			links[prereq[1]] = []int{prereq[0]}
		}
	}

	ts := &topSort{
		order: make([]int, numCourses),
		oIdx:  numCourses - 1,
	}

	seen := make([]bool, numCourses)
	stack := make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		if !seen[i] && topSortDfsWithCycleDetection(i, seen, stack, links, ts) {
			return []int{}
		}
	}

	return ts.order
}

type topSort struct {
	order []int
	oIdx  int
}

func (ts *topSort) add(v int) {
	ts.order[ts.oIdx] = v
	ts.oIdx--
}

func topSortDfsWithCycleDetection(c int, seen, stack []bool, m map[int][]int, ts *topSort) bool {
	// Mark the current node as visited and a part of recursion stack
	seen[c] = true
	stack[c] = true
	if l, ex := m[c]; ex {
		for _, v := range l {
			if stack[v] {
				return true
			}
			if !seen[v] && topSortDfsWithCycleDetection(v, seen, stack, m, ts) {
				return true
			}
		}
	}
	stack[c] = false
	ts.add(c)

	return false
}

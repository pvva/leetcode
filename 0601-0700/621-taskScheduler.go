package main

import "sort"

func main() {
	println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2), 8)
	println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0), 6)
	println(leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2), 16)
	println(leastInterval([]byte{'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D', 'E', 'E', 'F', 'F', 'G', 'G', 'H', 'H', 'I', 'I', 'J', 'J', 'K', 'K', 'L', 'L', 'M', 'M', 'N', 'N', 'O', 'O', 'P', 'P', 'Q', 'Q', 'R', 'R', 'S', 'S', 'T', 'T', 'U', 'U', 'V', 'V', 'W', 'W', 'X', 'X', 'Y', 'Y', 'Z', 'Z'}, 2), 52)

	println(leastIntervalBetter([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2), 8)
	println(leastIntervalBetter([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0), 6)
	println(leastIntervalBetter([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2), 16)
	println(leastIntervalBetter([]byte{'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D', 'E', 'E', 'F', 'F', 'G', 'G', 'H', 'H', 'I', 'I', 'J', 'J', 'K', 'K', 'L', 'L', 'M', 'M', 'N', 'N', 'O', 'O', 'P', 'P', 'Q', 'Q', 'R', 'R', 'S', 'S', 'T', 'T', 'U', 'U', 'V', 'V', 'W', 'W', 'X', 'X', 'Y', 'Y', 'Z', 'Z'}, 2), 52)
}

/*
https://leetcode.com/problems/task-scheduler/

You are given a char array representing tasks CPU need to do. It contains capital letters A to Z where each
letter represents a different task. Tasks could be done without the original order of the array. Each task
is done in one unit of time. For each unit of time, the CPU could complete either one task or just be idle.

However, there is a non-negative integer n that represents the cooldown period between two same tasks (the
same letter in the array), that is that there must be at least n units of time between any two same tasks.

You need to return the least number of units of times that the CPU will take to finish all the given tasks.


Example 1:

Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Explanation:
A -> B -> idle -> A -> B -> idle -> A -> B
There is at least 2 units of time between any two same tasks.

Example 2:

Input: tasks = ["A","A","A","B","B","B"], n = 0
Output: 6
Explanation: On this case any permutation of size 6 would work since n = 0.
["A","A","A","B","B","B"]
["A","B","A","B","A","B"]
["B","B","B","A","A","A"]
...
And so on.

Example 3:

Input: tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
Output: 16
Explanation:
One possible solution is
A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> idle -> idle -> A -> idle -> idle -> A


Constraints:

The number of tasks is in the range [1, 10000].
The integer n is in the range [0, 100].
*/

// first iteration
func leastInterval(tasks []byte, n int) int {
	if n == 0 || len(tasks) == 0 {
		return len(tasks)
	}

	counts := make([]int, 26)
	for _, v := range tasks {
		counts[int(v-'A')]++
	}
	sort.Ints(counts)

	i := len(counts) - 1
	t := (counts[i]-1)*(n+1) + 1
	fs := (counts[i] - 1) * n
	i--
	for ; i >= 0; i-- {
		if counts[i] == counts[i+1] {
			if fs > 0 {
				t++
				fs = fs - counts[i] + 1
				if fs < 0 {
					t = t - fs
					fs = 0
				}
			} else {
				t += counts[i]
			}
		} else {
			break
		}
	}
	for ; i >= 0; i-- {
		if fs > 0 {
			fs -= counts[i]
			if fs < 0 {
				t = t - fs
				fs = 0
			}
		} else {
			t += counts[i]
		}
	}

	return t
}

func leastIntervalBetter(tasks []byte, n int) int {
	if n == 0 || len(tasks) == 0 {
		return len(tasks)
	}

	counts := make([]int, 26)
	for _, v := range tasks {
		counts[int(v-'A')]++
	}
	max := 0
	mc := 0
	for _, v := range counts {
		if max < v {
			max = v
			mc = 1
		} else if max == v {
			mc++
		}
	}
	t := (max-1)*(n+1) + mc
	if len(tasks) > t {
		return len(tasks)
	}

	return t
}

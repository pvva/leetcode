package main

import "fmt"

func main() {
	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 7), []int{0, 0, 1, 1, 0, 0, 0, 0})
	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 14), []int{0, 0, 0, 0, 1, 1, 0, 0})
	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 21), []int{0, 0, 1, 1, 0, 0, 0, 0})
	fmt.Println(prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 1e9), []int{0, 0, 1, 1, 1, 1, 1, 0})
}

/*
https://leetcode.com/problems/prison-cells-after-n-days/

There are 8 prison cells in a row, and each cell is either occupied or vacant.

Each day, whether the cell is occupied or vacant changes according to the following rules:

If a cell has two adjacent neighbors that are both occupied or both vacant, then the cell becomes occupied.
Otherwise, it becomes vacant.
(Note that because the prison is a row, the first and the last cells in the row can't have two adjacent neighbors.)

We describe the current state of the prison in the following way: cells[i] == 1 if the i-th cell is occupied, else cells[i] == 0.

Given the initial state of the prison, return the state of the prison after N days (and N such changes described above.)


Example 1:

Input: cells = [0,1,0,1,1,0,0,1], N = 7
Output: [0,0,1,1,0,0,0,0]
Explanation:
The following table summarizes the state of the prison on each day:
Day 0: [0, 1, 0, 1, 1, 0, 0, 1]
Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
Day 7: [0, 0, 1, 1, 0, 0, 0, 0]

Example 2:

Input: cells = [1,0,0,1,0,0,1,0], N = 1000000000
Output: [0,0,1,1,1,1,1,0]
*/

// the sequence loops every 14 steps, this was found out by just analyzing the sequence of first 30 steps for
// different inputs
func prisonAfterNDays(cells []int, N int) []int {
	s := transform(intFromBitArray(cells))
	for i := 1; i <= (N-1)%14; i++ {
		s = transform(s)
	}

	return bitArrayFromInt(s)
}

func transform(v int) int {
	r := 0
	for i := 1; i <= 6; i++ {
		lb := (v & (1 << (i - 1))) >> (i - 1)
		rb := (v & (1 << (i + 1))) >> (i + 1)

		r = r | (^(lb^rb)&1)<<i
	}

	return r
}

func intFromBitArray(bits []int) int {
	r := 0
	for _, v := range bits {
		r = (r << 1) | v
	}

	return r
}

func bitArrayFromInt(v int) []int {
	r := make([]int, 8)

	for i := 0; i < 8; i++ {
		r[7-i] = v & 1
		v >>= 1
	}

	return r
}

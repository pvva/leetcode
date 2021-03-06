package main

import "fmt"

func main() {
	input := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(input)
	fmt.Println(input)
}

/*
https://leetcode.com/problems/game-of-life/

According to the Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised
by the British mathematician John Horton Conway in 1970."

Given a board with m by n cells, each cell has an initial state live (1) or dead (0). Each cell interacts with its
eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

Any live cell with fewer than two live neighbors dies, as if caused by under-population.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by over-population..
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

Write a function to compute the next state (after one update) of the board given its current state. The next state is
created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur
simultaneously.

Example:

Input:
[
  [0,1,0],
  [0,0,1],
  [1,1,1],
  [0,0,0]
]
Output:
[
  [0,0,0],
  [1,0,1],
  [0,1,1],
  [0,1,0]
]
Follow up:

Could you solve it in-place? Remember that the board needs to be updated at the same time: You cannot update some cells
first and then use their updated values to update other cells.
In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause
problems when the active area encroaches the border of the array. How would you address these problems?
*/

func gameOfLife(board [][]int) {
	lx := len(board)
	if lx == 0 {
		return
	}
	ly := len(board[0])
	lx--
	ly--

	for x, row := range board {
		for y, v := range row {
			nc := getNC(board, x, y, lx, ly)
			if nc >= 2 && nc <= 3 && ((v&1 == 1) || (nc == 3)) {
				board[x][y] |= 2
			}
		}
	}
	for x, row := range board {
		for y := range row {
			board[x][y] = (board[x][y] >> 1) & 1
		}
	}
}

func getNC(board [][]int, x, y, lx, ly int) int {
	c := 0
	if x >= 1 && y >= 1 {
		c += board[x-1][y-1] & 1
	}
	if x >= 1 && y >= 0 {
		c += board[x-1][y] & 1
	}
	if x >= 1 && y < ly {
		c += board[x-1][y+1] & 1
	}
	if x >= 0 && y >= 1 {
		c += board[x][y-1] & 1
	}
	if x >= 0 && y < ly {
		c += board[x][y+1] & 1
	}
	if x < lx && y >= 1 {
		c += board[x+1][y-1] & 1
	}
	if x < lx && y >= 0 {
		c += board[x+1][y] & 1
	}
	if x < lx && y < ly {
		c += board[x+1][y+1] & 1
	}

	return c
}

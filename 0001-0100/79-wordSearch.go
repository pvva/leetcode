package main

func main() {
	println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"), true)
	println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"), true)
	println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"), false)
}

/*
https://leetcode.com/problems/word-search/

Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally
or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.


Constraints:

board and word consists only of lowercase and uppercase English letters.
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3
*/

func exist(board [][]byte, word string) bool {
	for x, row := range board {
		for y := range row {
			if executeDfsSearch(board, &word, x, y, 0) {
				return true
			}
		}
	}

	return false
}

func executeDfsSearch(board [][]byte, word *string, x, y, from int) bool {
	if from == len(*word) {
		return true
	}
	if x < 0 || x == len(board) || y < 0 || y == len(board[0]) || board[x][y] != (*word)[from] {
		return false
	}

	c := board[x][y]
	board[x][y] = '-'
	success := executeDfsSearch(board, word, x, y+1, from+1)
	success = success || executeDfsSearch(board, word, x+1, y, from+1)
	success = success || executeDfsSearch(board, word, x, y-1, from+1)
	success = success || executeDfsSearch(board, word, x-1, y, from+1)

	if !success {
		board[x][y] = c
	}

	return success
}

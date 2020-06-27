package main

func main() {
	println(validTicTacToe([]string{"O  ", "   ", "   "}), false)
	println(validTicTacToe([]string{"XOX", " X ", "   "}), false)
	println(validTicTacToe([]string{"XXX", "   ", "OOO"}), false)
	println(validTicTacToe([]string{"XOX", "O O", "XOX"}), true)
}

/*
https://leetcode.com/problems/valid-tic-tac-toe-state/

A Tic-Tac-Toe board is given as a string array board. Return True if and only if it is possible to reach this board
position during the course of a valid tic-tac-toe game.
The board is a 3 x 3 array, and consists of characters " ", "X", and "O".  The " " character represents an empty square.

Here are the rules of Tic-Tac-Toe:

Players take turns placing characters into empty squares (" ").
The first player always places "X" characters, while the second player always places "O" characters.
"X" and "O" characters are always placed into empty squares, never filled ones.
The game ends when there are 3 of the same (non-empty) character filling any row, column, or diagonal.
The game also ends if all squares are non-empty.
No more moves can be played if the game is over.
Example 1:
Input: board = ["O  ", "   ", "   "]
Output: false
Explanation: The first player always plays "X".

Example 2:
Input: board = ["XOX", " X ", "   "]
Output: false
Explanation: Players take turns making moves.

Example 3:
Input: board = ["XXX", "   ", "OOO"]
Output: false

Example 4:
Input: board = ["XOX", "O O", "XOX"]
Output: true
*/

func validTicTacToe(board []string) bool {
	x := 0
	o := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 'X' {
				x++
			} else if board[i][j] == 'O' {
				o++
			}
		}
	}
	winX := hasWinningCombination(board, "XXX")
	winO := hasWinningCombination(board, "OOO")

	return !(o > x || x-o > 1 || (winX && winO) || (winX && o == x) || (winO && x > o))
}

func hasWinningCombination(board []string, row string) bool {
	return board[0] == row || board[1] == row || board[2] == row ||
		(board[0][0] == row[0] && board[1][0] == row[0] && board[2][0] == row[0]) ||
		(board[0][1] == row[0] && board[1][1] == row[0] && board[2][1] == row[0]) ||
		(board[0][2] == row[0] && board[1][2] == row[0] && board[2][2] == row[0]) ||
		(board[0][0] == row[0] && board[1][1] == row[0] && board[2][2] == row[0]) ||
		(board[0][2] == row[0] && board[1][1] == row[0] && board[2][0] == row[0])
}

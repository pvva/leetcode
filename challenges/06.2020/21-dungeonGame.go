package main

func main() {
	println(calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}), 7)
	println(calculateMinimumHP([][]int{{100}}), 1)
	println(calculateMinimumHP([][]int{{0, 0, 0}, {1, 1, -1}}), 1)
	println(calculateMinimumHP([][]int{{0, 1, 1}, {0, 0, -1}}), 1)
	println(calculateMinimumHP([][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}), 3)
	println(calculateMinimumHP([][]int{{3, -20, 30}, {-3, 4, 0}}), 1)
}

/*
The demons had captured the princess (P) and imprisoned her in the bottom-right corner of a dungeon. The dungeon
consists of M x N rooms laid out in a 2D grid. Our valiant knight (K) was initially positioned in the top-left
room and must fight his way through the dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer. If at any point his health point drops
to 0 or below, he dies immediately.

Some of the rooms are guarded by demons, so the knight loses health (negative integers) upon entering these rooms;
other rooms are either empty (0's) or contain magic orbs that increase the knight's health (positive integers).

In order to reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.

Write a function to determine the knight's minimum initial health so that he is able to rescue the princess.

For example, given the dungeon below, the initial health of the knight must be at least 7 if he follows the optimal
path RIGHT-> RIGHT -> DOWN -> DOWN.

-2 (K)	-3	3
-5	-10	1
10	30	-5 (P)

Note:

The knight's health has no upper bound.
Any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the
princess is imprisoned.
*/

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}
	xl := len(dungeon)
	yl := len(dungeon[0])

	memo := make([]int, yl)
	for x := xl - 1; x >= 0; x-- {
		for y := yl - 1; y >= 0; y-- {

			if x == xl-1 {
				if y == yl-1 {
					memo[y] = max(1, 1-dungeon[x][y])
				} else {
					memo[y] = max(1, memo[y+1]-dungeon[x][y])
				}
			} else {
				if y == yl-1 {
					memo[y] = max(1, memo[y]-dungeon[x][y])
				} else {
					right := max(1, memo[y+1]-dungeon[x][y])
					down := max(1, memo[y]-dungeon[x][y])
					memo[y] = min(right, down)
				}
			}

		}
	}

	return memo[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

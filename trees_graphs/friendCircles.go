package main

func main() {
	input := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	//input := [][]int{
	//	{1, 0, 0, 1},
	//	{0, 1, 1, 0},
	//	{0, 1, 1, 1},
	//	{1, 0, 1, 1},
	//}

	println(findCircleNum(input))
}

/*
There are N students in a class. Some of them are friends, while some are not. Their friendship is transitive in nature.
For example, if A is a direct friend of B, and B is a direct friend of C, then A is an indirect friend of C. And we
defined a friend circle is a group of students who are direct or indirect friends.

Given a N*N matrix M representing the friend relationship between students in the class. If M[i][j] = 1, then the ith and
jth students are direct friends with each other, otherwise not. And you have to output the total number of friend
circles among all the students.
*/
func findCircleNum(M [][]int) int {
	c := 0

	for x := 0; x < len(M); x++ {
		if M[x][x] == 1 {
			c++
			dfs(x, M)
		}
	}

	return c
}

func dfs(x int, M [][]int) {
	if M[x][x] != 1 {
		return
	}
	M[x][x] = 2
	for y := 0; y < len(M); y++ {
		if M[x][y] == 1 {
			M[x][y] = 2
			dfs(y, M)
		}
	}
}

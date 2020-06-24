package main

func main() {
	println(numTrees(3), 5)
	println(numTrees(4))
}

func numTrees(n int) int {
	return numTreesRec(n, map[int]int{0: 1, 1: 1, 2: 2})
}

func numTreesRec(n int, memo map[int]int) int {
	if v, ex := memo[n]; ex {
		return v
	}

	cnt := 0
	for i := 1; i <= n; i++ {
		cnt += numTreesRec(i-1, memo) * numTreesRec(n-i, memo)
	}
	memo[n] = cnt

	return cnt
}

package main

func main() {
	println(coinChange([]int{1, 2, 5}, 11), 3)
	println(coinChange([]int{2, 5, 10, 1}, 27), 4)
	println(coinChange([]int{1, 2, 5}, 100), 20)
}

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}

	bestSteps := make([]int, amount+1)
	for i := range bestSteps {
		bestSteps[i] = -1
	}
	bestSteps[0] = 0

	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if c > a {
				continue
			}
			t := bestSteps[a-c] + 1
			if t > 0 && (bestSteps[a] == -1 || t < bestSteps[a]) {
				bestSteps[a] = t
			}
		}
	}

	return bestSteps[amount]
}

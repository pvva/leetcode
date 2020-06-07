package main

func main() {
	println(canGetExactChange(94, []int{5, 10, 25, 100, 200}), false)
	println(canGetExactChange(75, []int{4, 17, 29}), true)
}

/*
Given a list of the available denominations, determine if it's possible to receive exact change for an amount
of money targetMoney. Both the denominations and target amount will be given in generic units of that currency.

Example 1
denominations = [5, 10, 25, 100, 200]
targetMoney = 94
output = false
Every denomination is a multiple of 5, so you can't receive exactly 94 units of money in this currency.

Example 2
denominations = [4, 17, 29]
targetMoney = 75
output = true
You can make 75 units with the denominations [17, 29, 29].
*/

func canGetExactChange(targetMoney int, denominations []int) bool {
	// O(n^2) where n is the length of denominations
	if targetMoney == 0 {
		return true
	}

	i := 0
	for i < len(denominations) && denominations[i] > targetMoney {
		i++
	}

	for ; i < len(denominations); i++ {
		if canGetExactChange(targetMoney-denominations[i], denominations) {
			return true
		}
	}

	return false
}

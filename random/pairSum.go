package main

func main() {
	println(numberOfWays([]int{}, 2), 0)
	println(numberOfWays([]int{1, 3, 4}, 2), 0)
	println(numberOfWays([]int{1, 2, 3, 4, 3}, 6), 2)
	println(numberOfWays([]int{1, 5, 3, 3, 3}, 6), 4)
}

/*
Given a list of n integers arr[0..(n-1)], determine the number of different pairs of elements within it which sum to k.
If an integer appears in the list multiple times, each copy is considered to be different; that is, two pairs are
considered different if one pair includes at least one array index which the other doesn't, even if they include
the same values.

n is in the range [1, 100,000].
Each value arr[i] is in the range [1, 1,000,000,000].
k is in the range [1, 1,000,000,000].

Example 1
n = 5
k = 6
arr = [1, 2, 3, 4, 3]
output = 2
The valid pairs are 2+4 and 3+3.

Example 2
n = 5
k = 6
arr = [1, 5, 3, 3, 3]
output = 4
There's one valid pair 1+5, and three different valid pairs 3+3 (the 3rd and 4th elements, 3rd and 5th elements,
and 4th and 5th elements).
*/

// based on simple math
func numberOfWays(arr []int, k int) int {
	counts := make(map[int]int)
	for _, v := range arr {
		counts[v]++
	}

	total := 0
	for first, firstCount := range counts {
		second := k - first
		if secondCount, ex := counts[second]; ex {
			if first == second {
				total += firstCount * (firstCount - 1) / 2 // same as combination of 2 from *count*
			} else {
				total += firstCount * secondCount
			}

			delete(counts, first)
			delete(counts, second)
		}
	}

	return total
}

package main

func main() {
	println(getPermutation(4, 8), "2143")
}

/*
The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.

Note:

Given n will be between 1 and 9 inclusive.
Given k will be between 1 and n! inclusive.

Example 1:

Input: n = 3, k = 3
Output: "213"
Example 2:

Input: n = 4, k = 9
Output: "2314"
*/

var factors = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func getPermutation(n int, k int) string {
	res := make([]byte, n)
	ri := 0
	a := make([]byte, n)
	for i := range a {
		a[i] = byte(i) + '1'
	}
	if k == 1 {
		return string(a)
	}

	for n > 1 {
		di := 0
		if k > 0 {
			di = (k - 1) / factors[n-1]
		}

		res[ri] = a[di]
		copy(a[di:], a[di+1:])

		k = k - di*factors[n-1]
		n--
		ri++
	}
	res[ri] = a[0]

	return string(res)
}

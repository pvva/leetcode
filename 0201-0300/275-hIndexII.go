package main

func main() {
	println(hIndex([]int{1}), 1)
	println(hIndex([]int{0}), 0)
	println(hIndex([]int{100}), 1)
	println(hIndex([]int{1, 1}), 1)
	println(hIndex([]int{1, 1, 1}), 1)
	println(hIndex([]int{1, 2, 3}), 2)
	println(hIndex([]int{1, 1, 1, 1, 2}), 1)
	println(hIndex([]int{1, 1, 1, 2, 2}), 2)
	println("--")
	println(hIndex2([]int{1}), 1)
	println(hIndex2([]int{0}), 0)
	println(hIndex2([]int{100}), 1)
	println(hIndex2([]int{1, 1}), 1)
	println(hIndex2([]int{1, 1, 1}), 1)
	println(hIndex2([]int{1, 2, 3}), 2)
	println(hIndex2([]int{1, 1, 1, 1, 2}), 1)
	println(hIndex2([]int{1, 1, 1, 2, 2}), 2)
}

/*
https://leetcode.com/problems/h-index-ii/

Given an array of citations sorted in ascending order (each citation is a non-negative integer) of a researcher,
write a function to compute the researcher's h-index.

According to the definition of h-index on Wikipedia: "A scientist has index h if h of his/her N papers have at
least h citations each, and the other N − h papers have no more than h citations each."

Example:

Input: citations = [0,1,3,5,6]
Output: 3
Explanation: [0,1,3,5,6] means the researcher has 5 papers in total and each of them had
             received 0, 1, 3, 5, 6 citations respectively.
             Since the researcher has 3 papers with at least 3 citations each and the remaining
             two with no more than 3 citations each, her h-index is 3.
Note:

If there are several possible values for h, the maximum one is taken as the h-index.

Follow up:

This is a follow up problem to H-Index, where citations is now guaranteed to be sorted in ascending order.
Could you solve it in logarithmic time complexity?
*/

func hIndex(citations []int) int {
	l, r := 0, len(citations)-1

	for l <= r {
		m := (r + l) / 2
		h := len(citations) - m

		if citations[m] < h {
			l = m + 1
		} else if citations[m] >= h {
			if m == 0 || r == m || (m > 0 && citations[m-1] < h) {
				return h
			}
			r = m
		} else {
			return h
		}
	}

	return 0
}

func hIndex2(citations []int) int {
	s := len(citations)
	l, r := 0, s-1

	for l <= r {
		m := (r + l) / 2
		h := s - m

		if citations[m] == h {
			return h
		} else if citations[m] < h {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return s - l
}

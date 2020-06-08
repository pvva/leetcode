package main

import "sort"

func main() {
	println(minOperations([]int{3, 1, 2}), 2)
	println(minOperations([]int{2, 5, 3, 6, 4, 7, 1}))
}

/*
In this problem, you are given an integer N, and a permutation, P of the integers from 1 to N, denoted as
(a_1, a_2, ..., a_N). You want to rearrange the elements of the permutation into increasing order, repeatedly
making the following operation:
Select a sub-portion of the permutation, (a_i, ..., a_j), and reverse its order.
Your goal is to compute the minimum number of such operations required to return the permutation to increasing order.

Example
If N = 3, and P = (3, 1, 2), we can do the following operations:
Select (1, 2) and reverse it: P = (3, 2, 1).
Select (3, 2, 1) and reverse it: P = (1, 2, 3).
output = 2
*/

func minOperations(arr []int) int {
	// the only idea I have is to use brute force
	// for given permutation calculate all possible permutations that can be constructed from it and make them be next set of permuations to check
	// for each next set of permutations to check increase the steps amount
	// continue for each current set of permutations, start from set containing just initial permutation
	// this corresponds to bfs search in the space of permutations

	desiredTarget := make([]int, len(arr))
	copy(desiredTarget, arr)
	sort.Ints(desiredTarget)
	targetHash := permHash(desiredTarget)

	seenPerms := make(map[uint]bool)
	perms := [][]int{arr}
	steps := -1

	for len(perms) > 0 {
		steps++
		newPerms := [][]int{}
		for _, perm := range perms {
			pHash := permHash(perm)

			if pHash == targetHash {
				return steps
			}
			seenPerms[pHash] = true

			// generate all possible perms from current one
			for i := 0; i < len(arr)-1; i++ {
				for j := i + 1; j < len(arr); j++ {
					tPerm := make([]int, len(arr))
					copy(tPerm, perm)
					reverse(tPerm[i : j+1])
					tHash := permHash(tPerm)
					if seenPerms[tHash] {
						continue
					}

					newPerms = append(newPerms, tPerm)
				}
			}

		}
		perms = newPerms
	}

	return steps
}

func permHash(arr []int) uint {
	hash := uint(0)

	for _, v := range arr {
		hash <<= 4
		hash |= uint(v)
	}

	return hash
}

func reverse(arr []int) {
	l := 0
	r := len(arr) - 1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

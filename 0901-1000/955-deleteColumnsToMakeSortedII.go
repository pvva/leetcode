package main

func main() {
	println(minDeletionSize([]string{"xga", "xfb", "yfa"}), 1)
	println(minDeletionSize([]string{"bwwdyeyfhc", "bchpphbtkh", "hmpudwfkpw", "lqeoyqkqwe", "riobghmpaa", "stbheblgao", "snlaewujlc", "tqlzolljas", "twdkexzvfx", "wacnnhjdis"}), 4)

	println(minDeletionSizeImproved([]string{"xga", "xfb", "yfa"}), 1)
	println(minDeletionSizeImproved([]string{"axx", "ayy", "baa", "bbb", "bcc"}), 0)
	println(minDeletionSizeImproved([]string{"azeee", "azeee"}), 0)
	println(minDeletionSizeImproved([]string{"zyx", "wvu", "tsr"}), 3)
	println(minDeletionSizeImproved([]string{"abx", "agz", "bgc", "bfc"}), 1)
	println(minDeletionSizeImproved([]string{"ousnatait", "xzswvwztr", "luknznxob"}), 4)
	println(minDeletionSizeImproved([]string{"bwwdyeyfhc", "bchpphbtkh", "hmpudwfkpw", "lqeoyqkqwe", "riobghmpaa", "stbheblgao", "snlaewujlc", "tqlzolljas", "twdkexzvfx", "wacnnhjdis"}), 4)
}

/*
https://leetcode.com/problems/delete-columns-to-make-sorted-ii/

We are given an array A of N lowercase letter strings, all of the same length.

Now, we may choose any set of deletion indices, and for each string, we delete all the characters in those indices.

For example, if we have an array A = ["abcdef","uvwxyz"] and deletion indices {0, 2, 3}, then the final array after
deletions is ["bef","vyz"].

Suppose we chose a set of deletion indices D such that after deletions, the final array has its elements in
lexicographic order (A[0] <= A[1] <= A[2] ... <= A[A.length - 1]).

Return the minimum possible value of D.length.

Example 1:

Input: ["ca","bb","ac"]
Output: 1
Explanation:
After deleting the first column, A = ["a", "b", "c"].
Now A is in lexicographic order (ie. A[0] <= A[1] <= A[2]).
We require at least 1 deletion since initially A was not in lexicographic order, so the answer is 1.

Example 2:

Input: ["xc","yb","za"]
Output: 0
Explanation:
A is already in lexicographic order, so we don't need to delete anything.
Note that the rows of A are not necessarily in lexicographic order:
ie. it is NOT necessarily true that (A[0][0] <= A[0][1] <= ...)

Example 3:

Input: ["zyx","wvu","tsr"]
Output: 3
Explanation:
We have to delete every column.

Note:

1 <= A.length <= 100
1 <= A[i].length <= 100
*/

// running time is O((n * m)^2), n - size of A, m - length of each string
// memory is O(m)
func minDeletionSize(A []string) int {
	if len(A) == 0 {
		return 0
	}

	dropToMakeOrder := make(map[int]bool)
	for i := 0; i < len(A)-1; i++ {
		newDrop := makeInLexOrder(A[i], A[i+1], dropToMakeOrder)
		if len(newDrop) > 0 {
			for k := range newDrop {
				dropToMakeOrder[k] = true
			}
			i = -1
		}
	}

	return len(dropToMakeOrder)
}

func makeInLexOrder(s1, s2 string, currentDrop map[int]bool) map[int]bool {
	dropToMakeOrder := make(map[int]bool)
	for i := range s1 {
		if currentDrop[i] {
			continue
		}
		if s1[i] > s2[i] {
			dropToMakeOrder[i] = true
		} else if s1[i] < s2[i] {
			break
		}
	}

	return dropToMakeOrder
}

// running time is O(n * m), n - size of A, m - length of each string
// memory is O(m)
func minDeletionSizeImproved(A []string) int {
	if len(A) == 0 {
		return 0
	}
	dropped := make(map[int]bool)
	searchDisorder(A, 0, len(A)-1, 0, len(A[0]), dropped)

	return len(dropped)
}

func searchDisorder(A []string, fromX, toX, fromY, toY int, dropped map[int]bool) {
	for y := fromY; y < toY; y++ {
		if dropped[y] {
			continue
		}
		drop := false
		monotonic := [][2]int{}
		cont := -1
		for x := fromX; x < toX; x++ {
			if A[x][y] > A[x+1][y] {
				drop = true
				dropped[y] = true
				break
			} else if A[x][y] == A[x+1][y] {
				if cont == -1 {
					cont = x
				}
			} else {
				if cont >= 0 {
					monotonic = append(monotonic, [2]int{cont, x})
					cont = -1
				}
			}
		}
		if cont >= 0 {
			monotonic = append(monotonic, [2]int{cont, toX})
		}

		if !drop {
			for _, rng := range monotonic {
				searchDisorder(A, rng[0], rng[1], y+1, toY, dropped)
			}
			break
		}
	}
}

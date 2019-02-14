package main

func main() {
	a := []int{0, 1, -1}
	b := []int{-1, 1, 0}
	c := []int{0, 0, 1}
	d := []int{-1, 1, 1}

	println(fourSumCount(a, b, c, d))
}

// Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.
func fourSumCount(A []int, B []int, C []int, D []int) int {
	result := 0
	hsum := make(map[int]int)

	for aIdx := 0; aIdx < len(A); aIdx++ {
		for bIdx := 0; bIdx < len(B); bIdx++ {
			hs := A[aIdx] + B[bIdx]
			if t, ex := hsum[hs]; ex {
				hsum[hs] = t + 1
			} else {
				hsum[hs] = 1
			}
		}
	}
	for cIdx := 0; cIdx < len(C); cIdx++ {
		for dIdx := 0; dIdx < len(D); dIdx++ {
			hs := -C[cIdx] - D[dIdx]
			if t, ex := hsum[hs]; ex {
				result += t
			}
		}
	}

	return result
}

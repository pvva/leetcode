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

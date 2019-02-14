package main

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	println(maxArea([]int{1, 2, 1}))
	println(maxArea2([]int{1, 2, 1}))

	println(maxArea([]int{1, 2, 4, 3}))
	println(maxArea2([]int{1, 2, 4, 3}))

	println(maxArea([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	println(maxArea2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func maxArea(height []int) int {
	area := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			c := min(height[i], height[j]) * (j - i)
			if area < c {
				area = c
			}
		}
	}

	return area
}

func maxArea2(height []int) int {
	area := 0
	i := 0
	e := len(height) - 1

	for i < e {
		a := min(height[i], height[e]) * (e - i)
		if a > area {
			area = a
		}
		if height[i] < height[e] {
			i++
		} else {
			e--
		}
	}

	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

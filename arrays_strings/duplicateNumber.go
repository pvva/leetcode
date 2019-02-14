package main

func main() {
	//println(findDuplicate([]int{1, 3, 4, 2, 2}))
	//println(findDuplicate([]int{3, 1, 3, 4, 2}))
	println(findDuplicate([]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}))
}

func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	s := nums[0]
	f := nums[s]

	for s != f {
		s = nums[s]
		f = nums[nums[f]]
	}

	f = 0
	for s != f {
		f = nums[f]
		s = nums[s]
	}

	return s
}

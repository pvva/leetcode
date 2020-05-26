package main

func main() {
	println(circularArrayLoopNaive([]int{2, -1, 1, 2, 2}), circularArrayLoop([]int{2, -1, 1, 2, 2}), true)
	println(circularArrayLoopNaive([]int{-2, 1, -1, -2, -2}), circularArrayLoop([]int{-2, 1, -1, -2, -2}), false)
	println(circularArrayLoopNaive([]int{-1, 2, -1, 2, -1, 3, 2}), circularArrayLoop([]int{-1, 2, -1, 2, -1, 3, 2}), true)
	println(circularArrayLoopNaive([]int{-1, 2}), circularArrayLoop([]int{-1, 2}), false)
	println(circularArrayLoopNaive([]int{-1, -2, -3, -4, -5}), circularArrayLoop([]int{-1, -2, -3, -4, -5}), false)
	println(circularArrayLoopNaive([]int{3, 1, 2}), circularArrayLoop([]int{3, 1, 2}), true)
}

/*
https://leetcode.com/problems/circular-array-loop/

You are given a circular array nums of positive and negative integers. If a number k at an index is positive,
then move forward k steps. Conversely, if it's negative (-k), move backward k steps. Since the array is circular,
you may assume that the last element's next element is the first element, and the first element's previous element
is the last element.

Determine if there is a loop (or a cycle) in nums. A cycle must start and end at the same index and the cycle's
length > 1. Furthermore, movements in a cycle must all follow a single direction. In other words, a cycle must
not consist of both forward and backward movements.

Example 1:

Input: [2,-1,1,2,2]
Output: true
Explanation: There is a cycle, from index 0 -> 2 -> 3 -> 0. The cycle's length is 3.

Example 2:

Input: [-1,2]
Output: false
Explanation: The movement from index 1 -> 1 -> 1 ... is not a cycle, because the cycle's length is 1.
By definition the cycle's length must be greater than 1.

Example 3:

Input: [-2,1,-1,-2,-2]
Output: false
Explanation: The movement from index 1 -> 2 -> 1 -> ... is not a cycle, because movement from index 1 -> 2 is
a forward movement, but movement from index 2 -> 1 is a backward movement. All movements in a cycle must follow
a single direction.
*/

func circularArrayLoopNaive(nums []int) bool {
	for i, v := range nums {
		seenIdx := make(map[int]bool)
		positive := v > 0

		cIdx := i
		for {
			if nums[cIdx] > 0 != positive {
				break
			}
			nextIdx := cIdx + nums[cIdx]
			if nextIdx >= len(nums) {
				nextIdx %= len(nums)
			} else if nextIdx < 0 {
				nextIdx = nextIdx % len(nums)
				if nextIdx < 0 {
					nextIdx += len(nums)
				}
			}

			if nextIdx == cIdx {
				break
			}
			if seenIdx[cIdx] {
				return true
			}
			seenIdx[cIdx] = true
			cIdx = nextIdx
		}
	}

	return false
}

func circularArrayLoop(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	for i := range nums {
		fi := i
		si := i
		positive := nums[fi] > 0

		for {
			fi = getNextIndexInCircularArray(nums, fi)
			if fi == -1 || nums[fi] > 0 != positive {
				break
			}
			si = getNextIndexInCircularArray(nums, si)
			if si == -1 || nums[si] > 0 != positive {
				break
			}
			si = getNextIndexInCircularArray(nums, si)
			if si == -1 || nums[si] > 0 != positive {
				break
			}

			if fi == si {
				return true
			}
		}
	}

	return false
}

func getNextIndexInCircularArray(nums []int, i int) int {
	next := (i + nums[i]) % len(nums)

	if next < 0 {
		next += len(nums)
	}
	if next == i {
		return -1
	}

	return next
}

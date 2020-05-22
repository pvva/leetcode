package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {
	println(largestNumber([]int{1, 20}), "201")
	println(largestNumber([]int{10, 2}), "210")
	println(largestNumber([]int{3, 30, 34, 5, 9}), "9534330")
	println(largestNumber([]int{0, 0}), "0")
}

/*
https://leetcode.com/problems/largest-number/

Given a list of non negative integers, arrange them such that they form the largest number.

Example 1:

Input: [10,2]
Output: "210"

Example 2:

Input: [3,30,34,5,9]
Output: "9534330"

Note: The result may be very large, so you need to return a string instead of an integer.
*/

func largestNumber(nums []int) string {
	numsStr := make([]string, len(nums))

	allZeros := true
	for i, v := range nums {
		allZeros = allZeros && v == 0
		numsStr[i] = strconv.Itoa(v)
	}
	if allZeros {
		return "0"
	}

	sort.SliceStable(numsStr, func(i, j int) bool {
		return firstGoesToTheLeft(numsStr[i], numsStr[j])
	})

	return strings.Join(numsStr, "")
}

func firstGoesToTheLeft(s1, s2 string) bool {
	toTheLeft := true
	i1 := 0
	i2 := 0

	for i1 < len(s1) && i2 < len(s2) {
		if s1[i1] < s2[i2] {
			toTheLeft = false
			break
		} else if s1[i1] > s2[i2] {
			return true
		}
		i1++
		i2++
	}

	if toTheLeft {
		if len(s1) > len(s2) {
			toTheLeft = firstGoesToTheLeft(s1[i1:], s2)
		} else if len(s1) < len(s2) {
			toTheLeft = firstGoesToTheLeft(s1, s2[i2:])
		}
	}

	return toTheLeft
}

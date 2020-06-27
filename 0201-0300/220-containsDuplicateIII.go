package main

func main() {
	println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0), true)
	println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2), true)
	println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3), false)
	println(containsNearbyAlmostDuplicate([]int{2, 2}, 3, 0), true)
	println(containsNearbyAlmostDuplicate([]int{1, 3, 6, 2}, 1, 2), true)

	println(containsNearbyAlmostDuplicateBuckets([]int{1, 2, 3, 1}, 3, 0), true)
	println(containsNearbyAlmostDuplicateBuckets([]int{1, 0, 1, 1}, 1, 2), true)
	println(containsNearbyAlmostDuplicateBuckets([]int{1, 5, 9, 1, 5, 9}, 2, 3), false)
	println(containsNearbyAlmostDuplicateBuckets([]int{2, 2}, 3, 0), true)
	println(containsNearbyAlmostDuplicateBuckets([]int{1, 3, 6, 2}, 1, 2), true)
}

/*
https://leetcode.com/problems/contains-duplicate-iii/

Given an array of integers, find out whether there are two distinct indices i and j in the array such that the
absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.

Example 1:

Input: nums = [1,2,3,1], k = 3, t = 0
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1, t = 2
Output: true
Example 3:

Input: nums = [1,5,9,1,5,9], k = 2, t = 3
Output: false
*/

// brute force, obvious
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			diff := abs(nums[i] - nums[j])
			if diff <= t {
				return true
			}
		}
	}

	return false
}

// bucket based approach
// the idea here is to calculate the bucket for current value nums[i]
// the bucket is calculated as nums[i] / t rounded down to nearest integer,
//    for example, if t is 4, then value 3 goes to bucket 0 (3/4=0.75) and value 4 goes to bucket 1,
//    so we have [0..3] => 0, [4..7] => 1 and so on
// in order to see if there are matches, we need to check not only the bucket of current number, but also
//    adjacent buckets, because if t is 4 as in example above and values are 0 and 4, then they satisfy conditions,
//    but appear in different buckets, namely bucket 0 and 1
// special case is when t equals 0, which means that we only need to check the bucket of current number, as numbers
//    must be equal
func containsNearbyAlmostDuplicateBuckets(nums []int, k int, t int) bool {
	buckets := make(map[int][2]int)

	for i, v := range nums {
		bucket := v
		offset := 0
		if t != 0 {
			bucket = v / t
			offset = 1
		}

		for bucketValue := bucket - offset; bucketValue <= bucket+offset; bucketValue++ {
			if b, ex := buckets[bucketValue]; ex {
				if abs(b[0]-v) <= t && abs(b[1]-i) <= k {
					return true
				}
			}
		}
		buckets[bucket] = [2]int{v, i}
	}

	return false
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

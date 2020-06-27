package main

func main() {
	println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}), 3)
	println(longestWPI([]int{6, 9, 6, 9, 6, 9, 9, 9, 9, 6, 6}), 11)
	println(longestWPI([]int{6, 6, 9}), 1)
}

/*
https://leetcode.com/problems/longest-well-performing-interval/

We are given hours, a list of the number of hours worked per day for a given employee.
A day is considered to be a tiring day if and only if the number of hours worked is (strictly) greater than 8.
A well-performing interval is an interval of days for which the number of tiring days is strictly larger than the
number of non-tiring days.

Return the length of the longest well-performing interval.

Example 1:

Input: hours = [9,9,6,0,6,6,9]
Output: 3
Explanation: The longest well-performing interval is [9,9,6].

Constraints:

1 <= hours.length <= 10000
0 <= hours[i] <= 16
*/

/*
Generic algorithm for finding longest consecutive subarray with sum at least K.
Let's denote S[i,j] as sum of elements from i to j index.
To find longest consecutive subarray with sum at least K we need for every index i find S[h,i] >= K for all h >= 0 && h <= i.
Let's use current running sum to find S[i] = S[0,i] (sum of elements from the beginning of the array to index i).
1) If S[i] >= K, then current length of the so far longest consecutive subarray is i+1.
2) If S[i] < K, then we need to find if exists such sum S[h,i] >= K for h >=0 && h <= i.
   Si = S[h-1] + S[h,i] =>
   S[i] - S[h-1] = S[h,i] >= K =>
   the task is to find sums that satisfy condition S[i] - S[h-1] >= K which is the same as S[h-1] <= S[i] - K
   Lower bound for this expression is S[h-1] = S[i] - K, because if S[h-1] is less then lower bound, then it already
   satisfies condition.
   If exists such S[h-1] equal to S[i] - K then length of the so far longest consecutive subarray is i - h + 1.
Let's store additional information of running sum => index only for the first occurrence of each running dum value.
Then if exists such pair S[h-1] => h-1 so that S[h-1] is equal to S[i] - K, then length can be calculated.
*/

// for this task let's consider hour > 8 as 1 and others as -1 and find the longest consecutive
// subarray that sums up to at least 1
func longestWPI(hours []int) int {
	maxLen := 0
	sums := make(map[int]int) // sum => index
	runningSum := 0

	for i, v := range hours {
		if v > 8 {
			runningSum++
		} else {
			runningSum--
		}
		cm := 0
		if runningSum >= 1 { // >= K
			cm = i + 1
		} else if s, ex := sums[runningSum-1]; ex { // s => h-1
			cm = i - s // i - h + 1
		}
		if _, ex := sums[runningSum]; !ex {
			sums[runningSum] = i
		}
		if maxLen < cm {
			maxLen = cm
		}
	}

	return maxLen
}

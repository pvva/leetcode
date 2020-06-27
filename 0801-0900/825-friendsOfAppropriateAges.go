package main

func main() {
	println(numFriendRequests([]int{20, 30, 100, 110, 120}), 3)
	println(numFriendRequests([]int{101, 56, 69, 48, 30}), 4)
	println(numFriendRequests([]int{8, 85, 24, 85, 69}), 4)
}

/*
https://leetcode.com/problems/friends-of-appropriate-ages/

Some people will make friend requests. The list of their ages is given and ages[i] is the age of the ith person.
Person A will NOT friend request person B (B != A) if any of the following conditions are true:

age[B] <= 0.5 * age[A] + 7
age[B] > age[A]
age[B] > 100 && age[A] < 100

Otherwise, A will friend request B.
Note that if A requests B, B does not necessarily request A.  Also, people will not friend request themselves.
How many total friend requests are made?

Example 1:

Input: [16,16]
Output: 2
Explanation: 2 people friend request each other.

Example 2:

Input: [16,17,18]
Output: 2
Explanation: Friend requests are made 17 -> 16, 18 -> 17.

Example 3:

Input: [20,30,100,110,120]
Output:
Explanation: Friend requests are made 110 -> 100, 120 -> 110, 120 -> 100.
*/

func numFriendRequests(ages []int) int {
	agesMap := make(map[int]int)

	for _, age := range ages {
		agesMap[age]++
	}

	cnt := 0
	for age, ageAmount := range agesMap {
		s := 0
		e := age
		// this is an error in conditions, age[B] > 100 && age[A] < 100 is not taken into account in correct solutions
		//if age > 100 {
		//	s = 100
		//} else {
		s = age/2 + 8
		//}
		for ; s <= e; s++ {
			if amount, ex := agesMap[s]; ex {
				if s == age {
					cnt += amount * (amount - 1)
				} else {
					cnt += amount * ageAmount
				}
			}
		}
	}

	return cnt
}

package main

import "math"

func main() {
	println(angleClock(12, 30), 165)
	println(angleClock(3, 30), 75)
	println(angleClock(3, 15), 7.5)
	println(angleClock(4, 50), 155)
	println(angleClock(12, 0), 0)
	println(angleClock(1, 57), 76.5)
}

/*
https://leetcode.com/problems/angle-between-hands-of-a-clock/

Given two numbers, hour and minutes. Return the smaller angle (in degrees) formed between the hour and the minute hand.

Example 1:

Input: hour = 12, minutes = 30
Output: 165

Example 2:

Input: hour = 3, minutes = 30
Output: 75

Example 3:

Input: hour = 3, minutes = 15
Output: 7.5

Example 4:

Input: hour = 4, minutes = 50
Output: 155

Example 5:

Input: hour = 12, minutes = 0
Output: 0

Constraints:

1 <= hour <= 12
0 <= minutes <= 59
Answers within 10^-5 of the actual value will be accepted as correct.
*/

func angleClock(hour int, minutes int) float64 {
	minuteAngle := 6 * float64(minutes)
	hourAngle := 30*float64(hour%12) + float64(minutes)/2

	a := math.Abs(minuteAngle - hourAngle)
	if a > 180 {
		return 360 - a
	}

	return a
}

package main

import (
	"container/heap"
	"sort"
)

func main() {
	println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}}), 3)
	println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}), 4)
	println(maxEvents([][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}), 4)
}

/*
https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/

Given an array of events where events[i] = [startDay[i], endDay[i]]. Every event i starts at startDay[i] and ends
at endDay[i].
You can attend an event i at any day d where startTime[i] <= d <= endTime[i]. Notice that you can only attend
one event at any time d.

Return the maximum number of events you can attend.

Example 1:

Input: events = [[1,2],[2,3],[3,4]]
Output: 3
Explanation: You can attend all the three events.
One way to attend them all is as shown.
Attend the first event on day 1.
Attend the second event on day 2.
Attend the third event on day 3.
*/

type eventsHeap [][]int

func (eh eventsHeap) Len() int {
	return len(eh)
}

func (eh eventsHeap) Less(i, j int) bool {
	return eh[i][1] < eh[j][1]
}

func (eh eventsHeap) Swap(i, j int) {
	eh[i], eh[j] = eh[j], eh[i]
}

func (eh *eventsHeap) Push(x interface{}) {
	item := x.([]int)
	*eh = append(*eh, item)
}

func (eh *eventsHeap) Pop() interface{} {
	item := (*eh)[len(*eh)-1]
	*eh = (*eh)[:len(*eh)-1]

	return item
}

func maxEvents(events [][]int) int {
	if len(events) <= 1 {
		return len(events)
	}

	// sort by starting day
	sort.SliceStable(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	maxDay := 0
	for _, event := range events {
		if maxDay < event[1] {
			maxDay = event[1]
		}
	}

	hp := &eventsHeap{}

	cnt := 0
	eIdx := 0
	for day := 1; day <= maxDay; {
		if hp.Len() == 0 && eIdx == len(events) {
			break
		}

		for eIdx < len(events) && events[eIdx][0] == day {
			heap.Push(hp, events[eIdx])
			eIdx++
		}

		if hp.Len() == 0 {
			day++
			continue
		}

		event := heap.Pop(hp).([]int)
		for event[1] < day && hp.Len() > 0 {
			event = heap.Pop(hp).([]int)
		}

		if event[1] >= day {
			cnt++
		}
		day++
	}

	return cnt
}

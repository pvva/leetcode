package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}),
		[]string{"JFK", "MUC", "LHR", "SFO", "SJC"})
	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}),
		[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"})
	fmt.Println(findItinerary([][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}),
		[]string{"JFK", "NRT", "JFK", "KUL"})

	fmt.Println(findItinerary2([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}),
		[]string{"JFK", "MUC", "LHR", "SFO", "SJC"})
	fmt.Println(findItinerary2([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}),
		[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"})
	fmt.Println(findItinerary2([][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}),
		[]string{"JFK", "NRT", "JFK", "KUL"})
}

/*
https://leetcode.com/problems/reconstruct-itinerary/

Given a list of airline tickets represented by pairs of departure and arrival airports [from, to], reconstruct
the itinerary in order. All of the tickets belong to a man who departs from JFK. Thus, the itinerary must begin with JFK.

Note:

If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order
when read as a single string. For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].
All airports are represented by three capital letters (IATA code).
You may assume all tickets form at least one valid itinerary.
One must use all the tickets once and only once.

Example 1:

Input: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
Output: ["JFK", "MUC", "LHR", "SFO", "SJC"]

Example 2:

Input: [["JFK","SFO"], ["JFK","ATL"], ["SFO","ATL"], ["ATL","JFK"], ["ATL","SFO"]]
Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"].
             But it is larger in lexical order.
*/

type destinationPoint struct {
	name    string
	visited bool
}

// O(E*log(E))
func findItinerary(tickets [][]string) []string {
	paths := make(map[string]*[]destinationPoint)

	for _, ticket := range tickets {
		tos, ex := paths[ticket[0]]
		if !ex {
			tos = &([]destinationPoint{})
			paths[ticket[0]] = tos
		}
		*tos = append(*tos, destinationPoint{name: ticket[1], visited: false})
	}
	for _, tos := range paths {
		sort.SliceStable(*tos, func(i, j int) bool {
			return (*tos)[i].name < (*tos)[j].name
		})
	}

	result := []string{"JFK"}
	findNextStep("JFK", paths, &result, len(tickets)+1)

	return result
}

func findNextStep(from string, paths map[string]*[]destinationPoint, path *[]string, maxlen int) bool {
	if len(*path) == maxlen {
		return true
	}
	tos := paths[from]
	if tos == nil {
		return false
	}

	visited := false
	for i := 0; i < len(*tos); i++ {
		if (*tos)[i].visited {
			continue
		}
		to := (*tos)[i].name
		(*tos)[i].visited = true
		*path = append(*path, to)

		if !findNextStep(to, paths, path, maxlen) {
			*path = (*path)[:len(*path)-1]
			(*tos)[i].visited = false
		} else {
			visited = true
		}
	}

	return visited
}

// this version is slower because of 2 heap operations - push and pop

type pointsHeap []string

func (ph pointsHeap) Len() int {
	return len(ph)
}

func (ph pointsHeap) Less(i, j int) bool {
	return ph[i] < ph[j]
}

func (ph pointsHeap) Swap(i, j int) {
	ph[i], ph[j] = ph[j], ph[i]
}

func (ph *pointsHeap) Push(x interface{}) {
	item := x.(string)
	*ph = append(*ph, item)
}

func (ph *pointsHeap) Pop() interface{} {
	item := (*ph)[len(*ph)-1]
	*ph = (*ph)[:len(*ph)-1]

	return item
}

// O(E*log(E))
func findItinerary2(tickets [][]string) []string {
	paths := make(map[string]*pointsHeap)

	for _, ticket := range tickets {
		tos, ex := paths[ticket[0]]
		if !ex {
			tos = &pointsHeap{}
			paths[ticket[0]] = tos
		}
		heap.Push(tos, ticket[1])
	}

	result := []string{}
	itineraryDfs("JFK", paths, &result)
	result = append(result, "JFK")

	l := 0
	r := len(result) - 1
	for l < r {
		result[l], result[r] = result[r], result[l]
		l++
		r--
	}

	return result
}

func itineraryDfs(from string, paths map[string]*pointsHeap, result *[]string) {
	tos := paths[from]
	if tos == nil {
		return
	}

	for tos.Len() > 0 {
		to := heap.Pop(tos).(string)
		itineraryDfs(to, paths, result)
		*result = append(*result, to)
	}
}

package main

import (
	"container/heap"
)

func main() {
	println(trapRainWater([][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}), 4)
	println(trapRainWater([][]int{
		{12, 13, 1, 12},
		{13, 4, 13, 12},
		{13, 8, 10, 12},
		{12, 13, 12, 12},
		{13, 13, 13, 13},
	}), 14)
	println(trapRainWater([][]int{
		{2, 3, 4},
		{5, 6, 7},
		{8, 9, 10},
		{11, 12, 13},
		{14, 15, 16},
	}), 0)
}

/*
https://leetcode.com/problems/trapping-rain-water-ii/

Given an m x n matrix of positive integers representing the height of each unit cell in a 2D elevation map,
compute the volume of water it is able to trap after raining.

Example:

Given the following 3x6 height map:
[
  [1,4,3,1,3,2],
  [3,2,1,3,2,4],
  [2,3,3,2,3,1]
]

Return 4.
*/

type mapElement struct {
	v    int
	x, y int
}

type rwPQ []mapElement

func (h rwPQ) Len() int {
	return len(h)
}

func (h rwPQ) Less(i, j int) bool {
	return h[i].v < h[j].v
}

func (h rwPQ) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *rwPQ) Push(x interface{}) {
	item := x.(mapElement)
	*h = append(*h, item)
}

func (h *rwPQ) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return item
}

func trapRainWater(heightMap [][]int) int {
	xLen := len(heightMap)
	if xLen <= 2 {
		return 0
	}
	yLen := len(heightMap[0])
	if yLen <= 2 {
		return 0
	}

	visitedMap := make([][]bool, xLen)
	for x := range visitedMap {
		visitedMap[x] = make([]bool, yLen)
	}

	pq := &rwPQ{}
	for y := 0; y < yLen; y++ {
		heap.Push(pq, mapElement{v: heightMap[0][y], x: 0, y: y})
		heap.Push(pq, mapElement{v: heightMap[xLen-1][y], x: xLen - 1, y: y})
		visitedMap[0][y] = true
		visitedMap[xLen-1][y] = true
	}
	for x := 1; x < xLen-1; x++ {
		heap.Push(pq, mapElement{v: heightMap[x][0], x: x, y: 0})
		heap.Push(pq, mapElement{v: heightMap[x][yLen-1], x: x, y: yLen - 1})
		visitedMap[x][0] = true
		visitedMap[x][yLen-1] = true
	}

	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	total := 0
	currentMax := 0
	for pq.Len() > 0 {
		me := heap.Pop(pq).(mapElement)
		if currentMax < me.v {
			currentMax = me.v
		}

		for _, dir := range directions {
			nx := me.x + dir[0]
			ny := me.y + dir[1]

			if nx >= 0 && nx < xLen && ny >= 0 && ny < yLen && !visitedMap[nx][ny] {
				visitedMap[nx][ny] = true
				v := heightMap[nx][ny]
				heap.Push(pq, mapElement{v: v, x: nx, y: ny})

				if v < currentMax {
					total += currentMax - v
				}
			}
		}
	}

	return total
}

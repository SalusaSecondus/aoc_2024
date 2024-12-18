package main

import (
	"container/heap"
	"strconv"
	"strings"
)

func day18_1(lines []string) int {
	g := parse18(lines)
	dest := [2]int{g.MaxX, g.MaxY}
	maxAge := 12
	if g.MaxX > 6 {
		maxAge = 1024
	}
	return dijk18(g, [2]int{0, 0}, dest, maxAge)
}

func day18_2(lines []string) string {
	g := parse18(lines)
	dest := [2]int{g.MaxX, g.MaxY}
	maxAge := 12
	if g.MaxX > 6 {
		maxAge = 1024
	}
	for dijk18(g, [2]int{0, 0}, dest, maxAge) > 0 {
		maxAge++
	}

	return lines[maxAge-1]
}

func parse18(lines []string) Grid[int] {
	result := NewGrid[int]()
	result.Default = -1
	for idx, l := range lines {
		parts := strings.Split(l, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		result.Set(x, y, idx+1)
	}
	result.MinX = 0
	result.MinY = 0
	if len(lines) > 25 {
		result.MaxX = 70
		result.MaxY = 70
	} else {
		result.MaxX = 6
		result.MaxY = 6
	}
	return result
}

func dijk18(g Grid[int], start, end Coord, maxTime int) int {
	queue := make(StateQueue, 1)
	queue[0] = &State16{
		locs: []Location16{Location16{coord: start}},
		cost: 0,
	}
	heap.Init(&queue)

	visited := make(map[Location16]bool)

	for queue.Len() > 0 {
		curr := heap.Pop(&queue).(*State16)
		// fmt.Printf("%d\tQ = %d, Pathlen = %d\n", maxTime, queue.Len(), curr.cost)
		if curr.Last().coord == end {
			return curr.cost
		}
		visited[curr.Last()] = true

		for _, link := range Connected18(g, curr.Last(), maxTime) {
			if !visited[link.dst] {
				nextLocs := make([]Location16, len(curr.locs))
				copy(nextLocs, curr.locs)
				nextLocs = append(nextLocs, link.dst)
				next := State16{
					locs: nextLocs,
					cost: curr.cost + link.cost,
				}
				heap.Push(&queue, &next)
				visited[link.dst] = true
			}
		}
	}
	return -1
}

func Connected18(g Grid[int], loc Location16, maxAge int) []Link16 {
	result := make([]Link16, 0)

	up := Up15.next(loc.coord)
	down := Down15.next(loc.coord)
	left := Left15.next(loc.coord)
	right := Right15.next(loc.coord)

	upVal := g.GetC(up)
	downVal := g.GetC(down)
	leftVal := g.GetC(left)
	rightVal := g.GetC(right)

	if g.InRangeC(up) && (upVal == -1 || upVal > maxAge) {
		result = append(result, Link16{src: loc, dst: Location16{coord: up}, cost: 1})
	}
	if g.InRangeC(down) && (downVal == -1 || downVal > maxAge) {
		result = append(result, Link16{src: loc, dst: Location16{coord: down}, cost: 1})
	}
	if g.InRangeC(left) && (leftVal == -1 || leftVal > maxAge) {
		result = append(result, Link16{src: loc, dst: Location16{coord: left}, cost: 1})
	}
	if g.InRangeC(right) && (rightVal == -1 || rightVal > maxAge) {
		result = append(result, Link16{src: loc, dst: Location16{coord: right}, cost: 1})
	}

	return result
}

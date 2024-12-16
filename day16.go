package main

import (
	"container/heap"
)

func day16_1(lines []string) int {
	grid, reindeer, end := parse16(lines)
	return dijk16(grid, reindeer, end)
}

func day16_2(lines []string) int {
	grid, reindeer, end := parse16(lines)
	return len(dijk16_2(grid, reindeer, end))
}

type Location16 struct {
	coord     Coord
	direction Dir15
}

func parse16(lines []string) (Grid[byte], Location16, Coord) {
	grid := NewGrid[byte]()
	grid.Default = '.'
	reindeer := Location16{direction: Right15}
	end := Coord{}

	for y, l := range lines {
		for x := range l {
			item := l[x]
			if item == 'S' {
				reindeer.coord = [2]int{x, y}
				grid.Set(x, y, '.')
			} else if item == 'E' {
				end = [2]int{x, y}
				grid.Set(x, y, '.')
			} else {
				grid.Set(x, y, item)
			}
		}
	}
	return grid, reindeer, end
}

type Link16 struct {
	src  Location16
	dst  Location16
	cost int
}

func dijk16(g Grid[byte], start Location16, end Coord) int {
	queue := make(StateQueue, 1)
	queue[0] = &State16{
		locs: []Location16{start},
		cost: 0,
	}
	heap.Init(&queue)

	visited := make(map[Location16]bool)

	for queue.Len() > 0 {
		curr := heap.Pop(&queue).(*State16)
		if curr.Last().coord == end {
			return curr.cost
		}
		visited[curr.Last()] = true

		for _, link := range Connected(g, curr.Last()) {
			if !visited[link.dst] {
				nextLocs := make([]Location16, len(curr.locs))
				copy(nextLocs, curr.locs)
				nextLocs = append(nextLocs, link.dst)
				next := State16{
					locs: nextLocs,
					cost: curr.cost + link.cost,
				}
				heap.Push(&queue, &next)
			}
		}
	}
	panic("unreachable")
}

func dijk16_2(g Grid[byte], start Location16, end Coord) map[Coord]bool {
	queue := make(StateQueue, 1)
	queue[0] = &State16{
		locs: []Location16{start},
		cost: 0,
	}
	heap.Init(&queue)

	bestCost := -1
	seats := make(map[Coord]bool)
	visited := make(map[Location16]int)

	for queue.Len() > 0 {
		curr := heap.Pop(&queue).(*State16)
		// fmt.Printf("Queue: %d\tCost %d\n", queue.Len(), curr.cost)
		if curr.Last().coord == end {
			// panic("Found!")
			if bestCost < 0 {
				bestCost = curr.cost
			}
			for _, l := range curr.locs {
				seats[l.coord] = true
			}
		}
		if bestCost > 0 && curr.cost > bestCost {
			return seats
		}
		visited[curr.Last()] = curr.cost

		for _, link := range Connected(g, curr.Last()) {
			priorCost, found := visited[link.dst]
			if !found || priorCost >= curr.cost+link.cost {

				nextLocs := make([]Location16, len(curr.locs))
				copy(nextLocs, curr.locs)
				nextLocs = append(nextLocs, link.dst)
				next := State16{
					locs: nextLocs,
					cost: curr.cost + link.cost,
				}
				heap.Push(&queue, &next)
			}
		}
	}
	return seats
}

func Connected(g Grid[byte], loc Location16) []Link16 {
	result := make([]Link16, 0)
	dst1 := loc
	dst2 := loc
	switch loc.direction {
	case Up15:
		dst1.direction = Left15
		dst2.direction = Right15
	case Down15:
		dst1.direction = Left15
		dst2.direction = Right15
	case Left15:
		dst1.direction = Up15
		dst2.direction = Down15
	case Right15:
		dst1.direction = Up15
		dst2.direction = Down15
	}
	result = append(result, Link16{src: loc, dst: dst1, cost: 1000})
	result = append(result, Link16{src: loc, dst: dst2, cost: 1000})

	nextStep := loc.direction.next(loc.coord)
	nextTile := g.GetC(nextStep)
	dst3 := Location16{
		coord:     nextStep,
		direction: loc.direction,
	}
	if nextTile == '.' {
		result = append(result, Link16{src: loc, dst: dst3, cost: 1})
	}
	return result
}

type State16 struct {
	locs []Location16
	cost int
}

func (s State16) Last() Location16 {
	return s.locs[len(s.locs)-1]
}

type StateQueue []*State16

func (sq StateQueue) Len() int {
	return len(sq)
}

func (sq StateQueue) Less(i, j int) bool {
	return sq[i].cost < sq[j].cost
}

func (sq StateQueue) Swap(i, j int) {
	sq[i], sq[j] = sq[j], sq[i]
}

func (sq *StateQueue) Push(x any) {
	item := x.(*State16)
	*sq = append(*sq, item)
}

func (sq *StateQueue) Pop() any {
	old := *sq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*sq = old[0 : n-1]
	return item
}

package main

import (
	"container/heap"
	"fmt"
	"slices"
)

func day20_1(lines []string) int {
	grid, start, end := parse20(lines)
	savingTarget := 100
	if grid.MaxX < 20 {
		savingTarget = 36
	}

	longest := dijk20(grid, start, end)
	paths := stack20(grid, start, end, longest, 2, savingTarget)

	return paths
}

func day20_2(lines []string) int {
	grid, start, end := parse20(lines)
	savingTarget := 100
	if grid.MaxX < 20 {
		savingTarget = 66
	}

	longest := dijk20(grid, start, end)
	paths := stack20(grid, start, end, longest, 20, savingTarget)

	return paths
}

func pathToString(path State20) string {
	result := ""

	for idx, c := range path.locs {
		if result != "" {
			result += "->"
		}
		if idx == path.cheat || idx == path.cheat+1 {
			result += "\033[1m" + c.String() + "\033[0m"
		} else {
			result += c.String()
		}
	}

	return result
}

func parse20(lines []string) (Grid[byte], Coord, Coord) {
	grid := NewGrid[byte]()
	grid.Default = '#'
	var start Coord
	var end Coord
	for y, l := range lines {
		for x := range l {
			val := l[x]
			if val == 'S' {
				start = [2]int{x, y}
				val = '.'
			} else if val == 'E' {
				end = [2]int{x, y}
				val = '.'
			}
			grid.Set(x, y, val)
		}
	}
	return grid, start, end
}

type State20 struct {
	locs  []Coord
	cost  int
	cheat int
}

func (s State20) Last() Coord {
	idx := len(s.locs) - 1
	return s.locs[idx]
}

type VKey20 struct {
	coord Coord
	cheat int
}

func (s State20) VKey() VKey20 {
	cheat := 0
	if s.cheat == len(s.locs)-1 {
		cheat = 1
	} else if s.cheat != NO_CHEAT {
		cheat = 2
	}
	return VKey20{
		coord: s.Last(),
		cheat: cheat,
	}
}

const NO_CHEAT = -2

func dijk20(g Grid[byte], start, end Coord) State20 {
	queue := make(State20Queue, 1)
	queue[0] = &State20{
		locs:  []Coord{start},
		cost:  0,
		cheat: -3,
	}
	heap.Init(&queue)

	// visited := make(map[VKey20]*State20)
	// result := make([]State20, 0)

	for queue.Len() > 0 {
		curr := heap.Pop(&queue).(*State20)
		// fmt.Printf("F = %d\tQ = %d, Pathlen = %d\n", len(result), queue.Len(), curr.cost)
		if curr.Last() == end {
			// result = append(result, *curr)
			if curr.cheat == -3 {
				return *curr
			}
			// visited[curr.VKey()] = curr
			continue
		}

		for _, link := range Connected20(g, curr, 0) {
			// if visited[link.VKey()] == nil {
			heap.Push(&queue, &link)
			// 	visited[link.VKey()] = &link
			// } else {
			// 	fmt.Printf("vvvvvvvv\nDiscarding path due to %+v\nOLD: %s\nNEW: %s\n^^^^^^^^\n\n",
			// 		link.VKey(), pathToString(*visited[link.VKey()]), pathToString(link))
			// }
		}
	}
	panic("Unreachable")
}

func stack20(g Grid[byte], start, end Coord, normalPath State20, cheatCount, savingTarget int) int {
	queue := make(State20Stack, 1)
	queue[0] = &State20{
		locs:  []Coord{start},
		cost:  0,
		cheat: NO_CHEAT,
	}
	heap.Init(&queue)

	shortcuts := make(map[Coord]int)
	for idx, c := range normalPath.locs {
		shortcuts[c] = normalPath.cost - idx
	}
	maxLen := normalPath.cost - savingTarget + 1

	cheats := make(map[[2]Coord]int)

	// visited := make(map[VKey20]*State20)
	// result := 0

	for queue.Len() > 0 {
		curr := heap.Pop(&queue).(*State20)
		if curr.Last() == end {
			cheats[curr.CheatSig(cheatCount)]++
			fmt.Printf("F = %d\tQ = %d, Pathlen = %d\n", len(cheats), queue.Len(), curr.cost)
			// if curr.cheat == NO_CHEAT {
			// 	return result
			// }
			// visited[curr.VKey()] = curr
			continue
		}

		for _, link := range Connected20(g, curr, cheatCount) {
			// if visited[link.VKey()] == nil {
			if link.DoneCheating(cheatCount) {
				shortcut, found := shortcuts[link.Last()]
				if found && shortcut+link.cost < maxLen {
					cheats[link.CheatSig(cheatCount)]++
				}
			} else if link.cost < maxLen {
				heap.Push(&queue, &link)
			}
			// 	visited[link.VKey()] = &link
			// } else {
			// 	fmt.Printf("vvvvvvvv\nDiscarding path due to %+v\nOLD: %s\nNEW: %s\n^^^^^^^^\n\n",
			// 		link.VKey(), pathToString(*visited[link.VKey()]), pathToString(link))
			// }
		}
	}
	return len(cheats)
}

func Connected20(g Grid[byte], curr *State20, cheatCount int) []State20 {
	result := make([]State20, 0)

	// if curr.Last()[0] == 7 && curr.Last()[1] == 8 {
	// 	fmt.Println("FOO!")
	// }

	up := Up15.next(curr.Last())
	down := Down15.next(curr.Last())
	left := Left15.next(curr.Last())
	right := Right15.next(curr.Last())

	result = maybeAdd(result, g, *curr, up, cheatCount)
	result = maybeAdd(result, g, *curr, down, cheatCount)
	result = maybeAdd(result, g, *curr, left, cheatCount)
	result = maybeAdd(result, g, *curr, right, cheatCount)

	return result
}

func (s State20) CheatSig(cheatCount int) [2]Coord {

	if s.CanCheat() {
		oob := [2]int{-1, -1}
		return [2]Coord{oob, oob}
	} else if s.IsCheating(cheatCount) {
		return [2]Coord{
			s.locs[s.cheat],
			s.locs[len(s.locs)-1],
		}
	} else {
		return [2]Coord{
			s.locs[s.cheat],
			s.locs[s.cheat+cheatCount-1],
		}
	}
}

func (s State20) CanCheat() bool {
	return s.cheat == NO_CHEAT
}

func (s State20) IsCheating(cheatCount int) bool {
	return !s.CanCheat() && s.cheat+cheatCount-1 > len(s.locs)
}

func (s State20) DoneCheating(cheatCount int) bool {
	return !s.CanCheat() && !s.IsCheating(cheatCount+1)
}
func maybeAdd(result []State20, g Grid[byte], curr State20, coord Coord, cheatCount int) []State20 {
	if slices.Contains(curr.locs, coord) {
		return result
	}
	if !g.InRangeC(coord) {
		return result
	}
	// canary := [2]int{7, 8}
	// if slices.Contains(curr.locs, canary) {
	// 	fmt.Println("FOO: " + pathToString(curr) + " ? " + coord.String())
	// }
	val := g.GetC(coord)
	if val == '.' {
		locs := slices.Clone(curr.locs)

		return append(result, State20{
			locs:  append(locs, coord),
			cheat: curr.cheat,
			cost:  curr.cost + 1,
		})
	} else if curr.CanCheat() {
		locs := slices.Clone(curr.locs)

		return append(result, State20{
			locs:  append(locs, coord),
			cheat: len(curr.locs),
			cost:  curr.cost + 1,
		})
	} else if curr.IsCheating(cheatCount) {
		locs := slices.Clone(curr.locs)

		return append(result, State20{
			locs:  append(locs, coord),
			cheat: curr.cheat,
			cost:  curr.cost + 1,
		})
	} else {
		return result
	}
}

type State20Queue []*State20

func (sq State20Queue) Len() int {
	return len(sq)
}

func (sq State20Queue) Less(i, j int) bool {
	if sq[i].cost < sq[j].cost {
		return true
	} else if sq[i].cost == sq[j].cost && sq[i].cheat < sq[j].cheat {
		return true
	}
	return false
}

func (sq State20Queue) Swap(i, j int) {
	sq[i], sq[j] = sq[j], sq[i]
}

func (sq *State20Queue) Push(x any) {
	item := x.(*State20)
	*sq = append(*sq, item)
}

func (sq *State20Queue) Pop() any {
	old := *sq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*sq = old[0 : n-1]
	return item
}

type State20Stack []*State20

func (sq State20Stack) Len() int {
	return len(sq)
}

func (sq State20Stack) Less(i, j int) bool {
	if sq[i].cost > sq[j].cost {
		return true
	} else if sq[i].cost == sq[j].cost && sq[i].cheat < sq[j].cheat {
		return true
	}
	return false
}

func (sq State20Stack) Swap(i, j int) {
	sq[i], sq[j] = sq[j], sq[i]
}

func (sq *State20Stack) Push(x any) {
	item := x.(*State20)
	*sq = append(*sq, item)
}

func (sq *State20Stack) Pop() any {
	old := *sq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*sq = old[0 : n-1]
	return item
}

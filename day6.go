package main

import (
	"fmt"
)

type Thing6 byte

const (
	None6    Thing6     = '.'
	Block6   Thing6     = '#'
	Visited6 Thing6     = 'X'
	UpT6     Direction6 = '^'
	DownT6   Direction6 = 'v'
	LeftT6   Direction6 = '<'
	RightT6  Direction6 = '>'
)

func (t Thing6) String() string {
	return fmt.Sprintf("%c", t)
}

type Direction6 byte

const (
	Up6    Direction6 = '^'
	Down6  Direction6 = 'v'
	Left6  Direction6 = '<'
	Right6 Direction6 = '>'
)

func day6_1(input []string) int {
	var result int
	grid, guard := parse6(input)
	for !guard.step(grid) {
		// fmt.Println(grid)
		// fmt.Println()
	}
	for _, v := range grid.Elements {
		if v == Visited6 {
			result++
		}
	}
	fmt.Println(grid)
	fmt.Println(len(grid.Elements))

	return result
}

func day6_2(input []string) int {
	var result int
	baseGrid, baseGuard := parse6(input)
	for !baseGuard.step(baseGrid) {
	}

	for coord, elem := range baseGrid.Elements {
		if elem != Visited6 {
			continue
		}
		grid, guard := parse6(input)
		grid.Elements[coord] = Block6
		// fmt.Printf("Trying\n%s\n", grid)

		for {
			done, loop := guard.step2(grid)
			if done {
				if loop {
					result++
					// fmt.Println(grid)
					// fmt.Println()
				}
				break
			}
		}
	}
	return result
}

type Guard struct {
	Loc Coord
	Dir Direction6
}

func parse6(lines []string) (Grid[Thing6], Guard) {
	result := NewGrid[Thing6]()
	result.Default = None6
	guard := Guard{}
	for y, line := range lines {
		for x, char := range line {
			elem := None6
			if char == '#' {
				elem = Block6
			} else if char != '.' {
				guard.Dir = Direction6(char)
				guard.Loc = [2]int{x, y}
				elem = Visited6
			}
			if elem != None6 {
				result.Set(x, y, elem)
			}
		}
	}
	return result, guard
}

func (g *Guard) step(grid Grid[Thing6]) bool {
	xStep := 0
	yStep := 0
	nextDir := Up6
	switch g.Dir {
	case Up6:
		yStep = -1
		nextDir = Right6
	case Down6:
		yStep = 1
		nextDir = Left6
	case Right6:
		xStep = 1
		nextDir = Down6
	case Left6:
		xStep = -1
		nextDir = Up6
	}

	newX := g.Loc[0] + xStep
	newY := g.Loc[1] + yStep
	if grid.Get(newX, newY) == Block6 {
		g.Dir = nextDir
		return false
	}
	if newX < grid.MinX || newX > grid.MaxX || newY < grid.MinY || newY > grid.MaxY {
		return true
	}
	g.Loc = [2]int{newX, newY}
	grid.Set(newX, newY, Visited6)
	return false
}

func (g *Guard) step2(grid Grid[Thing6]) (bool, bool) {
	xStep := 0
	yStep := 0
	nextDir := Up6
	switch g.Dir {
	case Up6:
		yStep = -1
		nextDir = Right6
	case Down6:
		yStep = 1
		nextDir = Left6
	case Right6:
		xStep = 1
		nextDir = Down6
	case Left6:
		xStep = -1
		nextDir = Up6
	}

	newX := g.Loc[0] + xStep
	newY := g.Loc[1] + yStep

	nextTile := grid.Get(newX, newY)
	if nextTile == Block6 {
		g.Dir = nextDir
		return false, false
	}
	if nextTile == Thing6(g.Dir) {
		return true, true
	}

	if newX < grid.MinX || newX > grid.MaxX || newY < grid.MinY || newY > grid.MaxY {
		return true, false
	}
	g.Loc = [2]int{newX, newY}
	grid.Set(newX, newY, Thing6(g.Dir))
	return false, false
}

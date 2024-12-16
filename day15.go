package main

import "fmt"

func day15_1(lines []string) int {
	grid, robot, dirs := parse15(lines)

	fmt.Printf("Initial state:\n%s\n", grid)
	for _, d := range dirs {
		if maybeMove(grid, robot, d) {
			robot = d.next(robot)
		}
		fmt.Printf("Move %c:\n%s\n", d, grid)
	}

	return score15(grid)
}

func day15_2(lines []string) int {
	grid, robot, dirs := parse15_2(lines)
	fmt.Printf("Initial state:\n%s\n", grid)
	for _, d := range dirs {
		if canMove2(grid, robot, d) {
			move2(grid, robot, d)
			robot = d.next(robot)
		}
		fmt.Printf("Move %c:\n%s\n", d, grid)
	}

	return score15(grid)
}

type Dir15 byte

const (
	Up15    Dir15 = '^'
	Down15  Dir15 = 'v'
	Left15  Dir15 = '<'
	Right15 Dir15 = '>'
)

func (d Dir15) next(c Coord) Coord {
	switch d {
	case Up15:
		return Coord([2]int{c[0], c[1] - 1})
	case Down15:
		return Coord([2]int{c[0], c[1] + 1})
	case Right15:
		return Coord([2]int{c[0] + 1, c[1]})
	case Left15:
		return Coord([2]int{c[0] - 1, c[1]})
	}
	panic("Unsupported direction")
}

func maybeMove(grid Grid[string], c Coord, d Dir15) bool {
	thing := grid.GetC(c)
	if thing == "." {
		return true
	} else if thing == "#" {
		return false
	} else if maybeMove(grid, d.next(c), d) {
		grid.SetC(d.next(c), thing)
		grid.SetC(c, ".")
		return true
	} else {
		return false
	}
}

func score15(g Grid[string]) int {
	result := 0
	for c, item := range g.Elements {
		if item == "O" || item == "[" {
			result += c[0]
			result += 100 * c[1]
		}
	}
	return result
}

func parse15(lines []string) (Grid[string], Coord, []Dir15) {
	g := NewGrid[string]()
	g.Default = "."
	robot := Coord([2]int{0, 0})
	dirs := make([]Dir15, 0)

	y := 0
	for _, l := range lines {
		if len(l) > 1 && l[0] == '#' {
			// We're in the grid
			for x := range l {
				g.Set(x, y, fmt.Sprintf("%c", l[x]))
				if l[x] == '@' {
					robot = Coord([2]int{x, y})
					g.Set(x, y, fmt.Sprintf("\033[0;31m%c\033[0m", l[x]))
				}
			}
			y++
		} else {
			// We're in the instructions
			for _, i := range l {
				if i == '<' || i == '>' || i == '^' || i == 'v' {
					dirs = append(dirs, Dir15(i))
				}
			}
		}
	}

	return g, robot, dirs
}

func parse15_2(lines []string) (Grid[string], Coord, []Dir15) {
	g := NewGrid[string]()
	g.Default = "."
	robot := Coord([2]int{0, 0})
	dirs := make([]Dir15, 0)

	y := 0
	for _, l := range lines {
		if len(l) > 1 && l[0] == '#' {
			// We're in the grid
			for x, thing := range l {
				switch thing {
				case '.':
					g.Set(2*x, y, ".")
					g.Set(2*x+1, y, ".")
				case '#':
					g.Set(2*x, y, "#")
					g.Set(2*x+1, y, "#")
				case 'O':
					g.Set(2*x, y, "[")
					g.Set(2*x+1, y, "]")
				case '@':
					g.Set(2*x, y, "\033[0;31m@\033[0m")
					g.Set(2*x+1, y, ".")
					robot = Coord([2]int{2 * x, y})
				}
			}
			y++
		} else {
			// We're in the instructions
			for _, i := range l {
				if i == '<' || i == '>' || i == '^' || i == 'v' {
					dirs = append(dirs, Dir15(i))
				}
			}
		}
	}

	return g, robot, dirs
}

func canMove2(grid Grid[string], c Coord, d Dir15) bool {
	thing := grid.GetC(c)
	if thing == "." {
		return true
	} else if thing == "#" {
		return false
	} else if d == Left15 || d == Right15 {
		return canMove2(grid, d.next(c), d)
	} else {
		// Up or down
		nextStep := d.next(c)
		if !canMove2(grid, nextStep, d) {
			return false
		}
		offset := nextStep
		if thing == "[" {
			offset[0]++
		} else if thing == "]" {
			offset[0]--
		} else {
			return true
		}
		return canMove2(grid, offset, d)
	}
}

func move2(grid Grid[string], c Coord, d Dir15) {
	thing := grid.GetC(c)
	nextStep := d.next(c)
	if thing == "#" || thing == "." {
		return
	}
	if d == Left15 || d == Right15 {
		move2(grid, nextStep, d)
		grid.SetC(nextStep, thing)
		grid.SetC(c, ".")
		return
	}
	// Up down
	move2(grid, nextStep, d)
	grid.SetC(nextStep, thing)
	grid.SetC(c, ".")

	selfOffset := c
	offset := nextStep
	if thing == "[" {
		selfOffset[0]++
		offset[0]++
	} else if thing == "]" {
		selfOffset[0]--
		offset[0]--
	} else {
		return
	}
	move2(grid, offset, d)
	grid.SetC(offset, grid.GetC(selfOffset))
	grid.SetC(selfOffset, ".")
}

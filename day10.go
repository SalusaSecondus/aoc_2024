package main

func day10_1(lines []string) int {
	result := 0
	grid, starts := parse10(lines)
	for _, curr := range starts {
		result += findTrails(curr, grid)
	}
	return result
}

func day10_2(lines []string) int {
	result := 0
	grid, starts := parse10(lines)
	for _, curr := range starts {
		result += findAllTrails(curr, grid)
	}
	return result
}

func parse10(lines []string) (Grid[byte], []Coord) {
	result := NewGrid[byte]()
	starts := make([]Coord, 0)
	for y := 0; y < len(lines); y++ {
		row := lines[y]
		for x := 0; x < len(row); x++ {
			result.Set(x, y, row[x]-'0')
			if row[x] == '0' {
				starts = append(starts, [2]int{x, y})
			}
		}
	}
	return result, starts
}

func findTrails(start Coord, g Grid[byte]) int {
	toCheck := make(map[Coord]int)
	toCheck[start] = 1
	// fmt.Printf("Trailhead %v\n", start)
	for nextHeight := byte(1); nextHeight <= 9; nextHeight++ {
		nextPoints := make(map[Coord]int)
		// fmt.Printf("\t%d<<\n", nextHeight)
		for curr, _ := range toCheck {
			x := curr[0]
			y := curr[1]
			if g.Get(x, y-1) == nextHeight {
				// fmt.Printf("\t\t(%d, %d)\n", x, y-1)
				nextPoints[[2]int{x, y - 1}] = 1
			}
			if g.Get(x, y+1) == nextHeight {
				// fmt.Printf("\t\t(%d, %d)\n", x, y+1)
				nextPoints[[2]int{x, y + 1}] = 1
			}
			if g.Get(x-1, y) == nextHeight {
				// fmt.Printf("\t\t(%d, %d)\n", x-1, y)
				nextPoints[[2]int{x - 1, y}] = 1
			}
			if g.Get(x+1, y) == nextHeight {
				// fmt.Printf("\t\t(%d, %d)\n", x+1, y)
				nextPoints[[2]int{x + 1, y}] = 1

			}
		}
		toCheck = nextPoints
	}
	return len(toCheck)
}

func findAllTrails(start Coord, g Grid[byte]) int {
	toCheck := make(map[string]Coord)
	toCheck[start.String()] = start
	// fmt.Printf("Trailhead %v\n", start)
	for nextHeight := byte(1); nextHeight <= 9; nextHeight++ {
		nextPoints := make(map[string]Coord)
		// fmt.Printf("\t%d<<\n", nextHeight)
		for path, curr := range toCheck {
			x := curr[0]
			y := curr[1]
			if g.Get(x, y-1) == nextHeight {
				// fmt.Printf("\t\t(%d, %d)\n", x, y-1)
				nextStep := Coord([2]int{x, y - 1})
				nextPoints[path+nextStep.String()] = nextStep
			}
			if g.Get(x, y+1) == nextHeight {
				// fmt.Printf("\t\t(%d, %d)\n", x, y+1)
				nextStep := Coord([2]int{x, y + 1})
				nextPoints[path+nextStep.String()] = nextStep
			}
			if g.Get(x-1, y) == nextHeight {
				// fmt.Printf("\t\t(%d, %d)\n", x-1, y)
				nextStep := Coord([2]int{x - 1, y})
				nextPoints[path+nextStep.String()] = nextStep
			}
			if g.Get(x+1, y) == nextHeight {
				// fmt.Printf("\t\t(%d, %d)\n", x+1, y)
				nextStep := Coord([2]int{x + 1, y})
				nextPoints[path+nextStep.String()] = nextStep

			}
		}
		toCheck = nextPoints
	}
	return len(toCheck)
}

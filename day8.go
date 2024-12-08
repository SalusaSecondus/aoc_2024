package main

func day8_1(lines []string) int {
	grid := toGrid(lines)
	antennas := make(map[string][]Coord)
	for coord, val := range grid.Elements {
		if val != "." {
			prev := antennas[val]
			if prev == nil {
				prev = make([]Coord, 0)
			}
			antennas[val] = append(prev, coord)
		}
	}

	dipoles := make(map[Coord]int)

	for _, locs := range antennas {
		for aIdx := 0; aIdx < len(locs)-1; aIdx++ {
			for bIdx := aIdx + 1; bIdx < len(locs); bIdx++ {
				a := locs[aIdx]
				b := locs[bIdx]
				xDiff := b[0] - a[0]
				yDiff := b[1] - a[1]

				// fmt.Printf("%v and %v have xDiff %d and yDiff %d\n", a, b, xDiff, yDiff)
				d1 := Coord([2]int{b[0] + xDiff, b[1] + yDiff})
				d2 := Coord([2]int{a[0] - xDiff, a[1] - yDiff})
				if inField8(grid, d1) {
					dipoles[d1]++
					// grid.Elements[d1] = "#"
				}
				if inField8(grid, d2) {
					dipoles[d2]++
					// grid.Elements[d2] = "#"
				}
			}
		}
	}
	// fmt.Println(grid)
	return len(dipoles)
}

func inField8(g Grid[string], coord Coord) bool {
	return coord[0] >= g.MinX && coord[0] <= g.MaxX && coord[1] >= g.MinY && coord[1] <= g.MaxY
}

func day8_2(input []string) int {
	grid := toGrid(input)
	antennas := make(map[string][]Coord)
	for coord, val := range grid.Elements {
		if val != "." {
			prev := antennas[val]
			if prev == nil {
				prev = make([]Coord, 0)
			}
			antennas[val] = append(prev, coord)
		}
	}

	lines := make([]Line8, 0)
	dipoles := make(map[Coord]int)

	for _, locs := range antennas {
		for aIdx := 0; aIdx < len(locs)-1; aIdx++ {
			for bIdx := aIdx + 1; bIdx < len(locs); bIdx++ {
				a := locs[aIdx]
				b := locs[bIdx]
				lines = append(lines, FindLine8(a, b))
			}
		}
	}

	for x := grid.MinX; x <= grid.MaxX; x++ {
		for y := grid.MinY; y <= grid.MaxY; y++ {
			point := Coord([2]int{x, y})
			for _, line := range lines {
				if line.Contains(point) {
					dipoles[point]++
					// grid.Elements[point] = "#"
					break
				}
			}
		}
	}
	// fmt.Println(grid)
	return len(dipoles)
}

type Line8 struct {
	Origin Coord
	xDiff  int
	yDiff  int
}

func FindLine8(a, b Coord) Line8 {
	xDiff := b[0] - a[0]
	yDiff := b[1] - a[1]

	if xDiff == 0 && yDiff == 0 {
		panic("a and b must be different")
	}

	return Line8{
		Origin: a,
		xDiff:  xDiff,
		yDiff:  yDiff,
	}
}

func (l Line8) Contains(point Coord) bool {
	xOffset := point[0] - l.Origin[0]
	yOffset := point[1] - l.Origin[1]

	if l.xDiff == 0 {
		return xOffset == 0
	}
	if l.yDiff == 0 {
		return yOffset == 0
	}
	if xOffset == 0 && yOffset == 0 {
		return true
	}
	if xOffset == 0 || yOffset == 0 {
		return false
	}

	xScale := float64(xOffset) / float64(l.xDiff)
	yScale := float64(yOffset) / float64(l.yDiff)
	// return math.Abs(xScale-yScale) < 0.001
	return xScale == yScale
}

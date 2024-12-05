package main

func day4_1(input []string) int {
	grid := toGrid(input)
	result := 0
	target := []string{"X", "M", "A", "S"}

	for x := grid.MinX; x <= grid.MaxX; x++ {
		for y := grid.MinY; y <= grid.MaxY; y++ {

			for xDir := -1; xDir <= 1; xDir++ {
				for yDir := -1; yDir <= 1; yDir++ {
					coord := [2]int{x, y}
					found := 1
					for t := 0; t < len(target); t++ {
						if grid.Elements[coord] != target[t] {
							found = 0
							break
						}
						coord[0] += xDir
						coord[1] += yDir
					}
					result += found
					// fmt.Printf("Found XMAS at (%d, %d) in dir (%d, %d)\n", x, y, xDir, yDir)
				}
			}
		}
	}
	return result
}

func day4_2(input []string) int {
	grid := toGrid(input)
	result := 0

	for x := grid.MinX; x <= grid.MaxX; x++ {
		for y := grid.MinY; y <= grid.MaxY; y++ {
			center := [2]int{x, y}
			nw := [2]int{x - 1, y - 1}
			ne := [2]int{x + 1, y - 1}
			sw := [2]int{x - 1, y + 1}
			se := [2]int{x + 1, y + 1}

			if grid.Elements[center] != "A" {
				continue
			}
			// fmt.Printf("Found "A" at %v\n", center)
			foundOne := (grid.Elements[nw] == "M" && grid.Elements[se] == "S") ||
				(grid.Elements[nw] == "S" && grid.Elements[se] == "M")

			foundTwo := (grid.Elements[ne] == "M" && grid.Elements[sw] == "S") ||
				(grid.Elements[ne] == "S" && grid.Elements[sw] == "M")

			if foundOne && foundTwo {
				result++
			}
		}
	}
	return result
}

package main

import "fmt"

func day12_1(lines []string) int {
	result := 0
	plots, _, _, _ := parse12(lines)
	for _, plot := range plots {
		fmt.Println(plot)
		result += plot.Price()
	}
	return result
}

func day12_2(lines []string) int {
	result := 0
	plots, coordMapping, maxX, maxY := parse12(lines)
	// fmt.Printf("Mapping:\n%v\n", coordMapping)
	plotSides := make([]int, len(plots))

	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			center := Coord([2]int{x, y})
			n := Coord([2]int{x, y - 1})
			s := Coord([2]int{x, y + 1})
			e := Coord([2]int{x + 1, y})
			w := Coord([2]int{x - 1, y})

			nw := Coord([2]int{x - 1, y - 1})
			sw := Coord([2]int{x - 1, y + 1})
			ne := Coord([2]int{x + 1, y - 1})
			se := Coord([2]int{x + 1, y + 1})

			cPlot := coordMapping[center]

			nPlot := coordMapping[n]
			sPlot := coordMapping[s]
			ePlot := coordMapping[e]
			wPlot := coordMapping[w]

			nwPlot := coordMapping[nw]
			nePlot := coordMapping[ne]
			swPlot := coordMapping[sw]
			sePlot := coordMapping[se]

			// fmt.Printf("Center: %s\n\t%d\t%d\t%d\n\t%d\t%d\t%d\n\t%d\t%d\t%d\n",
			// center, nwPlot, nPlot, nePlot, wPlot, cPlot, ePlot, swPlot, sPlot, sePlot)
			// Inside case
			// NE
			if nPlot != cPlot && ePlot != cPlot {
				plotSides[cPlot-1]++
			}
			// NW
			if nPlot != cPlot && wPlot != cPlot {
				plotSides[cPlot-1]++
			}
			// SE
			if sPlot != cPlot && ePlot != cPlot {
				plotSides[cPlot-1]++
			}
			// SW
			if sPlot != cPlot && wPlot != cPlot {
				plotSides[cPlot-1]++
			}

			// Outside case
			// NE
			if nPlot == nePlot && nePlot == ePlot && ePlot != cPlot && nPlot > 0 {
				plotSides[nPlot-1]++
			}
			// NW
			if nPlot == nwPlot && nwPlot == wPlot && wPlot != cPlot && nPlot > 0 {
				plotSides[nPlot-1]++
			}
			// SE
			if sPlot == sePlot && sePlot == ePlot && ePlot != cPlot && sPlot > 0 {
				plotSides[sPlot-1]++
			}
			// SW
			if sPlot == swPlot && swPlot == wPlot && wPlot != cPlot && sPlot > 0 {
				plotSides[sPlot-1]++
			}
		}
	}

	for idx, p := range plots {
		sides := plotSides[idx]
		price := len(p.spots) * sides
		fmt.Printf("A region of %s plants with price %d * %d = %d.\n",
			p.plant, len(p.spots), sides, price)
		result += price
	}
	return result
}

type Region12 struct {
	plant     string
	spots     map[Coord]int
	perimeter int
}

func parse12(lines []string) ([]Region12, map[Coord]int, int, int) {
	grid := toGrid(lines)
	result := make([]Region12, 0)

	visited := make(map[Coord]int)

	for x := grid.MinX; x <= grid.MaxX; x++ {
		for y := grid.MinY; y <= grid.MaxY; y++ {
			startPoint := Coord([2]int{x, y})
			_, alreadyVisited := visited[startPoint]
			if alreadyVisited {
				// fmt.Printf("Skipping %s because it is visited\n", startPoint)
				continue
			}
			currRegion := Region12{
				plant: grid.Elements[startPoint],
				spots: make(map[Coord]int),
			}
			currRegion.spots[startPoint]++
			// fmt.Printf("Starting new plot at %v for %s\n", startPoint, currRegion.plant)
			stack := []Coord{startPoint}
			for len(stack) > 0 {
				// fmt.Printf("Current stack size: %d\n", len(stack))
				// For each of the four directions, if it is different, we increase the perimete
				// If it is the same, then we add it to our current region and visited
				currSquare := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if visited[currSquare] > 0 {
					continue
				}
				visited[currSquare] = len(result) + 1

				currCrop := grid.Elements[currSquare]
				x := currSquare[0]
				y := currSquare[1]

				if currCrop != grid.Get(x-1, y) {
					currRegion.perimeter++
				} else {
					newCoord := [2]int{x - 1, y}
					if currRegion.spots[newCoord] == 0 {
						currRegion.spots[newCoord]++
						stack = append(stack, newCoord)
					}
				}
				if currCrop != grid.Get(x+1, y) {
					currRegion.perimeter++
				} else {
					newCoord := [2]int{x + 1, y}
					if currRegion.spots[newCoord] == 0 {
						currRegion.spots[newCoord]++
						stack = append(stack, newCoord)
					}
				}
				if currCrop != grid.Get(x, y-1) {
					currRegion.perimeter++
				} else {
					newCoord := [2]int{x, y - 1}
					if currRegion.spots[newCoord] == 0 {
						currRegion.spots[newCoord]++
						stack = append(stack, newCoord)
					}
				}
				if currCrop != grid.Get(x, y+1) {
					currRegion.perimeter++
				} else {
					newCoord := [2]int{x, y + 1}
					if currRegion.spots[newCoord] == 0 {
						currRegion.spots[newCoord]++
						stack = append(stack, newCoord)
					}
				}
			}
			// fmt.Printf("Finished plot: %s\n", currRegion)
			// fmt.Printf("\t%v\n", currRegion.spots)
			result = append(result, currRegion)
		}
	}
	return result, visited, grid.MaxX, grid.MaxY
}

func (r Region12) Price() int {
	return r.perimeter * len(r.spots)
}

func (r Region12) String() string {
	return fmt.Sprintf("A region of %s plants with price %d * %d = %d.",
		r.plant, len(r.spots), r.perimeter, r.Price())
}

//  Possible sides:
//  /-\  /-/  \-\  \-/
//
//

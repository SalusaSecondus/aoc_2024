package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func day13_1(lines []string) int {
	result := 0
	machines := parse13(lines)
	fmt.Printf("Machines: %+v\n", machines)
	for idx, m := range machines {
		a, b, cost := m.BestPath()
		fmt.Printf("Found path for machine %d. %d A and %d B for a cost of %d\n", idx, a, b, cost)
		if a > 0 {
			result += cost
		}
	}
	return result
}

func day13_1_beta(lines []string) int {
	result := 0
	machines := parse13(lines)
	fmt.Printf("Machines: %+v\n", machines)
	for idx, m := range machines {
		b, err := m.solveB()
		if err != nil {
			fmt.Printf("No solution for machine %d\n", idx+1)
		} else {
			remX := m.destX - b*m.bX
			remY := m.destY - b*m.bY
			a := 0
			if remX != 0 {
				a = remX / m.aX
			} else {
				a = remY / m.aY
			}
			cost := 3*a + b
			result += cost
			fmt.Printf("Found path for machine %d. %d A and %d B for a cost of %d\n", idx+1, a, b, cost)
		}
	}
	return result
}

func day13_2(lines []string) int {
	result := 0
	machines := parse13(lines)
	for idx, m := range machines {
		m.destX += 10000000000000
		m.destY += 10000000000000
		b, err := m.solveB()
		if err != nil {
			fmt.Printf("No solution for machine %d\n", idx)
		} else {
			a := (m.destX - b*m.bX) / m.aX

			cost := 3*a + b
			result += cost
			fmt.Printf("Found path for machine %d. %d A and %d B for a cost of %d\n", idx+1, a, b, cost)
		}
	}
	return result
}

type Machine13 struct {
	aX    int
	bX    int
	aY    int
	bY    int
	destX int
	destY int
}

func parse13(lines []string) []Machine13 {
	result := make([]Machine13, 0)

	buttonRe := regexp.MustCompile(`^Button (.): X\+(\d+), Y\+(\d+)`)
	prizeRe := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	currMachine := Machine13{}
	for _, l := range lines {
		matches := buttonRe.FindStringSubmatch(l)
		if len(matches) > 0 {
			if matches[1] == "A" {
				currMachine.aX, _ = strconv.Atoi(matches[2])
				currMachine.aY, _ = strconv.Atoi(matches[3])
			}
			if matches[1] == "B" {
				currMachine.bX, _ = strconv.Atoi(matches[2])
				currMachine.bY, _ = strconv.Atoi(matches[3])
			}
		} else {
			matches = prizeRe.FindStringSubmatch(l)
			if len(matches) > 0 {
				currMachine.destX, _ = strconv.Atoi(matches[1])
				currMachine.destY, _ = strconv.Atoi(matches[2])
				result = append(result, currMachine)
				currMachine = Machine13{}
			}
		}
	}
	return result
}

func (m Machine13) BestPath() (int, int, int) {
	cheapestCost := 0
	bestA := -1
	bestB := -1

	a := 0
	for {
		if bestA != -1 && 3*a > cheapestCost {
			break
		}
		baseX := a * m.aX
		baseY := a * m.aY
		// fmt.Printf("(%d, %d)\n", baseX, baseY)
		if baseX > m.destX || baseY > m.destY {
			break
		}
		b := 0
		for {
			cost := 3*a + b
			if bestA != -1 && cost > cheapestCost {
				break
			}
			finalX := baseX + b*m.bX
			finalY := baseY + b*m.bY
			// fmt.Printf("\t(%d, %d)\n", finalX, finalY)
			if finalX > m.destX || finalY > m.destY {
				break
			}
			if finalX == m.destX && finalY == m.destY {
				if bestA < 0 || cost <= cheapestCost {
					cheapestCost = cost
					bestA = a
					bestB = b
				}
			}
			b++
		}
		a++
	}
	return bestA, bestB, cheapestCost
}

func (m Machine13) solveB() (int, error) {
	top := m.destY*m.aX - m.destX*m.aY
	bottom := m.aX*m.bY - m.bX*m.aY

	if bottom == 0 {
		return 0, errors.New("division by zero")
	}
	quotiant := float64(top) / float64(bottom)
	// fmt.Printf("%d / %d = %f, %f, %v\n", top, bottom, quotiant, math.Ceil(quotiant), quotiant == math.Ceil(quotiant))
	b := int(quotiant)
	// check work
	a := (m.destX - b*m.bX) / m.aX

	finalX := a*m.aX + b*m.bX
	finalY := a*m.aY + b*m.bY
	// fmt.Printf("  a: %d, b: %d\n", a, b)
	// fmt.Printf("%d =? %d, %d =? %d\n", finalX, m.destX, finalY, m.destY)
	if finalX == m.destX && finalY == m.destY {
		return b, nil
	} else {
		return 0, errors.New("no solution")
	}
}

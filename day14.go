package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func day14_1(lines []string) int {
	robots, width, height := parse14(lines)

	time := 100
	for idx := range robots {
		currPos := robots[idx].pos
		currPos[0] += time * robots[idx].vX
		currPos[1] += time * robots[idx].vY
		currPos[0] = currPos[0] % width
		currPos[1] = currPos[1] % height
		robots[idx].pos = currPos
	}
	score, _ := score14(robots, width, height)
	return score
}

//	22
//
// 123
// 224
func day14_2(lines []string) int {
	robots, width, height := parse14(lines)

	secs := 4567
	step := 101
	for idx := range robots {
		currPos := robots[idx].pos
		currPos[0] += secs * robots[idx].vX
		currPos[1] += secs * robots[idx].vY
		currPos[0] = currPos[0] % width
		currPos[1] = currPos[1] % height
		robots[idx].pos = currPos
	}
	for {
		printRobots(robots, width, height)
		fmt.Printf("Time: %d\n\n", secs)
		// _, quads := score14(robots, width, height)
		// if quads[0] == quads[1] && quads[2] == quads[3] {
		// 	time.Sleep(1500 * time.Millisecond)
		// } else {
		time.Sleep(250 * time.Millisecond)

		// }
		for idx := range robots {

			currPos := robots[idx].pos
			currPos[0] += step * robots[idx].vX
			currPos[1] += step * robots[idx].vY
			currPos[0] = currPos[0] % width
			currPos[1] = currPos[1] % height
			robots[idx].pos = currPos
		}
		secs += step
	}
	return 0
}

func printRobots(robots []Robot14, width, height int) {
	locs := make(map[Coord]int)

	for _, r := range robots {
		locs[r.pos]++
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			val := locs[[2]int{x, y}]
			if val > 0 {
				fmt.Print(val)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func score14(robots []Robot14, width, height int) (int, []int) {
	middleCol := width / 2
	middleRow := height / 2

	quads := make([]int, 4)

	for _, r := range robots {
		q := 0
		if r.pos[0] == middleCol {
			continue
		}
		if r.pos[1] == middleRow {
			continue
		}
		if r.pos[0] > middleCol {
			q += 1
		}
		if r.pos[1] > middleRow {
			q += 2
		}
		quads[q]++
	}

	score := 1
	for _, q := range quads {
		score *= q
	}
	return score, quads
}

func parse14(lines []string) ([]Robot14, int, int) {
	result := make([]Robot14, 0)

	robotRe := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	var width int
	var height int
	if len(lines) > 20 {
		width = 101
		height = 103
	} else {
		width = 11
		height = 7
	}

	for _, l := range lines {
		matches := robotRe.FindStringSubmatch(l)
		pX, _ := strconv.Atoi(matches[1])
		pY, _ := strconv.Atoi(matches[2])
		vX, _ := strconv.Atoi(matches[3])
		if vX < 0 {
			vX += width
		}
		vY, _ := strconv.Atoi(matches[4])
		if vY < 0 {
			vY += height
		}
		pos := Coord([2]int{pX, pY})
		r := Robot14{
			pos:   pos,
			vX:    vX,
			vY:    vY,
			start: pos,
		}
		result = append(result, r)
	}

	return result, width, height
}

type Robot14 struct {
	start Coord
	vX    int
	vY    int
	pos   Coord
}

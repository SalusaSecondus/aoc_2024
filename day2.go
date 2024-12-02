package main

import (
	"cmp"
	"strconv"
	"strings"
)

func day2_1(input []string) int {
	reports := parse2(input)
	// fmt.Printf("%d, %d\n", len(reports), len(reports[0]))
	result := 0

	for i := 0; i < len(reports); i++ {
		if day2Safe(reports[i]) == -1 {
			result++
		}
	}

	return result
}

func day2_2(input []string) int {
	reports := parse2(input)

	result := 0

	for i := 0; i < len(reports); i++ {
		badIdx := day2Safe(reports[i])
		if badIdx == -1 {
			result++
		} else {
			if day2Safe(RemoveElement(reports[i], badIdx-1)) == -1 || day2Safe(RemoveElement(reports[i], badIdx)) == -1 || day2Safe(RemoveElement(reports[i], 0)) == -1 {
				result++
			}
		}
	}

	return result
}

func parse2(input []string) [][]int {
	var result [][]int
	for idx := 0; idx < len(input); idx++ {
		parts := strings.Split(input[idx], " ")
		levels := make([]int, len(parts))
		for i2 := 0; i2 < len(parts); i2++ {
			// fmt.Printf("Part %d is %s\n", i2, parts[i2])
			val, err := strconv.Atoi(parts[i2])
			check(err)
			levels[i2] = val
		}
		result = append(result, levels)
	}
	return result
}

func day2Safe(report []int) int {
	prior := 0
	direction := 0
	for x := 0; x < len(report); x++ {
		// fmt.Printf("%d ", reports[i][x])
		if prior == 0 {
			prior = report[x]
		} else {
			// fmt.Printf("%d ? %d\n", prior, reports[i][x])
			curr_dir := cmp.Compare(prior, report[x])
			curr_diff := Abs(prior - report[x])
			if curr_diff > 3 || curr_dir == 0 {
				return x
			}
			if direction == 0 {
				direction = curr_dir
			} else if direction != curr_dir {
				return x
			}
			prior = report[x]
		}
	}
	return -1
}

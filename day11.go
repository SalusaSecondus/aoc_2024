package main

import (
	"strconv"
	"strings"
)

func day11_1(lines []string) int {
	stones := parse11(lines[0])
	stones = iterate(25, stones)
	result := 0
	for _, v := range stones {
		result += v
	}
	return result
}

func day11_2(lines []string) int {
	stones := parse11(lines[0])
	stones = iterate(75, stones)
	result := 0
	for _, v := range stones {
		result += v
	}
	return result
}

func parse11(input string) map[int]int {
	result := make(map[int]int)
	parts := strings.Split(input, " ")
	for _, p := range parts {
		val, _ := strconv.Atoi(strings.TrimSpace(p))
		result[val]++
	}
	return result
}

func iterate(count int, input map[int]int) map[int]int {
	for i := 0; i < count; i++ {
		// fmt.Printf("After %d blinks\n%v\n\n", i, input)
		input = blink(input)
	}
	// fmt.Printf("After %d blinks\n%v\n\n", count, input)
	return input
}

func blink(input map[int]int) map[int]int {
	result := make(map[int]int)
	for k, v := range input {
		if k == 0 {
			result[1] += v
		} else {
			numStr := strconv.Itoa(k)
			numStrLen := len(numStr)
			if numStrLen%2 == 0 {
				leftStr := numStr[:numStrLen/2]
				rightStr := numStr[numStrLen/2:]
				leftNum, _ := strconv.Atoi(leftStr)
				rightNum, _ := strconv.Atoi(rightStr)
				result[leftNum] += v
				result[rightNum] += v
			} else {
				result[2024*k] += v
			}
		}
	}
	return result
}

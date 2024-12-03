package main

import (
	"regexp"
	"strconv"
	"strings"
)

func day3_1(input []string) int64 {
	var result int64
	joined := strings.Join(input, "")
	muls := findMuls(joined)
	for i := 0; i < len(muls); i++ {
		result += muls[i][0] * muls[i][1]
		// fmt.Println(result)
	}

	return result
}

func day3_2(input []string) int64 {
	var result int64
	joined := strings.Join(input, "")
	muls := findMuls2(joined)
	for i := 0; i < len(muls); i++ {
		result += muls[i][0] * muls[i][1]
		// fmt.Println(result)
	}

	return result
}

func findMuls(input string) [][]int64 {
	regex, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	check(err)
	matches := regex.FindAllStringSubmatch(input, -1)

	var result [][]int64
	for i := 0; i < len(matches); i++ {
		pair := matches[i]
		// fmt.Printf("Found %s with (%s, %s)\n", pair[0], pair[1], pair[2])
		a, err := strconv.Atoi(pair[1])
		check(err)
		b, err := strconv.Atoi(pair[2])
		check(err)
		ab := []int64{int64(a), int64(b)}
		result = append(result, ab)
	}
	return result
}

func findMuls2(input string) [][]int64 {
	regex, err := regexp.Compile(`don't\(\)|do\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	check(err)
	matches := regex.FindAllStringSubmatch(input, -1)

	var result [][]int64
	enabled := true
	for i := 0; i < len(matches); i++ {
		pair := matches[i]
		if pair[0] == "do()" {
			enabled = true
		} else if pair[0] == "don't()" {
			enabled = false
		} else if enabled {
			a, err := strconv.Atoi(pair[1])
			check(err)
			b, err := strconv.Atoi(pair[2])
			check(err)
			ab := []int64{int64(a), int64(b)}
			result = append(result, ab)

		}
	}
	return result
}

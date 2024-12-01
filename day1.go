package main

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

func day1_1(lines []string) int {
	result := 0

	list1, list2 := day1Parse(lines)

	for idx := 0; idx < len(list1); idx++ {
		diff := list1[idx] - list2[idx]
		if diff < 0 {
			diff = -diff
		}
		result += diff
	}
	return result
}

func day1_2(lines []string) int {
	result := 0

	list1, list2 := day1Parse(lines)

	freq := map[int]int{}

	for idx := 0; idx < len(list2); idx++ {
		freq[list2[idx]]++
	}

	for idx := 0; idx < len(list1); idx++ {
		count := freq[list1[idx]]
		result += list1[idx] * count
		// fmt.Printf("Found %d at freq %d for new value of %d\n", list1[idx], count, result)
	}
	return result
}

func day1Parse(lines []string) ([]int, []int) {
	var list1 []int
	var list2 []int

	for idx := 0; idx < len(lines); idx++ {
		line := lines[idx]
		parts := strings.SplitN(line, " ", 2)
		val1, err := strconv.Atoi(strings.Trim(parts[0], " "))
		check(err)
		val2, err := strconv.Atoi(strings.Trim(parts[1], " "))
		check(err)
		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	// Now, sort them
	slices.SortFunc(list1, cmp.Compare)
	slices.SortFunc(list2, cmp.Compare)

	return list1, list2
}

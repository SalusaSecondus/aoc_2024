package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day5_1(lines []string) int {
	result := 0
	input := parseInput(lines)
	for uIdx := range input.Updates {
		if input.IsValidUpdate(uIdx) {
			// fmt.Printf("Valid update %v\n", input.Updates[uIdx])
			middle := len(input.Updates[uIdx]) / 2
			result += input.Updates[uIdx][middle]
		}
	}
	return result
}

func day5_2(lines []string) int {
	result := 0
	input := parseInput(lines)
	for uIdx := range input.Updates {
		if !input.IsValidUpdate(uIdx) {
			for !input.IsValidUpdate(uIdx) {
				// fmt.Printf("Invalid update (%v) %v\n", input.IsValidUpdate(uIdx), input.Updates[uIdx])

				for _, rule := range input.Rules {
					first := rule[0]
					second := rule[1]

					a, aFound := input.UpdateMaps[uIdx][first]
					b, bFound := input.UpdateMaps[uIdx][second]
					if aFound && bFound && a > b {
						input.UpdateMaps[uIdx][first] = b
						input.UpdateMaps[uIdx][second] = a
						input.Updates[uIdx][b] = first
						input.Updates[uIdx][a] = second
					}

				}
			}
			// fmt.Printf("\tValid update (%v) %v\n", input.IsValidUpdate(uIdx), input.Updates[uIdx])
			middle := len(input.Updates[uIdx]) / 2
			result += input.Updates[uIdx][middle]
		}
	}
	return result
}

type Input struct {
	Rules   [][]int
	Updates [][]int

	UpdateMaps []map[int]int
}

func parseInput(input []string) Input {
	rules := make([][]int, 0)
	updates := make([][]int, 0)
	updateMaps := make([]map[int]int, 0)

	inUpdates := false
	for _, line := range input {
		if strings.TrimSpace(line) == "" {
			inUpdates = true
		} else if !inUpdates {
			parts := strings.Split(line, "|")
			a, err := strconv.Atoi(parts[0])
			check(err)
			b, err := strconv.Atoi(parts[1])
			check(err)
			rules = append(rules, []int{a, b})
		} else {
			parts := strings.Split(line, ",")
			uMap := make(map[int]int)
			pages := make([]int, 0)
			for idx, page := range parts {
				val, err := strconv.Atoi(page)
				check(err)
				pages = append(pages, val)
				foundIdx, found := uMap[val]
				if found {
					panic(fmt.Sprintf("Update %d had duplicate page %d at locations %d and %d", len(updates), val, foundIdx, idx))
				}
				uMap[val] = idx
			}
			updates = append(updates, pages)
			updateMaps = append(updateMaps, uMap)
		}
	}

	return Input{
		Rules:      rules,
		Updates:    updates,
		UpdateMaps: updateMaps,
	}
}

func (i Input) IsValidUpdate(idx int) bool {
	for _, rule := range i.Rules {
		first := rule[0]
		second := rule[1]

		a, aFound := i.UpdateMaps[idx][first]
		b, bFound := i.UpdateMaps[idx][second]

		if aFound && bFound && a > b {
			return false
		}
	}
	return true
}

package main

import "strings"

func day19_1(lines []string) int {
	towels, patterns := parse19(lines)

	result := 0
	cache := make(map[string]bool)
	for _, p := range patterns {
		if isPossible19(towels, p, &cache) {
			result++
		}
	}
	return result
}

func day19_2(lines []string) int {
	result := 0
	return result
}

func isPossible19(towels []string, pattern string, cache *map[string]bool) bool {
	if len(pattern) == 0 {
		return true
	}
	cAns, found := (*cache)[pattern]
	if found {
		return cAns
	}

	for _, t := range towels {
		if strings.HasPrefix(pattern, t) {
			suffix := pattern[len(t):]
			if isPossible19(towels, suffix, cache) {
				(*cache)[pattern] = true
				return true
			}
		}
	}
	(*cache)[pattern] = false
	return false
}

func parse19(lines []string) ([]string, []string) {
	towels := strings.Split(lines[0], ", ")
	patterns := make([]string, 0)
	for i := 2; i < len(lines); i++ {
		patterns = append(patterns, lines[i])
	}
	return towels, patterns
}

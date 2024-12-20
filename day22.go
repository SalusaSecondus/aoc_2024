package main

import (
	"fmt"
	"strconv"
)

func day22_1(lines []string) int {
	result := 0
	for _, seedStr := range lines {
		seed, err := strconv.Atoi(seedStr)
		check(err)
		secret := seed
		for idx := 1; idx <= 2000; idx++ {
			secret = Step(secret)
		}
		fmt.Printf("%d @ 2000 = %d\n", seed, secret)
		result += secret
	}
	return result
}

func day22_2(lines []string) int {
	allPrices := make([]BananaMap, 0)

	for _, seedStr := range lines {
		seed, err := strconv.Atoi(seedStr)
		check(err)
		secret := seed
		prices := make([]int, 1)
		prices[0] = seed % 10
		diffs := make([]int, 1)
		lastPrice := prices[0]
		for idx := 1; idx <= 2000; idx++ {
			secret = Step(secret)
			nextPrice := secret % 10
			nextDiff := nextPrice - lastPrice
			diffs = append(diffs, nextDiff)
			prices = append(prices, nextPrice)
			lastPrice = nextPrice
		}
		currMonkey := BananaMap(make(map[[4]int]int))
		for idx := 1; idx < 1997; idx++ {
			currMonkey.Put(diffs[idx:idx+4], prices[idx+3])
		}
		allPrices = append(allPrices, currMonkey)
	}

	allSigs := make(map[[4]int]int)
	for _, monkey := range allPrices {
		for sig := range monkey {
			allSigs[sig]++
		}
	}

	bestPrice := -1
	var bestSig [4]int
	for sig := range allSigs {
		currPrice := 0
		for _, monkey := range allPrices {
			currPrice += monkey[sig]
		}
		if currPrice > bestPrice {
			bestPrice = currPrice
			bestSig = sig
			fmt.Printf("New best %+v -> %d\n", bestSig, bestPrice)
		}
	}
	return bestPrice
}

func Mix(s, n int) int {
	return s ^ n
}

func Prune(s int) int {
	return s % 16777216
}

func MixPrune(s, n int) int {
	return Prune(Mix(s, n))
}

func Step(s int) int {
	s = MixPrune(s, s*64)
	// fmt.Printf("\ta: %d\n", s)
	s = MixPrune(s, s/32)
	// fmt.Printf("\tb: %d\n", s)
	s = MixPrune(s, s*2048)
	// fmt.Printf("\tc: %d\n", s)
	return s
}

type BananaMap map[[4]int]int

func (m *BananaMap) Put(s []int, price int) {
	if len(s) != 4 {
		panic("Wrong slice length")
	}
	var diffs [4]int
	copy(diffs[:], s[:4])
	if _, found := (*m)[diffs]; !found {
		(*m)[diffs] = price
	}
}

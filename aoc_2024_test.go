package main

import (
	"testing"
)

func TestDay1(t *testing.T) {
	smoke_input := readInput(1, true)
	smoke_1 := day1_1(smoke_input)
	if smoke_1 != 11 {
		t.Fatalf("Day 1_1. Expected 11 but got %d", smoke_1)
	}

	real_input := readInput(1, false)
	real_1 := day1_1(real_input)
	if real_1 != 1970720 {
		t.Fatalf("Day 1_1. Expected 1970720 but got %d", real_1)
	}

	smoke_2 := day1_2(smoke_input)
	if smoke_2 != 31 {
		t.Fatalf("Day 1_1. Expected 31 but got %d", smoke_2)
	}

	real_2 := day1_2(real_input)
	if real_2 != 17191599 {
		t.Fatalf("Day 1_1. Expected 17191599 but got %d", real_2)
	}
}

func TestDay2(t *testing.T) {
	assertIntDay(2, 1, true, 2, day2_1, t)
	assertIntDay(2, 2, true, 4, day2_2, t)

	assertIntDay(2, 1, false, 299, day2_1, t)
	assertIntDay(2, 2, false, 364, day2_2, t)
}

func TestDay3(t *testing.T) {
	assertInt64Day(3, 1, true, 161, day3_1, t)
	assertInt64Day(3, 2, true, 48, day3_2, t)

	assertInt64Day(3, 1, false, 155955228, day3_1, t)
	assertInt64Day(3, 2, false, 100189366, day3_2, t)
}

func TestDay4(t *testing.T) {
	assertIntDay(4, 1, true, 18, day4_1, t)
	assertIntDay(4, 2, true, 9, day4_2, t)

	assertIntDay(4, 1, false, 2414, day4_1, t)
	assertIntDay(4, 2, false, 1871, day4_2, t)
}

func TestDay5(t *testing.T) {
	assertIntDay(5, 1, true, 143, day5_1, t)
	assertIntDay(5, 2, true, 123, day5_2, t)

	assertIntDay(5, 1, false, 5964, day5_1, t)
	assertIntDay(5, 2, false, 4719, day5_2, t)
}

func TestDay6(t *testing.T) {
	assertIntDay(6, 1, true, 41, day6_1, t)
	assertIntDay(6, 2, true, 6, day6_2, t)

	assertIntDay(6, 1, false, 4647, day6_1, t)
	assertIntDay(6, 2, false, 1723, day6_2, t)
}

func BenchmarkDay62(b *testing.B) {
	input := readInput(6, false)
	for i := 0; i < b.N; i++ {
		day6_2(input)
	}
}

func TestDay7(t *testing.T) {
	day := 7
	assertIntDay(day, 1, true, 3749, day7_1, t)
	assertIntDay(day, 1, false, 5702958180383, day7_1, t)

	assertIntDay(day, 2, true, 11387, day7_2, t)
	assertIntDay(day, 2, false, 92612386119138, day7_2, t)
}

func BenchmarkDay7(b *testing.B) {
	lines := readInput(7, false)
	for i := 0; i < b.N; i++ {
		day7_2(lines)
	}
}

func TestDay8(t *testing.T) {
	day := 8
	assertIntDay(day, 1, true, 14, day8_1, t)
	assertIntDay(day, 1, false, 390, day8_1, t)

	assertIntDay(day, 2, true, 34, day8_2, t)
	assertIntDay(day, 2, false, 1246, day8_2, t)
}

func TestDay9(t *testing.T) {
	day := 9
	assertIntDay(day, 1, true, 1928, day9_1, t)
	assertIntDay(day, 1, false, 6463499258318, day9_1, t)

	assertIntDay(day, 2, true, 2858, day9_2, t)
	assertIntDay(day, 2, false, 6493634986625, day9_2, t)
}

func BenchmarkDay9_1(t *testing.B) {
	lines := readInput(9, false)
	for i := 0; i < t.N; i++ {
		day9_1(lines)
	}
}

func BenchmarkDay9_2(t *testing.B) {
	lines := readInput(9, false)
	for i := 0; i < t.N; i++ {
		day9_2(lines)
	}
}

func TestDay10(t *testing.T) {
	day := 10
	assertIntDay(day, 1, true, 36, day10_1, t)
	assertIntDay(day, 1, false, 472, day10_1, t)

	assertIntDay(day, 2, true, 81, day10_2, t)
	assertIntDay(day, 2, false, 969, day10_2, t)
}

func BenchmarkDay10_1(t *testing.B) {
	lines := readInput(10, false)
	for i := 0; i < t.N; i++ {
		day10_1(lines)
	}
}

func BenchmarkDay10_2(t *testing.B) {
	lines := readInput(10, false)
	for i := 0; i < t.N; i++ {
		day10_2(lines)
	}
}

func TestDay11(t *testing.T) {
	day := 11
	assertIntDay(day, 1, true, 55312, day11_1, t)
	assertIntDay(day, 1, false, 197157, day11_1, t)

	// assertIntDay(day, 2, true, 0, day11_2, t)
	assertIntDay(day, 2, false, 234430066982597, day11_2, t)
}

func BenchmarkDay11_2(t *testing.B) {
	lines := readInput(11, false)
	for i := 0; i < t.N; i++ {
		day11_2(lines)
	}
}

func TestDay12(t *testing.T) {
	day := 12
	assertIntDay(day, 1, true, 1930, day12_1, t)
	assertIntDay(day, 1, false, 1415378, day12_1, t)

	assertIntDay(day, 2, true, 1206, day12_2, t)
	assertIntDay(day, 2, false, 862714, day12_2, t)
}

func TestDay13(t *testing.T) {
	day := 13
	// assertIntDay(day, 1, true, 480, day13_1, t)
	// assertIntDay(day, 1, false, 29711, day13_1, t)
	assertIntDay(day, 1, true, 480, day13_1_beta, t)
	assertIntDay(day, 1, false, 29711, day13_1_beta, t)
	t.Logf("Part2\n")
	assertIntDay(day, 2, true, 875318608908, day13_2, t)
	assertIntDay(day, 2, false, 94955433618919, day13_2, t)
}

func BenchmarkDay13(t *testing.B) {
	lines := readInput(13, false)
	for i := 0; i < t.N; i++ {
		day13_2(lines)
	}
}

func TestDay14(t *testing.T) {
	day := 14
	assertIntDay(day, 1, true, 12, day14_1, t)
	assertIntDay(day, 1, false, 211692000, day14_1, t)

	// assertIntDay(day, 2, true, 0, day14_2, t)
	// assertIntDay(day, 2, false, 0, day14_2, t)
}

type dayIntFunc func([]string) int

func assertIntDay(day, part int, smoke bool, expected int, fn dayIntFunc, t *testing.T) {
	input := readInput(day, smoke)
	result := fn(input)
	prefix := "Real"
	if smoke {
		prefix = "Smoke"
	}
	if result != expected {
		t.Fatalf("%s %d_%d failed. Expected %d but got %d", prefix, day, part, expected, result)
	}
}

type dayInt64Func func([]string) int64

func assertInt64Day(day, part int, smoke bool, expected int64, fn dayInt64Func, t *testing.T) {
	input := readInput(day, smoke)
	result := fn(input)
	prefix := "Real"
	if smoke {
		prefix = "Smoke"
	}
	if result != expected {
		t.Fatalf("%s %d_%d failed. Expected %d but got %d", prefix, day, part, expected, result)
	}
}

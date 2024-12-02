package main

import "testing"

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
	smoke_input := readInput(2, true)
	smoke_1 := day2_1(smoke_input)
	if smoke_1 != 2 {
		t.Fatalf("Smoke 2_1. Expected 2 but got %d", smoke_1)
	}
	smoke_2 := day2_2(smoke_input)
	if smoke_2 != 4 {
		t.Fatalf("Smoke 2_1. Expected 4 but got %d", smoke_2)
	}

	real := readInput(2, false)
	real_1 := day2_1(real)
	if real_1 != 299 {
		t.Fatalf("Day 2_1. Expected 299 but got %d", real_1)
	}
	real_2 := day2_2(real)
	if real_2 != 364 {
		t.Fatalf("Day 2_2. Expected 364 but got %d", real_2)
	}
}

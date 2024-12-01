package main

import "testing"

func TestDay1_1(t *testing.T) {
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
	if real_2 != 1970720 {
		t.Fatalf("Day 1_1. Expected 1970720 but got %d", real_2)
	}
}

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day7_1(lines []string) int {
	result := 0
	equations := parse7(lines)
	legalOps := []Operator7{Add7, Mult7}
	for _, eq := range equations {
		ops := eq.Solve(legalOps)
		if ops != nil {
			// fmt.Println(eq.Formatted(ops))
			result += eq.Solution
		}

	}
	return result
}

func day7_2(lines []string) int {
	result := 0
	equations := parse7(lines)
	legalOps := []Operator7{Add7, Mult7, Concat7}
	for _, eq := range equations {
		ops := eq.Solve(legalOps)
		if ops != nil {
			// fmt.Println(eq.Formatted(ops))
			result += eq.Solution
		}

	}
	return result
}

type Equation7 struct {
	Solution int
	Operands []int
}

func (e Equation7) Solve(legalOps []Operator7) []Operator7 {
	ops := make([]Operator7, len(e.Operands)-1)
	if e.solveInner(legalOps, 0, ops, e.Operands[0]) {
		return ops
	} else {
		return nil
	}
}

func (e Equation7) Formatted(ops []Operator7) string {
	result := fmt.Sprintf("%d =? %d", e.Solution, e.Operands[0])
	for idx, op := range ops {
		result = fmt.Sprintf("%s %c %d", result, op, e.Operands[idx+1])
	}
	return result
}

func (e Equation7) solveInner(legalOps []Operator7, offset int, ops []Operator7, tally int) bool {
	if offset == len(ops) {
		return tally == e.Solution
	}
	for _, nextOp := range legalOps {
		ops[offset] = Operator7(nextOp)

		nextTally := tally
		next := e.Operands[offset+1]
		switch nextOp {
		case Mult7:
			nextTally *= next
		case Add7:
			nextTally += next
		case Concat7:
			digits := len(strconv.Itoa(next))
			nextTally *= int(math.Pow10(digits))
			nextTally += next
		}
		if nextTally <= e.Solution && e.solveInner(legalOps, offset+1, ops, nextTally) {
			return true
		}
	}
	return false
}

type Operator7 byte

const (
	Mult7   Operator7 = '*'
	Add7    Operator7 = '+'
	Concat7 Operator7 = '|'
)

func parse7(lines []string) []Equation7 {
	result := make([]Equation7, 0)
	for _, l := range lines {
		parts1 := strings.Split(l, ":")
		solution, err := strconv.Atoi(parts1[0])
		check(err)
		parts2 := strings.Split(strings.TrimSpace(parts1[1]), " ")
		operands := make([]int, 0)
		for _, p := range parts2 {
			op, err := strconv.Atoi(strings.TrimSpace(p))
			check(err)
			operands = append(operands, op)
		}
		result = append(result, Equation7{Solution: solution, Operands: operands})
	}
	return result
}

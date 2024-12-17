package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day17_1(lines []string) string {
	computer := parse17(lines)
	computer.Start()

	return computer.OutputString()
}

func day17_2(lines []string) int {
	computer := parse17(lines)

	starts := make([]int, 1)
	for idx := 0; idx < len(computer.Memory); idx++ {
		// fmt.Printf("Idx %d: Starts %d\n", idx, len(starts))
		nextStarts := make(map[int]bool)
		for _, s := range starts {
			for nextOp := 0; nextOp < 256; nextOp++ {
				a := nextOp<<(3*idx) | s
				if computer.IsQuine(a) {
					// fmt.Printf("\t%018b:\t%s\n", a, computer.OutputString())

					return a
				}
				if len(computer.Output) > idx+1 {
					// fmt.Printf("\t%018b:\t%s\n", a, computer.OutputString())
					a = ((nextOp & 0x7) << (3 * idx)) | s
					nextStarts[a] = true
				}
			}
		}
		starts = make([]int, 0)
		for k := range nextStarts {
			starts = append(starts, k)
		}

		sort.Ints(starts)
	}
	panic("Not found")
}

func day17_2_old(lines []string) int {
	computer := parse17(lines)

	a := 1 << (len(computer.Memory) * 3)
	for !computer.IsQuine(a) {
		if len(computer.Output) > 7 {
			fmt.Printf("%d: %d/%d\t%d\t%d\t%d\t%d\t%d\n", a, len(computer.Output), len(computer.Memory),
				a&0x3f,
				a&0x1f,
				a&0x67,
				a&0x87,
				a&0x1c7)
		}
		if computer.Memory[0] == 0 {
			a++
		} else {
			for {
				a++
				// if a&0x3f == 42 ||
				// 	a&0x1f == 19 ||
				// 	a&0x67 == 100 ||
				// 	a&0x87 == 134 ||
				// 	a&0x1c7 == 7 {
				// 	break
				// }
				if a&0x3f == 42 ||
					a&0x1f == 10 ||
					a&0x67 == 34 ||
					a&0x87 == 2 ||
					a&0x1c7 == 2 {
					break
				}
				if a&0x3f == 47 ||
					a&0x1f == 15 ||
					a&0x67 == 39 ||
					a&0x87 == 7 ||
					a&0x1c7 == 7 {
					break
				}
				// if a%1000000 == 0 {
				// 	fmt.Printf("%d\n", a)
				// }
			}
		}
	}
	fmt.Printf("Found %d -> %s\n", a, computer.OutputString())
	return a
}

func parse17(lines []string) Computer {
	aStr := lines[0][12:]
	bStr := lines[1][12:]
	cStr := lines[2][12:]

	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)
	c, _ := strconv.Atoi(cStr)

	memStr := lines[4][9:]
	mem := make([]int, 0)
	for _, val := range strings.Split(memStr, ",") {
		num, _ := strconv.Atoi(val)
		mem = append(mem, num)
	}

	return Computer{
		RegA:   a,
		RegB:   b,
		RegC:   c,
		IP:     0,
		Memory: mem,
	}
}

type Computer struct {
	RegA   int
	RegB   int
	RegC   int
	IP     int
	Memory []int
	Output []int
}

type Instruction byte

const (
	adv Instruction = 0
	bxl Instruction = 1
	bst Instruction = 2
	jnz Instruction = 3
	bxc Instruction = 4
	out Instruction = 5
	bdv Instruction = 6
	cdv Instruction = 7
)

func (c *Computer) Reset() {
	c.IP = 0
	c.RegA = 0
	c.RegB = 0
	c.RegC = 0
	c.Output = make([]int, 0)
}

func (c *Computer) operand(op int) int {
	if op < 4 {
		return op
	} else if op == 4 {
		return c.RegA
	} else if op == 5 {
		return c.RegB
	} else if op == 6 {
		return c.RegC
	} else {
		panic(fmt.Sprintf("Invalid operand %d", op))
	}
}

func (c Computer) OutputString() string {
	result := ""
	for _, val := range c.Output {
		if len(result) == 0 {
			result = strconv.Itoa(val)
		} else {
			result = fmt.Sprintf("%s,%d", result, val)
		}
	}
	return result
}

func (c *Computer) IsQuine(startA int) bool {
	c.Reset()
	c.RegA = startA
	for c.IP >= 0 && c.IP < len(c.Memory) {
		if c.Execute(Instruction(c.Memory[c.IP]), c.Memory[c.IP+1]) {
			c.IP += 2
		} else {
			// Also happens to mean we've got a new output
			if len(c.Output) > len(c.Memory) {
				return false
			}
			idx := len(c.Output) - 1
			if c.Output[idx] != c.Memory[idx] {
				return false
			}
		}
	}
	idx := len(c.Output) - 1
	return len(c.Output) == len(c.Memory) && c.Output[idx] == c.Memory[idx]
}

func (c *Computer) Execute(inst Instruction, op int) bool {
	result := true
	switch inst {
	case adv:
		c.RegA /= 1 << c.operand(op)
	case bxl:
		c.RegB ^= op
	case bst:
		c.RegB = c.operand(op) & 0x7
	case jnz:
		if c.RegA != 0 {
			c.IP = op
			result = false
		}
	case bxc:
		c.RegB ^= c.RegC
	case out:
		c.Output = append(c.Output, c.operand(op)&0x7)
	case bdv:
		c.RegB = c.RegA / (1 << c.operand(op))
	case cdv:
		c.RegC = c.RegA / (1 << c.operand(op))
	default:
		panic(fmt.Sprintf("Invalid instruction: %d", inst))
	}
	return result
}

func (c *Computer) Start() {
	for c.IP >= 0 && c.IP < len(c.Memory) {
		if c.Execute(Instruction(c.Memory[c.IP]), c.Memory[c.IP+1]) {
			c.IP += 2
		}
	}
}

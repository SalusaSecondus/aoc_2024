package main

import (
	"fmt"
	"strconv"
)

func day21_1(lines []string) int {
	result := 0

	// example := "029A"
	// numPad := NumPad21{}
	// d1 := buildDirPad(&numPad)
	// d2 := buildDirPad(&d1)
	// d3 := buildDirPad(&d2)
	d3 := robotChain(2)
	for _, l := range lines {
		d3.Reset()

		for idx := range l {
			d3.Press(l[idx])
			// fmt.Printf("\t%s\n", d3.trace)
		}
		first := l[:3]
		firstVal, _ := strconv.Atoi(first)
		fmt.Printf("%d * %d: %s\n", len(d3.Trace()[0]), firstVal, d3.Trace()[0])
		result += firstVal * len(d3.Trace()[0])
	}
	return result
}

func robotChain(len int) Keypad21 {
	var result Keypad21
	result = &NumPad21{}
	for i := 0; i < len; i++ {
		tmp := buildDirPad(result)
		result = &tmp
	}
	tmp := buildDirPad(result)
	return &tmp
}

func day21_2(lines []string) int {
	result := 0

	// example := "029A"
	// numPad := NumPad21{}
	// d1 := buildDirPad(&numPad)
	// d2 := buildDirPad(&d1)
	// d3 := buildDirPad(&d2)
	d3 := robotChain(5)
	for _, l := range lines {
		d3.Reset()

		for idx := range l {
			d3.Press(l[idx])
			// fmt.Printf("\t%s\n", d3.trace)
		}
		first := l[:3]
		firstVal, _ := strconv.Atoi(first)
		fmt.Printf("%d * %d: %s\n", len(d3.Trace()[0]), firstVal, d3.Trace()[0])
		for _, t := range d3.Trace() {
			fmt.Println(t)
		}
		result += firstVal * len(d3.Trace()[0])
	}
	return result
}

func buildDirPad[KP Keypad21](k KP) DirPad21[KP] {
	return DirPad21[KP]{
		keypad: k,
	}
}

type Keypad21 interface {
	Press(byte) string
	Reset()
	Trace() []string
	Find(byte) Coord
	Execute(byte) string
	Loc() Coord
	Val() byte
}

type NumPad21 struct {
	loc   Coord
	trace string
}

func (n NumPad21) Loc() Coord {
	return n.loc
}

func (n NumPad21) Val() byte {
	for guess := byte('0'); guess <= '9'; guess++ {
		if n.loc == n.Find(guess) {
			return guess
		}
	}
	if n.loc == n.Find('A') {
		return 'A'
	}
	panic("Unknown location")
}

func (n *NumPad21) Execute(val byte) string {
	fmt.Printf("\tEXECUTE %c\n", val)
	if val == 'A' {
		return fmt.Sprintf("%c", n.Val())
	}
	n.loc = Dir15(val).next(n.loc)
	return ""
}

func (n NumPad21) Find(val byte) Coord {
	switch val {
	case 'A':
		return [2]int{2, 3}
	case '0':
		return [2]int{1, 3}
	case '1':
		return [2]int{0, 2}
	case '2':
		return [2]int{1, 2}
	case '3':
		return [2]int{2, 2}
	case '4':
		return [2]int{0, 1}
	case '5':
		return [2]int{1, 1}
	case '6':
		return [2]int{2, 1}
	case '7':
		return [2]int{0, 0}
	case '8':
		return [2]int{1, 0}
	case '9':
		return [2]int{2, 0}
	default:
		panic(fmt.Sprintf("Unexpected val: %c", val))
	}
}

func (n *NumPad21) Press(val byte) string {
	dest := n.Find(val)
	path := ""

	if n.loc[1] == 3 && dest[0] == 0 {
		for n.loc[1] > dest[1] {
			path += "^"
			n.loc[1]--
		}
	}

	if n.loc[0] == 0 && dest[1] == 3 {
		for n.loc[0] < dest[0] {
			path += ">"
			n.loc[0]++
		}
	}

	for n.loc[1] < dest[1] {
		path += "v"
		n.loc[1]++
	}
	for n.loc[0] > dest[0] {
		path += "<"
		n.loc[0]--
	}

	for n.loc[0] < dest[0] {
		path += ">"
		n.loc[0]++
	}
	for n.loc[1] > dest[1] {
		path += "^"
		n.loc[1]--
	}
	n.trace += fmt.Sprintf("%c", val)
	return path + "A"
}

func (n *NumPad21) Reset() {
	n.trace = ""
	n.loc = n.Find('A')
}

func (n NumPad21) Trace() []string {
	return []string{n.trace}
}

type DirPad21[K Keypad21] struct {
	loc    Coord
	trace  string
	keypad K
}

func (d DirPad21[K]) Find(val byte) Coord {
	switch val {
	case '^':
		return [2]int{1, 0}
	case 'A':
		return [2]int{2, 0}
	case '<':
		return [2]int{0, 1}
	case 'v':
		return [2]int{1, 1}
	case '>':
		return [2]int{2, 1}
	default:
		panic(fmt.Sprintf("Unknown value %c", val))
	}
}

func (d *DirPad21[K]) Press(val byte) string {
	subPath := d.keypad.Press(val)

	path := ""
	for idx := range subPath {
		path += d.SubPress(subPath[idx])
		path += "A"
	}
	// path += d.SubPress('A')

	return path
}

func (d *DirPad21[K]) SubPress(val byte) string {
	dest := d.Find(val)
	path := ""

	for d.loc[1] < dest[1] {
		path += "v"
		d.loc[1]++
	}
	for d.loc[0] < dest[0] {
		path += ">"
		d.loc[0]++
	}

	for d.loc[1] > dest[1] {
		path += "^"
		d.loc[1]--
	}
	for d.loc[0] > dest[0] {
		path += "<"
		d.loc[0]--
	}
	d.trace += fmt.Sprintf("%c", val)

	return path
}

func (d DirPad21[K]) Trace() []string {
	prefix := []string{d.trace}
	return append(prefix, d.keypad.Trace()...)
}

func (d *DirPad21[K]) Reset() {
	d.keypad.Reset()
	d.loc = d.Find('A')
	d.trace = ""
}

func (d DirPad21[K]) Loc() Coord {
	return d.loc
}

func (d DirPad21[K]) Val() byte {
	if d.loc == d.Find(byte(Up15)) {
		return byte(Up15)
	}
	if d.loc == d.Find(byte(Down15)) {
		return byte(Down15)
	}
	if d.loc == d.Find(byte(Left15)) {
		return byte(Left15)
	}
	if d.loc == d.Find(byte(Right15)) {
		return byte(Right15)
	}
	if d.loc == d.Find('A') {
		return 'A'
	}
	panic(fmt.Sprintf("Unknown location: %s", d.loc))
}

func (d *DirPad21[K]) Execute(val byte) string {
	fmt.Printf("\tExecute %c\n", val)
	if val == 'A' {
		return d.keypad.Execute(d.Val())
	}

	d.loc = Dir15(val).next(d.loc)
	return ""
}

func tokenMap(token string) []string {
	switch token {
	case "A":
		return []string{"A"}
	case "vA":
		return []string{"v<A", "^>A"}
	case "<A":
		return []string{"v<<A", ">>^A"}
	case "^A":
		return []string{"<A", ">A"}
	case ">A":
		return []string{"vA", "^A"}
	case "^>A":
		return []string{"<A", "v>A", "^A"}
	case "v<A":
		return []string{"v<A", "<A", ">>^A"}
	case "<vA":
		return []string{"v<<A", ">A", ">^A"}
	case "<v<A":
		return []string{"v<<A", ">A", "<A", ">>^A"}
	case "v<<A":
		return []string{"v<A", "<A", "A", ">>^A"}
	default:
		panic("Unsupported token: " + token)

	}
}

package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day24_1(lines []string) int {
	circuit := parse24(lines)
	circuit.Evaluate()
	return circuit.Score()
}

func day24_2(lines []string) int {
	circuit := parse24(lines)

	fmt.Println(circuit.Graphviz())
	result := 0
	return result
}

type GateOp string

const (
	XOR GateOp = "XOR"
	OR  GateOp = "OR"
	AND GateOp = "AND"
)

func parse24(lines []string) Circuit {
	Wires := make(map[string]uint8)
	Gates := make([]Gate, 0)
	Connections := make(map[string][]Gate)

	inputRe := regexp.MustCompile(`^(\S+): (\d+)\s*$`)
	gateRe := regexp.MustCompile(`^(\S+) (\S+) (\S+) -> (\S+)`)

	for _, l := range lines {
		inputParts := inputRe.FindStringSubmatch(l)
		if inputParts != nil {
			value, _ := strconv.Atoi(inputParts[2])
			Wires[inputParts[1]] = uint8(value)
			continue
		}
		gateParts := gateRe.FindStringSubmatch(l)
		if gateParts != nil {
			wire1 := gateParts[1]
			wire2 := gateParts[3]
			output := gateParts[4]
			gate := Gate{
				Wire1:  wire1,
				Op:     GateOp(gateParts[2]),
				Wire2:  wire2,
				Output: output,
			}
			Gates = append(Gates, gate)

			Connections[wire1] = append(Connections[wire1], gate)
			Connections[wire2] = append(Connections[wire2], gate)
		}
	}
	return Circuit{
		Wires:       Wires,
		Gates:       Gates,
		Connections: Connections,
	}
}

type Circuit struct {
	Wires       map[string]uint8
	Gates       []Gate
	Connections map[string][]Gate
}

func (c *Circuit) Evaluate() {
	todo := make([]string, 0)
	for wire := range c.Wires {
		todo = append(todo, wire)
	}

	for len(todo) > 0 {
		curr := todo[len(todo)-1]
		todo = todo[:len(todo)-1]

		for _, g := range c.Connections[curr] {
			wire1, w1Found := c.Wires[g.Wire1]
			wire2, w2Found := c.Wires[g.Wire2]

			if w1Found && w2Found {
				value := g.Op.Eval(wire1, wire2)
				c.Wires[g.Output] = value
				todo = append(todo, g.Output)
			}
		}
	}
}

func (c Circuit) Score() int {
	result := 0
	for wire, value := range c.Wires {
		if wire[0] == 'z' {
			shift, err := strconv.Atoi(wire[1:])
			check(err)
			result |= int(value) << shift
		}
	}
	return result
}

func (c Circuit) Graphviz() string {
	result := "digraph G {\n"

	for w := range c.Wires {
		result += fmt.Sprintf("\t%s[shape=box, fillcolor=blue];\n", w)
	}
	result += "\n"
	for _, g := range c.Gates {
		var shape string
		switch g.Op {
		case AND:
			shape = "invtriangle"
		case OR:
			shape = "circle"
		case XOR:
			shape = "doublecircle"
		}
		result += fmt.Sprintf("\t%s[shape=%s];\n", g.Output, shape)
	}

	result += "\n"
	for _, g := range c.Gates {
		result += fmt.Sprintf("\t%s -> %s;\n", g.Wire1, g.Output)
		result += fmt.Sprintf("\t%s -> %s;\n", g.Wire2, g.Output)

	}
	result += "\n}"
	return result
}

type Gate struct {
	Op     GateOp
	Wire1  string
	Wire2  string
	Output string
}

func (o GateOp) Eval(wire1, wire2 uint8) uint8 {
	switch o {
	case XOR:
		return wire1 ^ wire2
	case AND:
		return wire1 & wire2
	case OR:
		return wire1 | wire2
	default:
		panic(fmt.Sprintf("Unknown GateOp: %s", o))
	}
}

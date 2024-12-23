package main

import (
	"slices"
	"strings"
)

func day23_1(lines []string) int {
	result := 0
	graph := parse23(lines)
	cliques := findCliques(graph)
	// fmt.Printf("%+v\n", cliques)
	for c := range cliques {
		if c[0][0] == 't' || c[1][0] == 't' || c[2][0] == 't' {
			result++
		}
	}
	return result
}

func day23_2(lines []string) string {
	graph := parse23(lines)
	ch := NewCliqueHolder()
	p := make(map[string]bool)
	for node := range graph {
		p[node] = true
	}
	bk2(make([]string, 0), p, make(map[string]bool), graph, &ch)
	slices.Sort(*ch.clique)
	// fmt.Printf("Largest clique: %+v\n", ch.clique)
	return strings.Join(*ch.clique, ",")
}

func findCliques(graph Graph23) map[[3]string]bool {
	result := make(map[[3]string]bool)

	for node1, links := range graph {
		for node2 := range links {
			for node3 := range graph[node2] {
				if graph[node3][node1] {
					clique := [3]string{node1, node2, node3}
					slices.Sort(clique[:])
					result[clique] = true
				}
			}
		}
	}
	return result
}

func bk2(r []string, p, x map[string]bool, graph Graph23, ch *CliqueHolder) {
	if len(p) == 0 && len(x) == 0 {
		ch.addClique(r)
		return
	}
	uList := make([]string, 0)
	for u := range p {
		uList = append(uList, u)
	}
	for u := range x {
		uList = append(uList, u)
	}

	for _, u := range uList {
		for v := range p {
			if graph[u][v] {
				// Skip neighbors
				continue
			}
			rV := slices.Clone(r)
			rV = append(rV, v)

			pN := make(map[string]bool)
			xN := make(map[string]bool)

			for neighbor := range graph[v] {
				if p[neighbor] {
					pN[neighbor] = true
				}
				if x[neighbor] {
					xN[neighbor] = true
				}
			}
			bk2(rV, pN, xN, graph, ch)
			delete(p, v)
			x[v] = true
		}
	}

}

type CliqueHolder struct {
	clique *[]string
}

func NewCliqueHolder() CliqueHolder {
	clique := make([]string, 0)
	return CliqueHolder{clique: &clique}
}

func (cl *CliqueHolder) addClique(clique []string) {
	if len(clique) > len(*cl.clique) {
		biggerClique := slices.Clone(clique)
		cl.clique = &biggerClique
	}
}

type Graph23 map[string]map[string]bool

func (g *Graph23) addLink(node1, node2 string) {
	if _, found := (*g)[node1]; !found {
		(*g)[node1] = make(map[string]bool)
	}
	if _, found := (*g)[node2]; !found {
		(*g)[node2] = make(map[string]bool)
	}
	(*g)[node1][node2] = true
	(*g)[node2][node1] = true
}

func parse23(lines []string) Graph23 {
	result := make(Graph23)

	for _, l := range lines {
		parts := strings.SplitN(l, "-", 2)

		result.addLink(parts[0], parts[1])
	}
	return result
}

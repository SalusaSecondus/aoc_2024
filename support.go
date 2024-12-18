package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"
)

func Abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func fileToLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var lines []string

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	file.Close()

	return lines, nil
}

func readInput(day int, smoke bool) []string {
	var filename string
	if smoke {
		filename = fmt.Sprintf("input/day%d_smoke.txt", day)
	} else {
		filename = fmt.Sprintf("input/day%d.txt", day)

	}
	lines, err := fileToLines(filename)

	if err != nil && !smoke && strings.Contains(err.Error(), "no such file") {
		loadDay(day)
		lines, err = fileToLines(filename)
	}
	check(err)

	return lines
}

func loadDay(day int) {
	fmt.Printf("Retrieving day %d\n", day)
	creds, err := fileToLines(".creds")
	check(err)
	cookieVal := creds[0]

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	req.Header.Add("Cookie", cookieVal)
	resp, err := http.DefaultClient.Do(req)
	check(err)
	fmt.Printf("Code: %d\nResponse: %s", resp.StatusCode, resp.Body)
	defer resp.Body.Close()
	value, err := io.ReadAll(resp.Body)
	check(err)
	filename := fmt.Sprintf("input/day%d.txt", day)
	os.WriteFile(filename, value, 0644)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RemoveElement[S ~[]E, E any](s S, i int) S {
	result := make([]E, len(s))
	copy(result, s)
	return slices.Delete(result, i, i+1)
}

type Coord [2]int

func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c[0], c[1])
}

type Grid[V comparable] struct {
	Elements               map[Coord]V
	MinX, MaxX, MinY, MaxY int
	Default                V
}

func NewGrid[V comparable]() Grid[V] {
	elements := make(map[Coord]V)
	return Grid[V]{
		Elements: elements,
	}
}

func (g Grid[V]) Get(x, y int) V {
	coord := [2]int{x, y}
	return g.GetC(coord)
}

func (g Grid[V]) GetC(c Coord) V {
	result, found := g.Elements[c]
	if found {
		return result
	} else {
		return g.Default
	}
}

func (g Grid[V]) Contains(x, y int) bool {
	coord := [2]int{x, y}
	_, found := g.Elements[coord]
	return found
}

func (g *Grid[V]) Set(x, y int, value V) (V, bool) {
	coord := [2]int{x, y}
	return g.SetC(coord, value)
}

func (g *Grid[V]) SetC(coord Coord, value V) (V, bool) {
	old, found := g.Elements[coord]
	if value != g.Default {
		g.Elements[coord] = value
	} else {
		delete(g.Elements, coord)
	}

	if !found && value != g.Default {
		if len(g.Elements) == 0 {
			g.MinX = coord[0]
			g.MaxX = coord[0]
			g.MinY = coord[1]
			g.MaxY = coord[1]
		} else {
			g.MinX = min(g.MinX, coord[0])
			g.MaxX = max(g.MaxX, coord[0])
			g.MinY = min(g.MinY, coord[1])
			g.MaxY = max(g.MaxY, coord[1])
		}
	}
	return old, found
}

func (g Grid[V]) Compress() Grid[V] {
	result := NewGrid[V]()
	result.Default = g.Default
	for k, v := range g.Elements {
		result.Set(k[0], k[1], v)
	}
	return result
}

func toGrid(input []string) Grid[string] {
	result := NewGrid[string]()
	for y := 0; y < len(input); y++ {
		row := input[y]
		for x := 0; x < len(row); x++ {
			elem := fmt.Sprintf("%c", row[x])
			result.Set(x, y, elem)
		}
	}
	return result
}

func (g Grid[V]) String() string {
	result := ""
	for y := g.MinY; y <= g.MaxY; y++ {
		for x := g.MinX; x <= g.MaxX; x++ {
			result = fmt.Sprintf("%s%v", result, g.Get(x, y))
		}
		result = result + "\n"
	}
	return result
}

func (g Grid[V]) Clone() Grid[V] {
	result := g
	result.Elements = make(map[Coord]V)
	for k, v := range g.Elements {
		result.Elements[k] = v
	}
	return result
}

func (g Grid[V]) InRangeC(coord Coord) bool {
	return coord[0] >= g.MinX && coord[0] <= g.MaxX && coord[1] >= g.MinY && coord[1] <= g.MaxY
}

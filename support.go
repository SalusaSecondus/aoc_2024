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

type Grid[V any] struct {
	Elements               map[Coord]V
	MinX, MaxX, MinY, MaxY int
	Default                V
}

func NewGrid[V any]() Grid[V] {
	elements := make(map[Coord]V)
	return Grid[V]{
		Elements: elements,
	}
}

func (g Grid[V]) Get(x, y int) V {
	coord := [2]int{x, y}
	result, found := g.Elements[coord]
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

	old, found := g.Elements[coord]
	g.Elements[coord] = value

	if !found {
		if len(g.Elements) == 0 {
			g.MinX = x
			g.MaxX = x
			g.MinY = y
			g.MaxY = y
		} else {
			g.MinX = min(g.MinX, x)
			g.MaxX = max(g.MaxX, x)
			g.MinY = min(g.MinY, y)
			g.MaxY = max(g.MaxY, y)
		}
	}
	return old, found
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

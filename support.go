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
}

func toGrid(input []string) Grid[byte] {
	elements := map[Coord]byte{}
	maxX := 0
	for y := 0; y < len(input); y++ {
		row := input[y]
		for x := 0; x < len(row); x++ {
			elem := row[x]
			coord := [2]int{x, y}
			elements[coord] = elem
		}
		maxX = max(len(row)-1, maxX)
	}
	return Grid[byte]{
		Elements: elements,
		MinX:     0,
		MinY:     0,
		MaxX:     maxX,
		MaxY:     len(input) - 1,
	}
}

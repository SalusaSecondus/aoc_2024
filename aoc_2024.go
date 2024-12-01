package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	goodbye()
	tmp := readInput(0, false)
	for i := 0; i < len(tmp); i++ {
		fmt.Printf("%d: %s\n", i, tmp[i])
	}
}

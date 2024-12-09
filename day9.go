package main

import (
	"strconv"
)

func day9_1(input []string) int {
	fs := parse9(input)
	// fmt.Println("Initial")
	for fs.Compact() {
		// fmt.Println(fs)
		// fmt.Printf("Min %d Max %d\n", fs.FirstEmpty, fs.LastValid)
	}
	return fs.Checksum()
}

func day9_2(input []string) int {
	fs := parse9(input)
	// fmt.Println(fs)
	// fmt.Printf("File %v\n", fs.FilesToCompress)
	// fmt.Printf("Free %v\n", fs.FreeBlocks)
	for fs.Compact2() {
		// fmt.Println(fs)
		// fmt.Printf("Files %v\n", len(fs.FilesToCompress))
		// fmt.Printf("Free %v\n", fs.FreeBlocks)

	}
	return fs.Checksum()
}

type FileSystem9 struct {
	Blocks          []int
	LastValid       int
	FirstEmpty      int
	FreeBlocks      [][2]int // idx, length
	FilesToCompress [][3]int // idx, id, length
}

func (f FileSystem9) String() string {
	result := ""
	for _, id := range f.Blocks {
		if id < 0 {
			result += "."
		} else {
			result += strconv.Itoa(id)
		}
	}
	return result
}

func (f *FileSystem9) Compact() bool {
	if f.FirstEmpty > f.LastValid {
		return false
	}
	if f.Blocks[f.FirstEmpty] >= 0 {
		panic("Trying to write to non-empty block!")
	}
	if f.Blocks[f.LastValid] < 0 {
		panic("Trying to compress empty block!")
	}
	f.Blocks[f.FirstEmpty] = f.Blocks[f.LastValid]
	f.Blocks[f.LastValid] = -1
	f.LastValid--
	for ; f.Blocks[f.LastValid] < 0; f.LastValid-- {
	}
	f.FirstEmpty++
	for ; f.Blocks[f.FirstEmpty] >= 0; f.FirstEmpty++ {
	}
	return true
}

func (f *FileSystem9) Compact2() bool {
	if len(f.FilesToCompress) == 0 {
		return false
	}
	nextFile := f.FilesToCompress[len(f.FilesToCompress)-1]
	// fmt.Printf("Compressing %v\n", nextFile)
	f.FilesToCompress = f.FilesToCompress[:len(f.FilesToCompress)-1]
	for idx := 0; idx < len(f.FreeBlocks); idx++ {
		if f.FreeBlocks[idx][0] > nextFile[0] {
			// fmt.Printf("Could not find to the left. id: %d\n", nextFile[1])
			continue
		}

		if f.FreeBlocks[idx][1] >= nextFile[2] {
			for wIdx := 0; wIdx < nextFile[2]; wIdx++ {
				f.Blocks[f.FreeBlocks[idx][0]+wIdx] = nextFile[1]
			}
			f.FreeBlocks[idx][0] += nextFile[2]
			f.FreeBlocks[idx][1] -= nextFile[2]
			for wIdx := nextFile[0]; wIdx < nextFile[0]+nextFile[2]; wIdx++ {
				f.Blocks[wIdx] = -1
			}

			// Find the existing block immediately before this file to re-allocate
			for fIdx := 0; fIdx < len(f.FreeBlocks); fIdx++ {
				if f.FreeBlocks[fIdx][0]+f.FreeBlocks[fIdx][1] == nextFile[0] {
					// Found it!
					f.FreeBlocks[fIdx][1] += nextFile[2]

					if fIdx < len(f.FreeBlocks)-1 && f.FreeBlocks[fIdx][0]+f.FreeBlocks[fIdx][1] == f.FreeBlocks[fIdx+1][0] {
						f.FreeBlocks[fIdx][1] += f.FreeBlocks[fIdx+1][1]
						f.FreeBlocks[fIdx+1][1] = 0
					}
				}
			}

			return true
		}
	}
	return true
}

func parse9(lines []string) FileSystem9 {
	currId := 0
	blocks := make([]int, 0)
	isFile := true
	maxBlock := 0
	firstEmpty := -1
	freeBlocks := make([][2]int, 0)
	filesToCompress := make([][3]int, 0)
	for _, elem := range lines[0] {
		count := int(elem - '0')
		if isFile {
			filesToCompress = append(filesToCompress, [3]int{len(blocks), currId, count})
			for i := 0; i < count; i++ {
				blocks = append(blocks, currId)
			}
			maxBlock = len(blocks) - 1
			currId++
		} else {
			freeBlocks = append(freeBlocks, [2]int{len(blocks), count})
			if firstEmpty == -1 {
				firstEmpty = len(blocks)
			}
			for i := 0; i < count; i++ {
				blocks = append(blocks, -1)
			}
		}
		isFile = !isFile
	}
	return FileSystem9{
		Blocks:          blocks,
		LastValid:       maxBlock,
		FirstEmpty:      firstEmpty,
		FreeBlocks:      freeBlocks,
		FilesToCompress: filesToCompress,
	}
}

func (f FileSystem9) Checksum() int {
	result := 0
	for idx, val := range f.Blocks {
		if val >= 0 {
			result += idx * val

		}
	}
	return result
}

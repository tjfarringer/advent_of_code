package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

type GridPosition struct {
	row, col int
}

// Part 1:  How many foods are in the fresh lists?
// Simple map and then loop is too slow.  Maybe store first,end within the map ?
func main() {
	start := time.Now()
	numSplitsP1 := 0
	numPathsP2 := 0
	var grid [][]rune

	file, err := os.Open("7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read input grid
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}

	// Part 1:  Want to know how many times the beam splits
	// Part 2:  How many timelines are there?
	//
	beamCol := make(map[int]bool)
	beamColCount := make(map[int]int)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'S' {
				beamCol[col] = true
				// 1 path to this square so far
				beamColCount[col] = 1
			}
			if row == 6 && col == 7 {
				println("debugging")
			}
			if grid[row][col] == '^' && beamCol[col] {
				numSplitsP1++
				beamCol[col] = false

				if col-1 >= 0 {
					beamCol[col-1] = true
					pathCount, exists := beamColCount[col-1]
					if !exists {
						pathCount = beamColCount[col]
					} else {
						pathCount = pathCount + beamColCount[col]
					}
					beamColCount[col-1] = pathCount
				}
				if col+1 < len(grid[row]) {
					beamCol[col+1] = true
					pathCount, exists := beamColCount[col+1]
					if !exists {
						pathCount = beamColCount[col]
					} else {
						pathCount = pathCount + beamColCount[col]
					}
					beamColCount[col+1] = pathCount
				}
				beamColCount[col] = 0
			}
		}
	}
	for _, value := range beamColCount {
		numPathsP2 += value
	}

	log.Printf("Number of splits -- part 1: %d\n", numSplitsP1)
	log.Printf("Number of paths -- part 2: %d\n", numPathsP2)
	log.Printf("Execution Time: %s\n", time.Since(start))
}

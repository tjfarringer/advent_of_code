package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

// forklifts can only access a roll of paper if there are
// fewer than four rolls of paper in the eight adjacent positions

func full_square_check(grid [][]rune, row, col int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nRow, nCol := row+i, col+j
			if nRow >= 0 && nRow < len(grid) && nCol >= 0 && nCol < len(grid[0]) {
				if grid[nRow][nCol] == '@' {
					count++
				}
			}
		}
	}
	return count
}

// func sliding_window_count(grid [][]rune, row, col, windowSize int) int {}

// Part 1:  How many rolls of paper are accessible to forklifts?
// Part 2:  Now rolls can be removed.  Count how many can be removed.
func main() {
	start := time.Now()
	var grid [][]rune
	accessiblePaper := 0
	removablePaperGlobal := 0
	removablePaperLoop := 1

	file, err := os.Open("4.txt")
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

	// Part 1
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '@' {
				adjacentCount := full_square_check(grid, r, c)
				if adjacentCount < 4 {
					accessiblePaper++
				}
			}
		}
	}
	// Part 2
	for removablePaperLoop > 0 {
		removablePaperLoop = 0
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[0]); c++ {
				if grid[r][c] == '@' {
					adjacentCount := full_square_check(grid, r, c)
					if adjacentCount < 4 {
						removablePaperLoop++
						removablePaperGlobal++
						grid[r][c] = '.' // remove the roll
					}
				}
			}
		}
	}

	log.Printf("Paper -- part 1: %d\n", accessiblePaper)
	log.Printf("Paper -- part 2: %d\n", removablePaperGlobal)
	log.Printf("Execution Time: %s\n", time.Since(start))
}

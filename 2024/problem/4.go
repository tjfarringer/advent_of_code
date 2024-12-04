package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func inGrid(grid [][]string, row int, col int) bool {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return false
	}
	return true
}

func checkForRestOfWord(r int, c int, rI int, cI int, grid [][]string) int {
	// If out of bounds return 0
	if !inGrid(grid, r+(2*rI), c+(2*cI)) || !inGrid(grid, r+(3*rI), c+(3*cI)) {
		return 0
	}
	// If rest of word then return 1
	if grid[r+(2*rI)][c+(2*cI)] == "A" && grid[r+(3*rI)][c+(3*cI)] == "S" {
		return 1
	}
	return 0
}

func checkXMas(grid [][]string, row int, col int) int {
	// The "A" cannot be on the edge of the grid
	if !inGrid(grid, row-1, col-1) || !inGrid(grid, row+1, col+1) {
		return 0
	}
	//  Check every other row
	for rI := -1; rI <= 1; rI += 2 {
		//	Check col -1 and 1
		for cI := -1; cI <= 1; cI += 2 {
			// Not the most elegant but putting together for speed
			if (grid[row+rI][col+cI] == "M" && grid[row-rI][col-cI] == "S") || (grid[row+rI][col+cI] == "S" && grid[row-rI][col-cI] == "M") {
				if (grid[row+rI][col+(cI*-1)] == "M" && grid[row+(rI*-1)][col+cI] == "S") || (grid[row+rI][col+(cI*-1)] == "S" && grid[row+(rI*-1)][col+cI] == "M") {
					return 1
				}
			}
		}
	}
	return 0
}

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2024/problem/puzzleInput/day_04_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Build grid
	uniqueWords := 0
	scanner := bufio.NewScanner(file)
	var searchGrid [][]string
	// Read each line from the file
	for scanner.Scan() {
		// Convert text to upper
		word := strings.ToUpper(scanner.Text())
		// Convert the word into a slice of characters
		var charSlice []string
		for _, r := range word {
			// Convert each rune to string and append to the slice
			charSlice = append(charSlice, string(r))
		}

		// Append the character slice to the grid
		searchGrid = append(searchGrid, charSlice)
	}

	for row := 0; row < len(searchGrid); row++ {
		for col := 0; col < len(searchGrid[row]); col++ {
			if searchGrid[row][col] == "X" {
				for rI := -1; rI <= 1; rI++ {
					for cI := -1; cI <= 1; cI++ {
						//if (rI == 0 && cI == 0) {
						//  Same square as what we are checking on
						//  Optimization for later
						//	continue
						//}
						if inGrid(searchGrid, row+rI, col+cI) && searchGrid[row+rI][col+cI] == "M" {
							uniqueWords += checkForRestOfWord(row, col, rI, cI, searchGrid)
						}
					}
				}
			}
		}
	}

	// Scan again for part 2
	xMasCount := 0
	for row := 0; row < len(searchGrid); row++ {
		for col := 0; col < len(searchGrid[row]); col++ {
			if searchGrid[row][col] == "A" {
				xMasCount += checkXMas(searchGrid, row, col)
			}
		}
	}
	log.Printf("Part 1 answer: %d", uniqueWords)
	log.Printf("Part 2 answer: %d", xMasCount)

}

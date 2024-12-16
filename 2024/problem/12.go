package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// need to know size and perimeter of the region

type iDireciton struct {
	colInc int
	rowInc int
}

func insideGrid(grid [][]string, row int, col int) bool {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return false
	}
	return true
}

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2024/problem/puzzleInput/day_12_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	startTime := time.Now()

	var searchGrid [][]string
	var visitedGrid [][]bool

	var allDirections []iDireciton
	allDirections = append(allDirections, iDireciton{colInc: 0, rowInc: -1})
	allDirections = append(allDirections, iDireciton{colInc: 0, rowInc: 1})
	allDirections = append(allDirections, iDireciton{colInc: 1, rowInc: 0})
	allDirections = append(allDirections, iDireciton{colInc: -1, rowInc: 0})

	// Step 1:  Parse the input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var lineS []string
		var boolS []bool
		line := scanner.Text()
		for _, str := range line {
			lineS = append(lineS, string(str))
			boolS = append(boolS, false)
		}

		searchGrid = append(searchGrid, lineS)
		visitedGrid = append(visitedGrid, boolS)
	}

	// 1 -- loop until you find an unexplored tile -- if none then done
	answer := 0
	for i := 0; i < len(searchGrid); i++ {
		for j := 0; j < len(searchGrid[i]); j++ {
			if visitedGrid[i][j] {
				continue
			} else {
				//  This is the first unexplored square
				//islandArea, islandPerm, corners
				islandArea, _, corners := exploreIsland(i, j, visitedGrid, searchGrid, allDirections)
				fmt.Printf("Num corners %s\n", corners)
				fmt.Printf("Cost %s\n", islandArea*(corners))
				answer += (islandArea * (corners))
				//answer += (islandArea * islandPerm)
			}
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Program took %s to run\n", elapsedTime)
	log.Printf("Answer: %d", answer)
}

func exploreIsland(row int, col int, visitedGrid [][]bool, searchGrid [][]string, allDirections []iDireciton) (int, int, int) {
	visitedGrid[row][col] = true
	// This should only be 1 for the current square, right ?
	islandArea := 1
	islandPerm := 4
	corners := 0

	for i := 0; i < len(allDirections); i++ {
		if insideGrid(searchGrid, row+allDirections[i].rowInc, col+allDirections[i].colInc) {
			if searchGrid[row][col] == searchGrid[row+allDirections[i].rowInc][col+allDirections[i].colInc] {

				// If the neighbor square is equal then 1 less from perimeter
				islandPerm -= 1
				if visitedGrid[row+allDirections[i].rowInc][col+allDirections[i].colInc] {
					continue
				} else {
					aI, pI, cI := exploreIsland(row+allDirections[i].rowInc, col+allDirections[i].colInc, visitedGrid, searchGrid, allDirections)
					islandArea += aI
					islandPerm += pI
					corners += cI
				}
			} else {

			}
		} else {
			//	out of bounds
		}
	}
	corners += checkCorners(row, col, searchGrid)
	return islandArea, islandPerm, corners
}

func checkCorners(row int, col int, searchGrid [][]string) int {
	numCorners := 0
	var cornerDirections []iDireciton
	cornerDirections = append(cornerDirections, iDireciton{colInc: -1, rowInc: -1})
	cornerDirections = append(cornerDirections, iDireciton{colInc: 1, rowInc: 1})
	cornerDirections = append(cornerDirections, iDireciton{colInc: 1, rowInc: -1})
	cornerDirections = append(cornerDirections, iDireciton{colInc: -1, rowInc: 1})

	// TODO:  Gap between needs to be different
	for i := 0; i < len(cornerDirections); i++ {
		// if inside grid
		if insideGrid(searchGrid, row+cornerDirections[i].rowInc, col+cornerDirections[i].colInc) {
			// and equals to the same
			if searchGrid[row+cornerDirections[i].rowInc][col] != searchGrid[row][col] &&
				searchGrid[row][col+cornerDirections[i].colInc] != searchGrid[row][col] {
				numCorners++
			}
			// concave corner -- like an L
			if searchGrid[row+cornerDirections[i].rowInc][col] == searchGrid[row][col] &&
				searchGrid[row][col+cornerDirections[i].colInc] == searchGrid[row][col] &&
				searchGrid[row+cornerDirections[i].rowInc][col+cornerDirections[i].colInc] != searchGrid[row][col] {
				numCorners++
			}
		}
	}
	return numCorners
}

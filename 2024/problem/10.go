package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Input -- height at each position using a scale (0 lowest; 9 highest)
// Good hiking path is as long as possible; even, gradual, uphill slope
//  starts at 0; ends at 9; always increases by a height of exactly 1;
// Only left, right, up, down (no diag)

//  Trailhead:  any position that starts one or more hiking trails
//  Trailhead score is the num of 9-height positions reachable from the trailhead

//  Goal:  sum of all trailhead scores

//  Loop through, find 0's, start DFS,

type HikingPosition struct {
	col int
	row int
}

type IncrementalDirection struct {
	colInc int
	rowInc int
}

func positionInGrid(grid [][]int, row int, col int) bool {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return false
	}
	return true
}

func findNextMove(grid [][]int, position HikingPosition, allDirections []IncrementalDirection) []HikingPosition {
	allPotentialMoves := []HikingPosition{}
	for i := 0; i < len(allDirections); i++ {
		if positionInGrid(grid, position.row+allDirections[i].rowInc, position.col+allDirections[i].colInc) {
			if grid[position.row][position.col]+1 == grid[position.row+allDirections[i].rowInc][position.col+allDirections[i].colInc] {
				allPotentialMoves = append(allPotentialMoves, HikingPosition{col: position.col + allDirections[i].colInc, row: position.row + allDirections[i].rowInc})
			}
		}
	}
	return allPotentialMoves
}

func dfs(searchGrid [][]int, currentPosition HikingPosition, allDirections []IncrementalDirection) (map[HikingPosition]struct{}, int) {
	uniquePaths := 0
	//  No such thing as a set in Go so we need to use a map
	peakSet := make(map[HikingPosition]struct{})
	// Return early due
	if searchGrid[currentPosition.row][currentPosition.col] == 9 {
		//return 1
		peakSet[currentPosition] = struct{}{}
		//currentPath = append(currentPath, HikingPosition{row: currentPosition.row, col: currentPosition.col})
		return peakSet, 1
	}
	//// Search until we find a 9
	//for searchGrid[currentPosition.row][currentPosition.col] != 9 {
	// Get all possible next moves
	nextMoves := findNextMove(searchGrid, currentPosition, allDirections)
	if len(nextMoves) == 0 {
		// No possible next moves
		return peakSet, 0
	} else {
		// TODO:  check that this is working correctly
		// TODO:  Search until val is 9 or there are no more moves
		for i := 0; i < len(nextMoves); i++ {
			currentPosition = nextMoves[i]
			// Add position to current path and pass down
			//currentPath = append(currentPath, currentPosition)
			// Search each of the next moves
			// If you get to a 9 there will be no next move
			//totalCount += dfs(searchGrid, currentPosition, allDirections)
			foundPeaks, uniquePathCount := dfs(searchGrid, currentPosition, allDirections)
			for peak := range foundPeaks {
				peakSet[peak] = struct{}{}
			}
			uniquePaths += uniquePathCount
			//currentPath = append(currentPath, peakPath...)
		}
	}
	//return totalCount
	return peakSet, uniquePaths
}

func buildTrailMap(searchGrid [][]int, origin HikingPosition, allDirections []IncrementalDirection) int {
	numTrails := 0
	numPeaks := 0

	foundPeaks, uniquePath := dfs(searchGrid, origin, allDirections)
	numPeaks += len(foundPeaks)
	numTrails += uniquePath
	return numTrails
}

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2024/problem/puzzleInput/day_10_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	startTime := time.Now()

	potentialTrailHeads := make(map[HikingPosition]int)

	var allDirections []IncrementalDirection
	allDirections = append(allDirections, IncrementalDirection{colInc: 0, rowInc: -1})
	allDirections = append(allDirections, IncrementalDirection{colInc: 0, rowInc: 1})
	allDirections = append(allDirections, IncrementalDirection{colInc: 1, rowInc: 0})
	allDirections = append(allDirections, IncrementalDirection{colInc: -1, rowInc: 0})

	var searchGrid [][]int
	rowNum := 0

	// Step 1:  Parse the input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		colNum := 0
		line := scanner.Text()

		// Convert each string number to an integer and store it in a slice
		var intNumbers []int

		for _, str := range line {
			num, _ := strconv.Atoi(string(str))
			intNumbers = append(intNumbers, num)
			if num == 0 {
				potentialTrailHeads[HikingPosition{col: colNum, row: rowNum}] = 0
			}
			colNum++
		}

		searchGrid = append(searchGrid, intNumbers)
		rowNum++
	}

	// Step 2:  Go through all potential trailheads; search the grid; update the score
	totalTrailheadScore := 0
	for k, v := range potentialTrailHeads {
		potentialTrailHeads[k] = v + buildTrailMap(searchGrid, k, allDirections)
		totalTrailheadScore += potentialTrailHeads[k]
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Program took %s to run\n", elapsedTime)
	//  Part 1 they only wanted the number of distinct peaks
	log.Printf("Answer: %d", totalTrailheadScore)
	//	Part 2 they want the number of distinct paths
}

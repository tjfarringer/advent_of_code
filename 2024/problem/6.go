package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Goal:  How many positions will the guard visit before leaving the map ?
type Direction struct {
	colInc int
	rowInc int
}

type Position struct {
	col int
	row int
}

// https://stackoverflow.com/questions/8307478/how-to-find-out-element-position-in-slice
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func guardInGrid(grid [][]string, row int, col int) bool {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return false
	}
	return true
}

func moveGuard(grid [][]string, guardPosition Position, currGuardDirection string, guardDirection map[string]Direction, possibleGuardDirections []string, visitedPositions int, visitedPositionMap map[Position]Direction) (Position, string, int) {
	value, _ := guardDirection[currGuardDirection]
	colInc := value.colInc
	rowInc := value.rowInc

	guardPosition.col += colInc
	guardPosition.row += rowInc

	// If out of bounds that's fine just return
	// Only thing to change is if the new square is blocked
	if !guardInGrid(grid, guardPosition.row, guardPosition.col) {
		return guardPosition, currGuardDirection, visitedPositions
	} else if grid[guardPosition.row][guardPosition.col] != "." {
		// Square is blocked, go back to original position
		guardPosition.col -= colInc
		guardPosition.row -= rowInc
		// new direction
		return guardPosition, possibleGuardDirections[(SliceIndex(len(possibleGuardDirections), func(i int) bool { return possibleGuardDirections[i] == currGuardDirection })+1)%4], visitedPositions
	}
	if _, exists := visitedPositionMap[guardPosition]; !exists {
		// Only increment if we haven't visited yet
		visitedPositions++
	}
	return guardPosition, currGuardDirection, visitedPositions

}

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2024/problem/puzzleInput/day_06_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	possibleGuardDirections := []string{"^", ">", "v", "<"}
	guardDirection := make(map[string]Direction)
	guardDirection["^"] = Direction{colInc: 0, rowInc: -1}
	guardDirection["v"] = Direction{colInc: 0, rowInc: 1}
	guardDirection[">"] = Direction{colInc: 1, rowInc: 0}
	guardDirection["<"] = Direction{colInc: -1, rowInc: 0}

	currentGuardDirection := "^"
	var guardPosition Position
	var searchGrid [][]string
	visitedPositionMap := make(map[Position]Direction)
	visitedPositions := 1
	barrierCount := 0
	// Read each line from the file
	scanner := bufio.NewScanner(file)
	rowNum := 0
	for scanner.Scan() {
		// Convert text to upper
		word := strings.ToUpper(scanner.Text())
		// Convert the word into a slice of characters
		var charSlice []string
		colNum := 0
		for _, r := range word {
			// If we found the guard char set it
			_, exists := guardDirection[string(r)]
			if exists {
				currentGuardDirection = string(r)
				guardPosition = Position{col: colNum, row: rowNum}
				visitedPositionMap[guardPosition] = guardDirection[currentGuardDirection]
				charSlice = append(charSlice, string("."))
			} else {
				// Convert each rune to string and append to the slice
				charSlice = append(charSlice, string(r))
			}
			colNum++
		}
		// Append the character slice to the grid
		searchGrid = append(searchGrid, charSlice)
		rowNum++
	}

	//	Brute force algo for Part 1 :)
	for guardInGrid(searchGrid, guardPosition.row, guardPosition.col) {
		guardPosition, currentGuardDirection, visitedPositions = moveGuard(searchGrid, guardPosition, currentGuardDirection, guardDirection, possibleGuardDirections, visitedPositions, visitedPositionMap)
		visitedPositionMap[guardPosition] = guardDirection[currentGuardDirection]

		// Check for part 2
		potGuardPosition := Position{guardPosition.col + 1, guardPosition.row}
		if guardInGrid(searchGrid, potGuardPosition.row, potGuardPosition.col) {
			// If direction the same then I think you can put a barrier here
			if visitedPositionMap[guardPosition] == visitedPositionMap[potGuardPosition] && guardInGrid(searchGrid, guardPosition.row+guardDirection[currentGuardDirection].rowInc, guardPosition.col+guardDirection[currentGuardDirection].colInc) {
				barrierCount += 1
			}
		}
	}

	// Part 2 -- how many positions can an obstruction be placed to cause an infinite loop ? -- cannot place at the current position
	//  Idea -- if the position to the immediate right has been visited, and has the same direction as the current one, then maybe that causes an infinite loop ?
	//for guardInGrid(searchGrid, guardPosition.row, guardPosition.col) {
	//	guardPosition, currentGuardDirection, visitedPositions = moveGuard(searchGrid, guardPosition, currentGuardDirection, guardDirection, possibleGuardDirections, visitedPositions, visitedPositionMap)
	//	visitedPositionMap[guardPosition] = guardDirection[currentGuardDirection]
	//}

	log.Printf("Part 1 answer: %d", visitedPositions)
	log.Printf("Part 2 answer: %d", barrierCount)
}

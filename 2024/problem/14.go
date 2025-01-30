package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// robots move in a straight line
// p=x,y
// if they go over the edge they teleport to the otherside

//  Where will the robots be in 100 rounds ?
//  What if you multiply by 100, then take mod by the length of the grid, then move the robot by that amount

//count the number of robots in each quadrant after 100 seconds.
// Robots that are exactly in the middle (horizontally or vertically) don't count

// Multiply the safety score for each quad

type RobotPosition struct {
	col int
	row int
}
type RobotMoves struct {
	colDiff int
	rowDiff int
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2024/problem/puzzleInput/day_14_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	robotPosition := make(map[int]RobotPosition)
	robotMovement := make(map[int]RobotMoves)
	lineNum := 0

	quad1 := 0
	quad2 := 0
	quad3 := 0
	quad4 := 0

	gridWidth := 101  //101
	gridHeight := 103 //103

	// Parse input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		cord := parts[0]
		cordParts := strings.Split(cord, ",")
		// Convert strings to integers
		x, _ := strconv.Atoi(strings.Split(cordParts[0], "=")[1])
		y, _ := strconv.Atoi(cordParts[1])
		pos := RobotPosition{col: x, row: y}
		robotPosition[lineNum] = pos

		moveDiff := parts[1]
		moveDiffParts := strings.Split(moveDiff, ",")
		// Convert strings to integers
		x, _ = strconv.Atoi(strings.Split(moveDiffParts[0], "=")[1])
		y, _ = strconv.Atoi(moveDiffParts[1])
		movePos := RobotMoves{colDiff: x, rowDiff: y}
		robotMovement[lineNum] = movePos
		lineNum++
	}

	// Find new positions for all robots
	for index := range lineNum {
		currentPosition := robotPosition[index]
		movement := robotMovement[index]

		newCol := mod((currentPosition.col + (movement.colDiff * 100)), gridWidth)
		newRow := mod((currentPosition.row + (movement.rowDiff * 100)), gridHeight)

		if newCol == int(math.Floor(float64(gridWidth/2))) || newRow == int(math.Floor(float64(gridHeight/2))) {
			continue
		} else if newCol < int(math.Floor(float64(gridWidth/2))) && newRow < int(math.Floor(float64(gridHeight/2))) {
			quad1++
		} else if newCol > int(math.Floor(float64(gridWidth/2))) && newRow < int(math.Floor(float64(gridHeight/2))) {
			quad2++
		} else if newCol > int(math.Floor(float64(gridWidth/2))) && newRow > int(math.Floor(float64(gridHeight/2))) {
			quad3++
		} else if newCol < int(math.Floor(float64(gridWidth/2))) && newRow > int(math.Floor(float64(gridHeight/2))) {
			quad4++
		}
	}

	// sum up
	log.Printf("Answer: %d", (quad1 * quad2 * quad3 * quad4))

	//	space which is 101 tiles wide and 103 tiles tall

}

// 10x10
// at 4x4 -- because it's 0 indexed
// Moving up 21 spaces == -17

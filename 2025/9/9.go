package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

func main() {
	start := time.Now()

	file, err := os.Open("9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gridPoints := []Point{}
	// loop through and add points to a list
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		gridPoints = append(gridPoints, Point{x, y})
	}

	// Part 1: Find the largest rectangle that can be formed by the points
	largestRectSize := 0
	for _, point := range gridPoints {
		for _, otherPoint := range gridPoints {
			if point.x == otherPoint.x && point.y == otherPoint.y {
				// i.e. the same point
				continue
			} else {
				rectSize := (point.x - otherPoint.x + 1) * (point.y - otherPoint.y + 1)
				if rectSize > largestRectSize {
					largestRectSize = rectSize
				}
			}
		}
	}
	// Part 2:  Can only use red or green tiles
	// red tiles are defined in the input
	// every red tile is connected to the red tile before and after it by a straight line of green tiles
	// Every tile inside this loop is also green

	log.Printf("Largest rectangle size -- part 1: %d", largestRectSize)
	log.Printf("Execution Time: %s\n", time.Since(start))
}

package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Presents:
// index:
// # is part of the shape and . is not

// Second section lists the regions under the trees
// Each line starts with the width and length of the region
// The rest of the line describes the presents that need to fit into that region
// by listing the quantity of each shape of present

// Presents can be rotated and flipped
// but they have to always be placed perfectly on the grid

// Goal:
// The Elves need to know how many of the regions can fit the presents listed

type Present struct {
	ID   int
	Grid [][]bool // true is part of the present, false is not
	Size int
}

type Tree struct {
	width    int
	height   int
	Presents []int
}

func main() {
	start := time.Now()
	// Create data structures
	regionsThatFit := 0
	presentSize := make(map[int]int)

	data, err := os.ReadFile("12.txt")
	if err != nil {
		log.Fatal(err)
	}

	sections := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	// Parse present sizes
	for _, section := range sections {
		lines := strings.Split(section, "\n")
		if strings.Contains(lines[0], "x") {
			// Tree -- skip
			continue
		}
		id, _ := strconv.Atoi(strings.TrimSuffix(lines[0], ":"))
		presentSize[id] = strings.Count(section, "#") + strings.Count(section, ".")
	}

	// Parse trees
	for _, line := range strings.Split(string(data), "\n") {
		if !strings.Contains(line, "x") {
			continue
		}
		parts := strings.SplitN(line, ": ", 2)
		dims := strings.Split(parts[0], "x")
		w, _ := strconv.Atoi(dims[0])
		h, _ := strconv.Atoi(dims[1])

		// Parse present counts
		counts := strings.Fields(parts[1])
		totalPresentArea := 0
		for id, countStr := range counts {
			count, _ := strconv.Atoi(countStr)
			totalPresentArea += presentSize[id] * count
		}

		gridArea := w * h

		// Count if presents fit by area
		// This is a cheeky way to solve part 1;
		if totalPresentArea <= gridArea {
			regionsThatFit++
		}

		// Part 2:

		

	}

	log.Printf("Regions that fit: %d", regionsThatFit)
	log.Printf("Execution Time: %s\n", time.Since(start))
}

package main

import (
	"bufio"
	"embed"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//go:embed puzzleInput/day_01_input.txt
var content embed.FS

func main() {
	file, err := os.Open("puzzleInput/day_01_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a slice to store integers
	var leftList []int
	var rightList []int

	// Read file line by line
	// Store numbers in the slice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		//log.Printf("p1: %q", parts[0])
		//log.Printf("p2: %q", parts[1])
		//log.Printf("p3: %q", parts[2])

		if len(parts) != 2 {
			log.Printf("bad line -- skipping: %q", line)
			continue
		}

		// convert into integers
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("Error converting first number: %v\n", err)
			continue
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Printf("Error converting second number: %v\n", err)
			continue
		}
		// Append the number to the slice
		leftList = append(leftList, num1)
		rightList = append(rightList, num2)
	}

	// Sort the slices in ascending order
	sort.Ints(leftList)
	sort.Ints(rightList)

	// Loop through and find the diff
	totalDiff := 0
	for i := 0; i < len(leftList); i++ {
		absDiff := math.Abs(float64(leftList[i] - rightList[i]))
		totalDiff += int(absDiff)
	}
	log.Printf("Part 1 answer -- total diff: %d", totalDiff)

	//	PART 2:  How often does each number in leftList appear in the rightList ?
	//	Calc a similarity score by adding up each number in the Left after * by num of times it appears in the right List

	//	Step 1 -- create a count map from rightList
	countMap := make(map[int]int)

	for _, num := range rightList {
		countMap[num]++
	}
	//  Step 2 -- loop through left list, *, and add to final val
	//	What happens if not in list ?
	simScore := 0
	for _, num := range leftList {
		countVal, found := countMap[num]
		if found {
			simScore += (num * countVal)
		} else {
			simScore += 0
		}
	}
	log.Printf("Part 2 answer -- sim score: %d", simScore)

}

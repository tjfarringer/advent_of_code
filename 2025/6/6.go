package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Part 1:  Need to parse the numbers and operators from the file
// Then do the operation and sum up the results
func main() {
	start := time.Now()
	part1Answer := 0
	part2Answer := 0
	numbers := make(map[int][]int)
	operatorMap := make(map[int]rune)

	// open file
	file, err := os.Open("6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		// Split the line into parts
		parts := strings.Fields(line)
		for i, part := range parts {
			if part == "+" {
				operatorMap[i] = '+'
			} else if part == "*" {
				operatorMap[i] = '*'
			} else {
				num, err := strconv.Atoi(part)
				if err != nil {
					log.Fatal(err)
				}
				numbers[i] = append(numbers[i], num)
			}
		}
	}

	// Compute the answer
	for problem := 0; problem < len(numbers); problem++ {
		problemAnswer := 1
		for _, num := range numbers[problem] {
			if operatorMap[problem] == '+' {
				problemAnswer += num
			} else if operatorMap[problem] == '*' {
				problemAnswer *= num
			}
		}
		part1Answer += problemAnswer
		if operatorMap[problem] == '+' {
			part1Answer -= 1
		}
	}

	// Part 2:  Different way to compile the numbers...
	// read file line by line
	// After first scanner, seek back to start
	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	problemSpace := []string{}
	for scanner2.Scan() {
		text := scanner2.Text()
		problemSpace = append(problemSpace, text)
	}
	// Find problem operators
	operators := strings.Fields(problemSpace[len(problemSpace)-1])
	slices.Reverse(operators)

	// Loop through all values in the line
	problems := make(map[int][]int)
	problemPointer := 0
	problemString := ""

	for numIndex := len(problemSpace[0]) - 1; numIndex >= 0; numIndex-- {
		allNotBlank := false
		// Set string to the first val
		problemString = string(problemSpace[0][numIndex])
		if problemSpace[0][numIndex] != ' ' {
			allNotBlank = true
		}
		// Start at 1 to skip the first line
		// Compile the number
		for subValIndex := 1; subValIndex < len(problemSpace)-1; subValIndex++ {
			problemString += string(problemSpace[subValIndex][numIndex])
			// if all values are blank then we move to the next problem
			if problemSpace[subValIndex][numIndex] != ' ' {
				allNotBlank = true
			}
		}
		if allNotBlank {
			num, err := strconv.Atoi(strings.TrimSpace(problemString))
			if err != nil {
				log.Fatal(err)
			}
			// Store number entry attached to the problem
			problems[problemPointer] = append(problems[problemPointer], num)
		} else {
			// At this point we need to move onto the next problem
			problemPointer++
		}
	}

	for problemNum, numbers := range problems {
		operator := operators[problemNum]
		var result int
		if operator == "+" {
			result = 0
			for _, num := range numbers {
				result += num
			}
		} else if operator == "*" {
			result = 1
			for _, num := range numbers {
				result *= num
			}
		}

		part2Answer += result
	}

	log.Printf("Part 1 answer: %d", part1Answer)
	log.Printf("Part 2 answer: %d", part2Answer)
	log.Printf("Execution Time: %s\n", time.Since(start))
}

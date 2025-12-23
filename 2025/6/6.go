package main

import (
	"bufio"
	"log"
	"os"
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
	problemLength := make(map[int]int)
	numbers := make(map[int][]int)
	// numbersP2 := make(map[int][]string)
	operatorMap := make(map[int]rune)
	operatorMapP2 := make(map[int]rune)

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
	problemPointer := 0
	positionSum := 1
	var lastLine string
	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	problemSpace := []string{}
	for scanner2.Scan() {
		text := scanner2.Text()
		problemSpace = append(problemSpace, text)
	}
	
	for i := len(line) - 1; i >= 0; i-- {
		char := string(line[i])



	// Loop through and figure out how many positions each problem is
	for i := 1; i < len(lastLine); i++ {
		if string(lastLine[i]) == " " {
			positionSum += 1
			continue
		} else if string(lastLine[i]) != " " {
			// subtract 1 because we don't want to
			// count the space between problems
			problemLength[problemPointer] = (positionSum - 1)
			positionSum = 1
			problemPointer++
		}
	}

	// now compile the problems
	file.Seek(0, 0)
	scanner3 := bufio.NewScanner(file)
	for scanner3.Scan() {
		line := scanner3.Text()
		if line == "" {
			continue
		}
		for i := len(line) - 1; i >= 0; i-- {
			char := string(line[i])
			problemLen := problemLength[problemPointer-1]
			if char == "+" {
				operatorMapP2[problemPointer] = '+'
			} else if char == "*" {
				operatorMapP2[problemPointer] = '*'
			} else {
				// numbersP2[problemPointer][problemLen] = append(numbersP2[problemPointer][problemLen], char)
			}

			problemLen--
			if problemLen < 0 {
				problemPointer--
				problemLen = problemLength[problemPointer-1]
			}
		}
	}

	log.Printf("Part 1 answer: %d", part1Answer)
	log.Printf("Part 2 answer: %d", part2Answer)
	log.Printf("Execution Time: %s\n", time.Since(start))
}

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Goal - multiply some numbers
// mul(44,46) multiplies 44 by 46
// Many invalid chars that should be ignored
// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))

// Idea: scan until you find m, then check if the following chars exactly match the required pattern

func toInteger(s string) (int, bool) {
	// Try to convert the string to an integer
	i, err := strconv.Atoi(s)
	if err != nil {
		// If there's an error, return 0 and false
		return 0, false
	}
	// If successful, return the integer and true
	return i, true
}

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2024/problem/puzzleInput/day_03_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file line by line
	answer := 0
	multiBool := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()

		// Loop through each character of the string
		for i := 0; i < len(instruction)-7; i++ {
			if instruction[i:i+7] == "don't()" {
				multiBool = false
			}
			if instruction[i:i+4] == "do()" {
				multiBool = true
			}
			if instruction[i:i+4] == "mul(" {
				// Find the index of the next comma
				nextCommaIndex := strings.Index(instruction[i+4:], ",")
				// If -1 then not valid instruction
				if nextCommaIndex != -1 {
					// Before the first comma
					part1 := instruction[i+4:][:nextCommaIndex]
					// After the first comma
					part2 := instruction[i+4:][nextCommaIndex+1:]

					if part1Num, ok := toInteger(part1); ok {
						// Find the index of the next comma
						nextParIndex := strings.Index(part2, ")")
						// Before the first comma
						part1Par := part2[:nextParIndex]
						if part1ParNum, ok := toInteger(part1Par); ok {
							if multiBool {
								// Only add to the answer if the multi flag is set to true
								answer += (part1Num * part1ParNum)
							}
						}
					}
				}
			}
		}
		log.Printf("Answer: %d", answer)
	}
}

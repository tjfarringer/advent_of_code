package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2023/puzzle_input/day_01_puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	myMap := make(map[rune]bool)

	// Rune from a string is diff than rune from an integer
	xS := string("0123456789")
	for _, char := range xS {
		myMap[char] = true
	}

	finalAnswer := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {

		println(scanner.Text())
		// Loop over each line
		// pos, char
		intermediateAnswer := ""
		for _, char := range scanner.Text() {

			if myMap[char] {
				//do something here
				intermediateAnswer = intermediateAnswer + string(char)
				break
			}
		}
		// loop backwards to get last number value
		rs := []rune(scanner.Text())
		max := len(rs) - 1
		for i := range rs {
			r := rs[max-i]
			if _, ok := myMap[r]; ok {
				//do something here
				intermediateAnswer = intermediateAnswer + string(r)
				break
			}
		}
		// string to int
		i, err := strconv.Atoi(intermediateAnswer)
		if err != nil {
			panic(err)
		} else {
			finalAnswer += i
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println(finalAnswer)
}

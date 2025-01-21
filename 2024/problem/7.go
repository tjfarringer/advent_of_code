package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//which test values could possibly be produced by placing any combination of operators into their
// calibration equations
//  Each line is an equation;  test value is on left of :
//  Operators are eval left to right

//  Only two types of operators (+, *)
//  Which equations could possibly be true ?

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2024/problem/puzzleInput/day_07_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	finalAnswer := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		equation := scanner.Text()
		answer, values, ok := strings.Cut(equation, ":")
		target, _ := strconv.Atoi(answer)
		if !ok {
			//	TODO:  panic ?
		}
		strs := strings.Split(values, " ")

		// Create a slice with our values
		var numbers []int
		for _, str := range strs {
			if str != "" {
				num, err := strconv.Atoi(str)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					return
				}
				numbers = append(numbers, num)
			}
		}

		//	To solve I think we want to do a backtracking algorithm
		if backtracking(target, numbers, 1, numbers[0]) {
			finalAnswer += target
		}
	}
	log.Printf("Answer: %d", finalAnswer)
}

func backtracking(target int, numbers []int, depth int, i_val int) bool {
	if (depth >= len(numbers)) && (i_val != target) {
		return false
	} else if (depth >= len(numbers)) && (i_val == target) {
		return true
	} else if i_val+numbers[depth] <= target {
		// Trying addition
		if backtracking(target, numbers, depth+1, i_val+numbers[depth]) {
			return true
		}
	}
	// Try multiplication
	if i_val*numbers[depth] <= target {
		if backtracking(target, numbers, depth+1, i_val*numbers[depth]) {
			return true
		}
	}
	// Try concatenation

	// Convert integers to strings
	iValS := strconv.Itoa(i_val)
	nextIntS := strconv.Itoa(numbers[depth])
	concatResult := iValS + nextIntS
	concatResultI, _ := strconv.Atoi(concatResult)
	if concatResultI <= target {
		if backtracking(target, numbers, depth+1, concatResultI) {
			return true
		}
	}
	// base case is to return false
	return false
}

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Pages need to be printed in a specific order
//X|Y means that if both page number X and page number Y are to be produced
//as part of an update,
//page number X must be printed at some point before page number Y

//input 1: rules
//input 2: pages to print

//Goal:  which updates are already in the right order ?
//  Add up the middle page # in the row and return the sum

// Idea:
//  Create two maps -- 1 with key: page #, val: page order
//  1 with key: page #, val: array of rules

//  Go through all the rules, check if the placement of the page is <
//  If any rule fails then you break

// Function to append to the map at a specific key
func appendToMap(k int, v int, customMap map[int][]int) {
	// Check if the key exists in the map
	if _, exists := customMap[k]; exists {
		// If the key exists, append the value to the slice at this key
		customMap[k] = append(customMap[k], v)
	} else {
		// If the key does not exist, create a new slice with the value
		customMap[k] = []int{v}
	}
}

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2024/problem/puzzleInput/day_05_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a slice to store the rules
	ruleMap := make(map[int][]int)
	//var rules []string
	var printJobs []string
	parseRules := true
	answer := 0
	p2answer := 0

	// Use bufio.Scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			// At this point we will scan the paper order
			parseRules = false
			continue
		}
		if parseRules {
			// Create a map with all the rules
			ruleParts := strings.Split(scanner.Text(), "|")
			num1, _ := strconv.Atoi(ruleParts[0])
			num2, _ := strconv.Atoi(ruleParts[1])

			appendToMap(num1, num2, ruleMap)

		} else {
			// Append each line to the job slice
			printJobs = append(printJobs, scanner.Text())
		}
	}

	// For each print job check the order
	for _, job := range printJobs {
		paperOrder := make(map[int]int)
		positionPaperOrder := make(map[int]int)
		newPaperOrder := make(map[int]int)

		// Split the string by commas
		parts := strings.Split(job, ",")

		// Build a map with positions per paper #
		for i, v := range parts {
			// key: paper #; value: position in print-job
			// TODO:  Do we need this err value ?
			num, _ := strconv.Atoi(v)
			paperOrder[num] = i
			positionPaperOrder[i] = num
		}

		// Check all rules against each print job order
		ruleIter := 0
		meetsAllRules := true
		for ruleKey, ruleVal := range ruleMap {
			for i := 0; i < len(ruleVal); i++ {
				n1Value, n1Exists := paperOrder[ruleKey]
				n2Value, n2Exists := paperOrder[ruleVal[i]]

				// This is required because if it doesn't exist in the map then the returned value is 0
				if !n1Exists || !n2Exists {
					// If either value isn't in the paper map then skip this rule
					//continue
				} else if n1Value > n2Value {
					// fail in this case
					meetsAllRules = false
					break
				}
			}

			// If the last rule and we haven't broken yet
			if meetsAllRules && ruleIter == (len(ruleMap)-1) {
				// In this case we should add the middle number
				answer += positionPaperOrder[len(paperOrder)/2]
			} else if !meetsAllRules {
				// If it doesn't meet the rules then stop checking
				break
			}
			ruleIter++
		}

		// If it's out of order we need to fix the order
		if !meetsAllRules {
			//	TODO: Fix the order
			var newOrder []int
			for _, v := range parts {
				num, _ := strconv.Atoi(v)
				// If first or no rules for that entry put at end of slice
				// If rules then take all entries from rules map, find min in placement map.. put the value there
				_, exists := ruleMap[num]
				//log.Printf("Part 1 answer: %d", ruleI)
				//log.Printf("Part 1 answer: %d", ruleV)
				if len(newOrder) == 0 || !exists {
					newOrder = append(newOrder, num)
					newPaperOrder[num] = len(newOrder) - 1
				} else {
					//	If here then there are rules for this number
					insertPosition := len(newOrder)
					for i := 0; i < len(ruleMap[num]); i++ {
						value, exists := newPaperOrder[ruleMap[num][i]]
						if exists {
							if value < insertPosition {
								insertPosition = value
							}
							newPaperOrder[ruleMap[num][i]] += 1
						}
					}
					newOrder = append(newOrder[:insertPosition], append([]int{num}, newOrder[insertPosition:]...)...)
					newPaperOrder[num] = insertPosition
				}
			}
			p2answer += newOrder[len(newOrder)/2]
		}

	}
	log.Printf("Part 1 answer: %d", answer)
	log.Printf("Part 2 answer: %d", p2answer)

}

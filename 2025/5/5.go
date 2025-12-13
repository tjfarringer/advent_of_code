package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Part 1:  How many foods are in the fresh lists?
// Simple map and then loop is too slow.  Maybe store first,end within the map ?
func main() {
	start := time.Now()
	freshFoodP1 := 0
	freshFoodP2 := 0
	// key: first, val: end
	freshFoodMap := make(map[int]int)

	// open file
	file, err := os.Open("5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// skip the blank line
			log.Printf("Empty line -- going to next line: %q", line)
			continue
		}
		parts := strings.Split(line, "-")
		// Loop through lines that aren't a range
		if len(parts) != 2 {
			valKey, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Printf("Error converting line start to a number: %v\n", err)
				continue
			}
			for k, v := range freshFoodMap {
				if k <= valKey && v >= valKey {
					freshFoodP1++
					break
				}
			}
		} else {
			// Create fresh food map
			rangeAdded := false
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Printf("Error converting start to a number: %v\n", err)
				continue
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Printf("Error converting end to a number: %v\n", err)
				continue
			}
			// concat the ranges in the map
			for k, v := range freshFoodMap {
				if k <= start && v >= end {
					// range already covered, don't need to add it
					rangeAdded = true
					continue
				} else if k <= start && v < end && start >= k && start <= v {
					// range partially covered, need to update the end
					rangeAdded = true
					freshFoodMap[k] = end
				} else if k > start && v >= end && end >= k && end <= v {
					// range partially covered, need to update the start
					rangeAdded = true
					freshFoodMap[start] = v
					// remove the old range
					delete(freshFoodMap, k)
				}
			}
			// For P2 -- deduplicate the ranges in the map
			for k, v := range freshFoodMap {
				for k2, v2 := range freshFoodMap {
					if k2 == k && v2 == v {
						// same range
						continue
					}
					if k2 <= k && v2 >= v {
						// remove bc range is covered
						delete(freshFoodMap, k)
					} else if k2 <= k && v2 < v && k >= k2 && k <= v2 {
						// range partially covered, update both
						delete(freshFoodMap, k)
						delete(freshFoodMap, k2)
						freshFoodMap[k2] = v
					} else if k2 > k && v2 >= v && v >= k2 && v <= v2 {
						// range partially covered, need to update the start
						delete(freshFoodMap, k)
						delete(freshFoodMap, k2)
						freshFoodMap[k] = v2
					}

				}
			}
			if !rangeAdded {
				// range not covered, need to add it
				freshFoodMap[start] = end
			}
		}
	}

	// Part 2:  How many foods are in the fresh lists?
	for k, v := range freshFoodMap {
		freshFoodP2 += v - k + 1
	}

	log.Printf("Fresh food -- part 1: %d\n", freshFoodP1)
	log.Printf("Fresh food -- part 2: %d\n", freshFoodP2)
	log.Printf("Execution Time: %s\n", time.Since(start))
}

// check for valid product-ids
// firstid-lastid, nextrange, ...
// invalid ids are any id that are made ONLY of a sequence of digits repeated twice(!!)
// ex: 55, 6464, 123123
// none have leading zeros
// return a sum of all the invalid ids

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	invalidSumP1 := 0
	invalidSumP2 := 0

	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		for _, item := range items {
			parts := strings.Split(strings.TrimSpace(item), "-")
			if len(parts) == 2 {
				startNum, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
				endNum, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err1 != nil || err2 != nil {
					log.Printf("Error converting to a number: %v\n", err1)
					continue
				}

				// The sequence repeats at most twice //
				// So we can simply compare first/second half //
				for num := startNum; num <= endNum; num++ {
					totalNumSum := 0
					numStr := strconv.Itoa(num)
					// Part 1
					// odd length numbers cannot be invalid
					if len(numStr)%2 == 0 {
						mid := len(numStr) / 2
						firstHalf := numStr[:mid]
						secondHalf := numStr[mid:]
						if firstHalf == secondHalf {
							// invalid number found
							invalidSumP1 += num
						}
					}
					// Part 2
					// invalid numbers are made only of some sequence of digits repeated at least twice
					// sum all digits of num
					for i := 0; i < len(numStr); i++ {
						totalNumSum += int(numStr[i] - '0')
					}
					// if subset sum and total sum are the same
					// and the length of the number is divisible by the length of the subset
					subsetSum := 0
					for pos := 0; pos < len(numStr)/2; pos++ {
						subsetSum += int(numStr[pos] - '0')
						if totalNumSum%subsetSum == 0 && len(numStr)%(pos+1) == 0 {
							// check if the number is made of repeated subsets
							subset := numStr[:pos+1]
							numRepeat := len(numStr) / (pos + 1)
							repeatedStr := ""
							for r := 0; r < numRepeat; r++ {
								repeatedStr += subset
							}
							if repeatedStr == numStr {
								// invalid number found
								invalidSumP2 += num
								break
							}
						}
					}
				}
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Sum of invalid IDs -- part 1: %d\n", invalidSumP1)
	fmt.Printf("Sum of invalid IDs -- part 2: %d\n", invalidSumP2)
	fmt.Printf("Execution time: %s\n", elapsed)
}

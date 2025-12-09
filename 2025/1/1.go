// Problem:
// crack a safe
// 0 to 99, click for each number
// L (lower #s) or R (higher #s) means to turn the dial which direction
// since it's a circle you can wrap around
// starts at 90
// Password == # times the dial is left pointing at 0 after any rotation in the sequence

// P2:  Need to count if the dial ever points at 0
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func main() {
	start := time.Now()

	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file line by line
	// Store numbers in the slice
	startValue := 50
	password1 := 0
	password2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		turnNum, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Printf("Error converting to a number: %v\n", err)
			continue
		}

		if direction == "L" {
			newVal := mod(startValue-turnNum, 100)
			if newVal == 0 {
				password1++
			}
			// find full rotations
			password2 += int(math.Floor(float64(turnNum) / 100.0))
			if (newVal > startValue && startValue > 0) || newVal == 0 {
				password2++
			}
			startValue = mod(startValue-turnNum, 100)
		} else {
			newVal := mod(startValue+turnNum, 100)
			if newVal == 0 {
				password1++
			}
			// find full rotations
			password2 += int(math.Floor(float64(turnNum) / 100.0))
			if (newVal < startValue && startValue > 0) || newVal == 0 {
				password2++
			}
			startValue = mod(startValue+turnNum, 100)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Password 1: %d\n", password1)
	fmt.Printf("Password 2: %d\n", password2)
	fmt.Printf("Execution time: %v\n", elapsed)
	return
}

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func charToInt(c byte) int {
	return int(c - '0')
}

func findMaxDigitPos(s string, start, end int) (int, int) {
	maxDigit := 0
	maxPos := start
	for i := start; i < end; i++ {
		if digit := charToInt(s[i]); digit > maxDigit {
			maxDigit = digit
			maxPos = i
		}
	}
	return maxDigit, maxPos
}

func main() {
	start := time.Now()
	batterySumP1 := 0
	batterySumP2 := 0

	file, err := os.Open("3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		batteryBank := scanner.Text()

		// Part 1
		maxTenVal, maxTenPos := findMaxDigitPos(batteryBank, 0, len(batteryBank)-1)
		maxUnitVal, _ := findMaxDigitPos(batteryBank, maxTenPos+1, len(batteryBank))

		batterySumP1 += maxTenVal*10 + maxUnitVal

		// Part 2
		// Note: strings.Builder is more efficient than string concatenation
		// (+=) because it avoids creating new strings on each append.
		var intermediateMaxVal strings.Builder
		// start is the first value the first battery can be at
		// likewise for end
		start := 0
		end := len(batteryBank) - 12
		// loop until we have 12 batteries
		for intermediateMaxVal.Len() < 12 {
			// greedy search
			maxVal, maxPos := findMaxDigitPos(batteryBank, start, end+1)
			// update for next loop
			start = maxPos + 1
			end = len(batteryBank) - (12 - intermediateMaxVal.Len() - 1)
			// record max digit in our string builder
			intermediateMaxVal.WriteString(string(maxVal + '0'))

		}
		// convert to int and add to sum
		val, _ := strconv.Atoi(intermediateMaxVal.String())
		batterySumP2 += val

	}

	log.Printf("Battery Sum -- part 1: %d\n", batterySumP1)
	log.Printf("Battery Sum -- part 2: %d\n", batterySumP2)
	log.Printf("Execution Time: %s\n", time.Since(start))
}

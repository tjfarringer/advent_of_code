package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Data is of many reports, 1 report per line
// list of numbers called levels, separated by spaces
// Goal:  Which reports are safe;  safe if:
// Levels are either all increasing or decreasing
// any two adj levels diff by at least 1 and at most 3
// Return:  how many reports are safe?

func removeAtIndex(i int, slice []int) []int {
	// Check if the index is out of bounds
	if i < 0 || i >= len(slice) {
		//fmt.Println("Index out of bounds")
		// Return the original slice if index is invalid
		return slice
	}

	// Remove the element at position i
	return append(slice[:i], slice[i+1:]...)
}

func part2Fncn(reportInfo [][]int, i int, levelNums []int, reportTransformed map[int]bool, reportNum int, reportStatus map[int]bool) [][]int {
	var vWoLevel []int

	if (!reportTransformed[reportNum]) && (!reportStatus[reportNum]) {
		for i := 1; i < len(levelNums); i++ {
			// This is how stackoverflow recommended making a deep copy in go
			vWoLevel = append([]int(nil), levelNums...)
			vWoLevel = removeAtIndex(i, vWoLevel)
			reportInfo = append(reportInfo, vWoLevel)
		}
		reportTransformed[reportNum] = true
	}
	// If the report is already transformed or marked as safe return reportInfo as is
	return reportInfo
}

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2024/problem/puzzleInput/day_02_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numSafeReports := 0
	part2 := true

	// If part 2 then add failures back into the check queue with the problem floor removed

	reportNum := 0
	// Level information for the report
	//reportInfo := make(map[int][]int)
	var reportInfo [][]int
	// Key to determine if the report has been processed and is safe
	reportStatus := make(map[int]bool)
	// Report transformed already
	reportTransformed := make(map[int]bool)
	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Split(report, " ")

		// Convert the slice of strings to a slice of integers
		// TODO:  Add floor number as the first element
		var levelNums []int
		levelNums = append([]int{reportNum}, levelNums...)
		for _, s := range levels {
			// Convert each string to an integer
			num, err := strconv.Atoi(s)
			if err != nil {
				//fmt.Println("Error converting string to integer:", err)
				continue
			}
			levelNums = append(levelNums, num)
		}

		//reportInfo[reportNum] = levelNums
		reportInfo = append(reportInfo, levelNums)
		reportStatus[reportNum] = false
		reportTransformed[reportNum] = false
		reportNum++
	}

	// For every report
	//for _, levelNums := range reportInfo {
	for len(reportInfo) > 0 {
		reportSafe := true
		rollingLvlDiff := 0
		levelNums := reportInfo[0]
		reportNum = levelNums[0]
		// Has this report been transformed previously at the time of grabbing it?
		//log.Printf("Checking report: - %d.", reportNum)
		// Check each level -- start with 2 because position 0 is the report number
		for i := 2; i < len(levelNums); i++ {
			// diff of 1 is ok; diff of 3 is ok;
			if math.Abs(float64(levelNums[i]-levelNums[i-1])) > 3 || math.Abs(float64(levelNums[i]-levelNums[i-1])) < 1 {
				reportSafe = false
				//log.Printf("Report is not safe.  Floor - %d - isn't the right height diff from prev", i)
				if part2 {
					reportInfo = part2Fncn(reportInfo, i, levelNums, reportTransformed, reportNum, reportStatus)
				}
				break
			} else if (rollingLvlDiff < 0) && (levelNums[i]-levelNums[i-1]) > 0 {
				reportSafe = false
				//log.Printf("Report is not safe.  Floor - %d - isn't always decreasing", i)
				if part2 {
					reportInfo = part2Fncn(reportInfo, i, levelNums, reportTransformed, reportNum, reportStatus)
				}
				break
			} else if (rollingLvlDiff > 0) && (levelNums[i]-levelNums[i-1]) < 0 {
				reportSafe = false
				//log.Printf("Report is not safe.  Floor - %d - isn't always increasing", i)
				if part2 {
					reportInfo = part2Fncn(reportInfo, i, levelNums, reportTransformed, reportNum, reportStatus)
				}
				break
			} else {
				rollingLvlDiff += (levelNums[i] - levelNums[i-1])
			}
		}
		if (!reportStatus[reportNum]) && (reportSafe) {
			//log.Printf("Adding report - %d - as safe", i)
			numSafeReports++
			reportStatus[reportNum] = true
		}
		// Remove entry from the queue
		reportInfo = reportInfo[1:]
	}
	log.Printf("Part 1 (or 2) answer -- num safe reports: %d", numSafeReports)

}

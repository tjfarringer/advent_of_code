package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Machine struct {
	target   []bool
	buttons  [][]int
	joltages []int
}

func parseMachine(line string) Machine {
	// Parse target state: [.##.]
	targetRegex := regexp.MustCompile(`\[([.#]+)\]`)
	targetMatch := targetRegex.FindStringSubmatch(line)
	if len(targetMatch) < 2 {
		log.Fatal("Could not parse target state")
	}

	target := make([]bool, len(targetMatch[1]))
	for i, char := range targetMatch[1] {
		target[i] = char == '#'
	}

	// Parse buttons: (3) (1,3) (2) (2,3) (0,2) (0,1)
	buttonRegex := regexp.MustCompile(`\(([0-9,]+)\)`)
	buttonMatches := buttonRegex.FindAllStringSubmatch(line, -1)

	buttons := make([][]int, len(buttonMatches))
	for i, match := range buttonMatches {
		if len(match) < 2 {
			continue
		}
		indices := strings.Split(match[1], ",")
		buttons[i] = make([]int, len(indices))
		for j, indexStr := range indices {
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				log.Fatal("Could not parse button index:", indexStr)
			}
			buttons[i][j] = index
		}
	}

	// Parse joltages: {61,15,50,14,45,50}
	joltageRegex := regexp.MustCompile(`\{([0-9,]+)\}`)
	joltageMatch := joltageRegex.FindStringSubmatch(line)
	var joltages []int
	if len(joltageMatch) >= 2 {
		joltageStrs := strings.Split(joltageMatch[1], ",")
		joltages = make([]int, len(joltageStrs))
		for i, jStr := range joltageStrs {
			j, err := strconv.Atoi(jStr)
			if err != nil {
				log.Fatal("Could not parse joltage:", jStr)
			}
			joltages[i] = j
		}
	}

	return Machine{target: target, buttons: buttons, joltages: joltages}
}

// Used for Part 1
func findMinButtonPresses(machine Machine) (int, int) {
	numButtons := len(machine.buttons)
	numLights := len(machine.target)
	minPresses := math.MaxInt32
	joltageForMin := 0

	// Try all possible combinations of button presses (0 or 1 per button)
	for combination := 0; combination < (1 << numButtons); combination++ {
		currentState := make([]bool, numLights)
		pressCount := 0
		joltageSum := 0

		for buttonIndex := 0; buttonIndex < numButtons; buttonIndex++ {
			if (combination & (1 << buttonIndex)) != 0 {
				pressCount++
				// Sum joltage for this button (joltages index by button)
				if buttonIndex < len(machine.joltages) {
					joltageSum += machine.joltages[buttonIndex]
				}
				for _, lightIndex := range machine.buttons[buttonIndex] {
					if lightIndex < numLights {
						currentState[lightIndex] = !currentState[lightIndex]
					}
				}
			}
		}

		matches := true
		for i := 0; i < numLights; i++ {
			if currentState[i] != machine.target[i] {
				matches = false
				break
			}
		}

		if matches && pressCount < minPresses {
			minPresses = pressCount
			joltageForMin = joltageSum
		}
	}

	if minPresses == math.MaxInt32 {
		return -1, 0
	}
	return minPresses, joltageForMin
}

// findJoltagePresses solves the linear system using Gaussian elimination with RREF
// Each position i: sum(presses[btn] for btn affecting i) = joltages[i]
// Used for Part 2
func findJoltagePresses(machine Machine) int {
	numButtons := len(machine.buttons)
	numPositions := len(machine.joltages)

	// Build augmented matrix
	// Form is of [A|b] where A[i][j] = 1 if button j affects position i
	matrix := make([][]float64, numPositions)
	for i := 0; i < numPositions; i++ {
		matrix[i] = make([]float64, numButtons+1)
		matrix[i][numButtons] = float64(machine.joltages[i])
	}
	for btnIdx, positions := range machine.buttons {
		for _, pos := range positions {
			if pos < numPositions {
				matrix[pos][btnIdx] = 1
			}
		}
	}

	// Transform matrix to Reduced Row Echelon Form
	// Pivot columns (columns with leading 1s) correspond to determined variables.
	pivotRow := 0
	pivotCols := make([]int, 0)
	pivotColSet := make(map[int]bool)
	for col := 0; col < numButtons && pivotRow < numPositions; col++ {
		maxRow := pivotRow
		for row := pivotRow + 1; row < numPositions; row++ {
			if math.Abs(matrix[row][col]) > math.Abs(matrix[maxRow][col]) {
				maxRow = row
			}
		}
		if math.Abs(matrix[maxRow][col]) < 1e-9 {
			continue
		}

		matrix[pivotRow], matrix[maxRow] = matrix[maxRow], matrix[pivotRow]
		pivotCols = append(pivotCols, col)
		pivotColSet[col] = true

		pivot := matrix[pivotRow][col]
		for k := col; k <= numButtons; k++ {
			matrix[pivotRow][k] /= pivot
		}

		for row := 0; row < numPositions; row++ {
			if row != pivotRow && math.Abs(matrix[row][col]) > 1e-9 {
				factor := matrix[row][col]
				for k := col; k <= numButtons; k++ {
					matrix[row][k] -= factor * matrix[pivotRow][k]
				}
			}
		}
		pivotRow++
	}

	// Identify free variables
	// These do not have leading 1s in their column
	// which means they can take on any value
	freeVars := []int{}
	for col := 0; col < numButtons; col++ {
		if !pivotColSet[col] {
			freeVars = append(freeVars, col)
		}
	}

	// Find minimum total presses across all valid non-negative integer solutions
	minTotal := math.MaxInt32

	// Find minimum solution
	var trySolution func(freeIdx int, freeVals []int)
	trySolution = func(freeIdx int, freeVals []int) {
		if freeIdx == len(freeVars) {
			// Compute pivot variable values from RREF
			solution := make([]float64, numButtons)
			for i, fv := range freeVars {
				solution[fv] = float64(freeVals[i])
			}
			for i := len(pivotCols) - 1; i >= 0; i-- {
				col := pivotCols[i]
				val := matrix[i][numButtons]
				for j := col + 1; j < numButtons; j++ {
					val -= matrix[i][j] * solution[j]
				}
				solution[col] = val
			}

			// Check if valid non-negative integers
			total := 0
			for _, val := range solution {
				rounded := int(math.Round(val))
				if rounded < 0 || math.Abs(val-float64(rounded)) > 1e-6 {
					return
				}
				total += rounded
			}

			if total < minTotal {
				minTotal = total
			}
			return
		}

		// Try values 0 to maxJoltage for this free variable
		maxVal := 0
		for _, j := range machine.joltages {
			if j > maxVal {
				maxVal = j
			}
		}
		for v := 0; v <= maxVal; v++ {
			freeVals[freeIdx] = v
			trySolution(freeIdx+1, freeVals)
		}
	}

	if len(freeVars) == 0 {
		trySolution(0, []int{})
	} else {
		trySolution(0, make([]int, len(freeVars)))
	}

	if minTotal == math.MaxInt32 {
		return -1
	}
	return minTotal
}

func main() {
	start := time.Now()

	file, err := os.Open("10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	minBtnPressPart1 := 0
	minBtnPressPart2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		machine := parseMachine(line)
		minPresses, _ := findMinButtonPresses(machine)
		joltagePresses := findJoltagePresses(machine)

		if minPresses == -1 {
			log.Printf("No solution found for machine: %s", line)
			continue
		}

		log.Printf("Machine: %d presses (part1), %d presses (part2)", minPresses, joltagePresses)
		minBtnPressPart1 += minPresses
		if joltagePresses > 0 {
			minBtnPressPart2 += joltagePresses
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total button presses part 1: %d", minBtnPressPart1)
	log.Printf("Total button presses part 2: %d", minBtnPressPart2)
	log.Printf("Execution Time: %s\n", time.Since(start))
}

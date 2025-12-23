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
	target  []bool
	buttons [][]int
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

	return Machine{target: target, buttons: buttons}
}

func findMinButtonPresses(machine Machine) int {
	numButtons := len(machine.buttons)
	numLights := len(machine.target)
	minPresses := math.MaxInt32

	// Try all possible combinations of button presses
	for combination := 0; combination < (1 << numButtons); combination++ {
		// Start with all lights off
		currentState := make([]bool, numLights)
		pressCount := 0

		// For each button, check if it's pressed in this combination
		for buttonIndex := 0; buttonIndex < numButtons; buttonIndex++ {
			if (combination & (1 << buttonIndex)) != 0 {
				// This button is pressed (odd number of times)
				pressCount++
				// Toggle all lights this button affects
				for _, lightIndex := range machine.buttons[buttonIndex] {
					if lightIndex < numLights {
						currentState[lightIndex] = !currentState[lightIndex]
					}
				}
			}
		}

		// Check if current state matches target
		matches := true
		for i := 0; i < numLights; i++ {
			if currentState[i] != machine.target[i] {
				matches = false
				break
			}
		}

		if matches && pressCount < minPresses {
			minPresses = pressCount
		}
	}

	if minPresses == math.MaxInt32 {
		return -1 // No solution found
	}

	return minPresses
}

func main() {
	start := time.Now()

	file, err := os.Open("10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPresses := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		machine := parseMachine(line)
		minPresses := findMinButtonPresses(machine)

		if minPresses == -1 {
			log.Printf("No solution found for machine: %s", line)
			continue
		}

		log.Printf("Machine requires %d presses", minPresses)
		totalPresses += minPresses
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total button presses required: %d", totalPresses)
	log.Printf("Execution Time: %s\n", time.Since(start))
}

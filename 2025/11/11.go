package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

// Input:  list of devices and their outputs
// Each line gives the name of a device followed by a list of the devices to which its outputs are attached
// Data only ever flows from a device through its outputs; it can't flow backwards

// Start with "you"
// find every path from you to "out"
// return the number of unique paths

type Device struct {
	name    string
	outputs []string
}

type PathState struct {
	node       string
	visitedDac bool
	visitedFft bool
}

func parseDevice(line string) Device {
	var outputs []string
	parts := strings.Fields(strings.ReplaceAll(line, ":", " "))
	name := parts[0]
	for i := 1; i < len(parts); i++ {
		outputs = append(outputs, parts[i])
	}
	return Device{name: name, outputs: outputs}
}

func reverseTopoSort(graph map[string][]string) []string {
	// num dependencies for each device
	inDegree := make(map[string]int)

	for node := range graph {
		inDegree[node] = 0
	}

	// Count in-degrees
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}

	// Find leaf nodes (0-degrees)
	var queue []string
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	var result []string

	// Process nodes with 0 in-degree first
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		// For each neighbor of current node
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

func main() {
	start := time.Now()
	devices := make(map[string][]string)
	revTopologicalOrder := make(map[string][]string)

	file, err := os.Open("11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// build the graph
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		device := parseDevice(line)
		devices[device.name] = device.outputs
		for _, output := range device.outputs {
			revTopologicalOrder[output] = append(revTopologicalOrder[output], device.name)
		}
	}

	// Sort based on how many dependencies each device has
	revTopOrder := reverseTopoSort(revTopologicalOrder)
	numPaths := make(map[string]int)
	for _, device := range revTopOrder {
		if device == "out" {
			numPaths[device] = 1
		} else {
			numPaths[device] = 0
			for _, neighbor := range devices[device] {
				numPaths[device] += numPaths[neighbor]
			}
		}
	}

	// Part 2: paths from "svr" to "out" that visit both "dac" and "fft"
	numPathsWithConstraints := make(map[PathState]int)
	for _, device := range revTopOrder {
		for _, visitedDac := range []bool{false, true} {
			for _, visitedFft := range []bool{false, true} {
				state := PathState{device, visitedDac, visitedFft}
				// if at end, check constraint
				if device == "out" {
					if visitedDac && visitedFft {
						numPathsWithConstraints[state] = 1
					} else {
						numPathsWithConstraints[state] = 0
					}
				} else {
					numPathsWithConstraints[state] = 0

					// Calculate new visited states when moving from this device
					newVisitedDac := visitedDac || (device == "dac")
					newVisitedFft := visitedFft || (device == "fft")

					// Sum paths from all neighbors
					for _, neighbor := range devices[device] {
						neighborState := PathState{neighbor, newVisitedDac, newVisitedFft}
						numPathsWithConstraints[state] += numPathsWithConstraints[neighborState]
					}
				}
			}
		}
	}

	// Get result for "svr" starting with appropriate initial state
	svrVisitedDac := ("svr" == "dac")
	svrVisitedFft := ("svr" == "fft")
	svrInitialState := PathState{"svr", svrVisitedDac, svrVisitedFft}

	log.Printf("Part 1 - Total paths from 'you' to 'out': %d", numPaths["you"])
	log.Printf("Part 2 - Paths from 'svr' to 'out' via dac and fft: %d", numPathsWithConstraints[svrInitialState])
	log.Printf("Execution Time: %s\n", time.Since(start))
}

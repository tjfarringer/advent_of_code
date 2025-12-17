package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// which junction boxes to connect so electricty can go to every box?
// each line is 1 box and it's x,y,z coordinates
// In a loop, connect the boxes that are closest to each other

// Connect the 1,000 pairs that are closest to together
// answer := muliplty together the size of the 3 largest circuits

type Point struct {
	x, y, z int
}

type PairDistance struct {
	junctionBox1, junctionBox2 Point
	distance                   float64
}

func straight_line_distance(p1 Point, p2 Point) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y) + (p1.z-p2.z)*(p1.z-p2.z)))
}

func main() {
	start := time.Now()

	file, err := os.Open("8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read line by line and store the points in a slice
	junctionBoxes := []Point{}
	// circuitSize := make(map[int]int)
	circuitMap := make(map[Point]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		z, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatal(err)
		}
		junctionBoxes = append(junctionBoxes, Point{x, y, z})
	}

	// Connect the boxes that are closest to each other
	// closestDistanceMap := make(map[float64]bool)
	var box1 Point
	var box2 Point
	// Find all distances -- only need to compare once; O(n^2);
	var allPairs []PairDistance
	for iBox := 0; iBox < len(junctionBoxes); iBox++ {
		// assign the index to the circuit map
		circuitMap[junctionBoxes[iBox]] = iBox
		// circuitSize[iBox] = 1
		for jBox := iBox + 1; jBox < len(junctionBoxes); jBox++ {
			// skip if the boxes are the same
			if iBox == jBox {
				continue
			}
			dist := straight_line_distance(junctionBoxes[iBox], junctionBoxes[jBox])
			allPairs = append(allPairs, PairDistance{junctionBoxes[iBox], junctionBoxes[jBox], dist})
		}
	}
	// Sort all pairs by distance
	sort.Slice(allPairs, func(i, j int) bool {
		return allPairs[i].distance < allPairs[j].distance
	})

	connections := 0
	round := 0
	// log.Printf("Total pairs available: %d", len(allPairs))
	// for connections < 1000 && round < len(allPairs) {
	// 	if round >= len(allPairs) {
	// 		log.Printf("Ran out of pairs at round %d", round)
	// 		break
	// 	}
	// 	box1 = allPairs[round].junctionBox1
	// 	box2 = allPairs[round].junctionBox2
	// 	// If the boxes are already in the same circuit, skip
	// 	if circuitMap[box1] == circuitMap[box2] {
	// 		// but increment to next distance
	// 		connections++
	// 		round++
	// 		continue
	// 	} else {
	// 		box1Circuit := circuitMap[box1]
	// 		box2Circuit := circuitMap[box2]

	// 		newCircuit := min(box1Circuit, box2Circuit)

	// 		for k, v := range circuitMap {
	// 			if v == box2Circuit || v == box1Circuit {
	// 				circuitMap[k] = newCircuit
	// 			}
	// 		}

	// 		connections++
	// 		round++
	// 	}
	// }

	// log.Printf("Final: connections=%d, rounds=%d, available_pairs=%d",
	// 	connections, round, len(allPairs))

	// // Find number occurrences in each circuit
	// for _, value := range circuitMap {
	// 	circuitSize[value]++
	// }

	// // Part 1:
	// // Find the largest 3 circuits
	// largestCircuits := []int{}
	// for _, value := range circuitSize {
	// 	largestCircuits = append(largestCircuits, value)
	// }
	// sort.Sort(sort.Reverse(sort.IntSlice(largestCircuits)))
	// productP1 := largestCircuits[0] * largestCircuits[1] * largestCircuits[2]

	// PART 2:
	// Loop until it's all one cluster
	// Multiply the x value of the last two points that need to be connected
	finalClusterNumber := min(circuitMap[allPairs[0].junctionBox1], circuitMap[allPairs[0].junctionBox2])
	finalClusterSize := 1
	part2Answer := 0
	round = 0
	log.Printf("Total pairs available: %d", len(allPairs))
	for finalClusterSize != len(junctionBoxes) && round < len(allPairs) {
		if round >= len(allPairs) {
			log.Printf("Ran out of pairs at round %d", round)
			break
		}
		box1 = allPairs[round].junctionBox1
		box2 = allPairs[round].junctionBox2
		// If the boxes are already in the same circuit, skip
		if circuitMap[box1] == circuitMap[box2] {
			// but increment to next distance
			round++
			continue
		} else {
			box1Circuit := circuitMap[box1]
			box2Circuit := circuitMap[box2]

			newCircuit := min(box1Circuit, box2Circuit)

			for k, v := range circuitMap {
				if v == box2Circuit || v == box1Circuit {
					circuitMap[k] = newCircuit
					if newCircuit == finalClusterNumber {
						finalClusterSize++
					}
				}
			}

			round++
			part2Answer = box1.x * box2.x
		}
	}

	log.Printf("Made %d connections out of %d rounds", connections, round)

	// log.Printf("Product of largest 3 circuits -- part 1: %d\n", productP1)
	log.Printf("Product of last two points -- part 2: %d\n", part2Answer)
	log.Printf("Execution Time: %s\n", time.Since(start))
}

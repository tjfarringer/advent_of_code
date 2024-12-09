package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// The disk map uses a dense format to represent the layout of files and free space on the disk.
//The digits alternate between indicating the length of a file and the length of free space.

//Each file on disk also has an ID number based on the order of the files
//as they appear before they are rearranged, starting with ID 0

// Goal:  file blocks one at a time from the end of the disk to the leftmost free space block

// Final step -- update filesystem checksum
//  Add up the result of muliplying each file-id by block position
//	 If free space, skip it

//  Idea -- for each file, put into a map
//  [0]: [0, 1]
//  [1]:  [3, 4, 5] ...

//  Free space stored in a min heap

//  Loop until min from heap is > total-file-slots

// Function to delete an element from a slice at a given index
func deleteFromSlice(slice []int, index int) []int {
	// Remove the element by slicing the array and joining the parts before and after the index
	return append(slice[:index], slice[index+1:]...)
}

// MinHeap is a type for a slice of integers that implements heap.Interface
type MinHeap []int

// Implementing heap.Interface for MinHeap

// Len is the number of elements in the collection.
func (h MinHeap) Len() int { return len(h) }

// Less reports whether the element with index i should sort before the element with index j.
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap swaps the elements with indexes i and j.
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap.
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the minimum element from the heap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2024/problem/puzzleInput/day_09_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	startTime := time.Now()

	fileLocations := make(map[int][]int)
	// Starting location, # free spots
	freeLocations := make(map[int]int)
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Println("Error reading line:", scanner.Err())
		return
	}
	puzzleInput := scanner.Text()
	parsedPuzzleInput := strings.Split(puzzleInput, "")
	fileNum := 0
	diskOffset := 0
	usedFileBlocks := 0

	// Create a MinHeap and initialize it
	h := &MinHeap{}
	heap.Init(h)

	//  Create data structure for problem
	for index, num := range parsedPuzzleInput {
		// If a file
		if index%2 == 0 {
			var fLocation []int
			rangeInt, _ := strconv.Atoi(num)
			// num represents how many spaces this file exists in
			for i := 0; i < rangeInt; i++ {
				// append all locations of the file into fLocation slice
				fLocation = append(fLocation, (i + diskOffset))
			}
			fileLocations[fileNum] = fLocation
			fileNum++
			// Iterate diskOffset between files
			diskOffset += rangeInt
			usedFileBlocks += rangeInt
		} else {
			//	Free space
			freeSpaceCount, _ := strconv.Atoi(num)
			freeLocations[diskOffset] = freeSpaceCount
			for i := 0; i < freeSpaceCount; i++ {
				heap.Push(h, (i + diskOffset))
			}
			// Update disk offset after parsing the free space
			diskOffset += freeSpaceCount
		}
	}

	//	PART 1:  Main loop
	//  If here then you know a swap needs to be made
	//swapFileNum := fileNum
	//for (*h)[0] < usedFileBlocks {
	//	for i := 0; i < len(fileLocations[swapFileNum]); i++ {
	//		if fileLocations[swapFileNum][i] > (*h)[0] {
	//			heap.Push(h, fileLocations[swapFileNum][i])
	//			minFreeSpace := heap.Pop(h)
	//			intValue, _ := minFreeSpace.(int)
	//			fileLocations[swapFileNum][i] = intValue
	//		}
	//	}
	//	swapFileNum--
	//}

	//  PART 2:  Main loop
	//  Only move an entire file;  only try to move it once; start with largest file id and work to the smallest
	for i := (fileNum - 1); i > 0; i-- {
		// Extract the keys from the map
		var freeLocationsKeys []int
		for k := range freeLocations {
			freeLocationsKeys = append(freeLocationsKeys, k)
		}
		// Sort the keys
		sort.Ints(freeLocationsKeys)

		for freeLocationKeysi, k := range freeLocationsKeys {
			if k > fileLocations[i][0] {
				// if location > file location -- break
				break
			} else if freeLocations[k] < len(fileLocations[i]) {
				// space is too small; keep it moving
				continue
			} else {
				remainder := freeLocations[k] - len(fileLocations[i])
				if remainder > 0 {
					freeLocations[k+freeLocations[k]-remainder] = remainder
				}
				for j := 0; j < len(fileLocations[i]); j++ {
					fileLocations[i][j] = k + j
				}
				delete(freeLocations, k)
				deleteFromSlice(freeLocationsKeys, freeLocationKeysi)
				// File has moved, cannot move again, go to next file
				break
			}
		}
	}

	//	Lastly, find the checksum
	checkSumVal := 0
	for key, values := range fileLocations {
		for _, value := range values {
			checkSumVal += key * value
		}
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("Program took %s to run\n", elapsedTime)
	log.Printf("Answer: %d", checkSumVal)
}

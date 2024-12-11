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

//  Stones are in a straight line; each has a number which can change
//  Stones can split into two, causing all other stones to shift
//  Each blink:
//  If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
//	If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
//		The left half of the digits are engraved on the new left stone, and the right half of the digits are
//		engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
//	If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by
//		2024 is engraved on the new stone.

type stoneStep struct {
	stone string
	round int
}

func main() {
	file, err := os.Open("/Users/talmadge.farringer/Documents/side_projects/advent_of_code/2024/problem/puzzleInput/day_11_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	startTime := time.Now()

	//cache := make(map[string][]string)
	cacheStoneCount := make(map[stoneStep]int)
	var stones []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		for x := 0; x < len(splitLine); x++ {
			stones = append(stones, splitLine[x])
		}
	}

	roundCount := 75
	answer := 0
	for i := 0; i < len(stones); i++ {
		answer += findStoneCount(stones[i], roundCount, cacheStoneCount)
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("Program took %s to run\n", elapsedTime)
	log.Printf("Answer: %d", answer)
}

func findStoneCount(stone string, r int, cache map[stoneStep]int) int {
	var v int

	if r == 0 {
		// No rounds left return itself
		return 1
	}

	//  Already saw this entry
	cacheValue, exists := cache[stoneStep{stone, r}]
	if exists {
		return cacheValue
	}

	if stone == "0" {
		v = findStoneCount("1", r-1, cache)
	} else if len(stone)%2 == 0 {
		value := stone
		left := value[:len(value)/2]
		leftNum, _ := strconv.Atoi(left)
		// TODO:  cut off leading 0's -- so if it's "00" => "0"
		right := value[len(value)/2:]
		rightNum, _ := strconv.Atoi(right)
		v = findStoneCount(strconv.Itoa(leftNum), r-1, cache) + findStoneCount(strconv.Itoa(rightNum), r-1, cache)
	} else {
		num, _ := strconv.Atoi(stone)
		v = findStoneCount(strconv.Itoa(num*2024), r-1, cache)
	}
	cache[stoneStep{stone: stone, round: r}] = v
	return v
}

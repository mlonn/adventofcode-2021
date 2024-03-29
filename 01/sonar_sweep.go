package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"time"
)

// Part1 Part 1 of puzzle
func CountIncreasing(input []int) int {
	increase := 0
	for index := range input {
		if index < 1 {
			continue
		}

		if input[index] > input[index-1] {
			increase++
		}
	}
	return increase
}

// Part2 Part2 of puzzle
func CountIncresingSlidingWindow(input []int) int {
	increase := 0
	for i := range input {
		if i < 3 {
			continue
		}
		if input[i] > input[i-3] {
			increase++
		}
	}
	return increase
}

func main() {
	stringInput := utils.Input(1)
	input := utils.Ints(stringInput)
	start := time.Now()
	fmt.Println("Part 1: ", CountIncreasing(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", CountIncresingSlidingWindow(input), "Time", time.Since(start))
}

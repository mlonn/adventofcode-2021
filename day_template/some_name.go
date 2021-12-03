package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(input string) int {
	v := 0

	return v
}

// Part2 Part 2 of puzzle
func Part2(input string) int {
	v := 0

	return v
}

func main() {
	input := parse(utils.Input(0))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "Time", time.Since(start))
}

func parse(input string) string {
	return input
}

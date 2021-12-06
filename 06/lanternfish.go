package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(input []int) int {
	return solve(input, 80)
}

// Part2 Part 2 of puzzle
func Part2(input []int) int {
	return solve(input, 256)
}

func solve(input []int, iterations int) int {
	var days = [9]int{}
	for _, n := range input {
		days[n]++
	}
	for i := 0; i < iterations; i++ {
		willSpawn := days[0]
		for j := 0; j <= 7; j++ {
			days[j] = days[j+1]
		}
		days[6] += willSpawn
		days[8] = willSpawn
	}
	sum := 0
	for _, day := range days {
		sum += day
	}
	return sum
}

func main() {
	input := parse(utils.Input(6))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "Time", time.Since(start))
}

func parse(input string) []int {
	s := strings.Split(input, ",")
	var numbers []int
	for _, ns := range s {
		n, _ := strconv.Atoi(ns)
		numbers = append(numbers, n)
	}

	return numbers
}

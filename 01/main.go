package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(input string) int {
	s := strings.Split(input, "\n")
	increase := 0
	var numbers []int
	for _, depth := range s {
		i, _ := strconv.Atoi(depth)
		numbers = append(numbers, i)
	}
	for index := range numbers {
		if index < 1 {
			continue
		}

		if  numbers[index] > numbers[index-1] {
			increase++
		}
	}
	return increase
}

// Part2 Part2 of puzzle
func Part2(input string) int {
	s := strings.Split(input, "\n")
	var numbers []int
	for _, depth := range s {
		i, _ := strconv.Atoi(depth)
		numbers = append(numbers, i)
	}
	increase := 0
	for i := range numbers {
		if i < 3 {
			continue
		}
		a := sum(numbers[i-3:i])
		b := sum(numbers[i-2:i+1])
		if b > a {
			increase++
		}
	}
	return increase
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	input := utils.Input(2021, 1)
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "Time", time.Since(start))
}

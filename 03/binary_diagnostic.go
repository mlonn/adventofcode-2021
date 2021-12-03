package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(input []string) int {
	wordLength := len(input[0])
	ones := make([]int, wordLength)
	for _, line := range input {
		for i, bit := range line {
			if bit == '1' {
				ones[i] = ones[i] + 1
			}

		}
	}

	gamma, epsilon := 0, 0
	for i, numberOfOnes := range ones {
		bit := wordLength - i - 1
		if numberOfOnes >= len(input)/2 {
			gamma |= 1 << bit
		} else {
			epsilon |= 1 << bit
		}
	}

	return gamma * epsilon
}

// Part2 Part 2 of puzzle
func Part2(input []string) int {
	o2 := reduce(input, 0, mostCommon)
	co2 := reduce(input, 0, leastCommon)
	return o2 * co2
}

func mostCommon(input []string, bit int) uint8 {
	one, zero := countBits(input, bit)
	if one >= zero {
		return '1'
	} else {
		return '0'
	}
}

func leastCommon(input []string, bit int) uint8 {
	one, zero := countBits(input, bit)
	if zero <= one {
		return '0'
	} else {
		return '1'
	}
}

func countBits(input []string, bit int) (int, int) {
	one, zero := 0, 0
	for _, line := range input {
		if line[bit] == '1' {
			one++
		} else {
			zero++
		}
	}
	return one, zero
}

func reduce(input []string, bit int, correctFunc func(input []string, bit int) uint8) int {
	if len(input) == 1 {
		value, _ := strconv.ParseInt(input[0], 2, 64)
		return int(value)
	}

	correct := correctFunc(input, bit)
	var left []string

	for _, line := range input {
		if line[bit] == correct {
			left = append(left, line)
		}
	}
	return reduce(left, bit+1, correctFunc)
}

func main() {
	input := parse(utils.Input(3))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "Time", time.Since(start))
}

func parse(input string) []string {
	return utils.Strings(input)
}

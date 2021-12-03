package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(input []int) int {
	wordLength := 12
	ones := make([]int, wordLength)
	for _, number := range input {
		for i := 0; i < wordLength; i++ {
			bit := (number >> i) & 1
			if bit == 1 {
				ones[i]++
			}

		}
	}

	gamma, epsilon := 0, 0
	for bit, numberOfOnes := range ones {
		if numberOfOnes >= len(input)/2 {
			gamma |= 1 << bit
		} else {
			epsilon |= 1 << bit
		}
	}

	return gamma * epsilon
}

// Part2 Part 2 of puzzle
func Part2(input []int) int {
	o2 := reduce(input, 11, mostCommon)
	co2 := reduce(input, 11, leastCommon)
	return o2 * co2
}

func mostCommon(input []int, bit int) int {
	one, zero := countBits(input, bit)
	if one >= zero {
		return 1
	} else {
		return 0
	}
}

func leastCommon(input []int, bit int) int {
	one, zero := countBits(input, bit)
	if zero < one {
		return 0
	} else {
		return 1
	}
}

func countBits(input []int, bit int) (int, int) {
	one, zero := 0, 0
	for _, number := range input {
		value := (number >> bit) & 1
		if value == 1 {
			one++
		} else {
			zero++
		}
	}
	return one, zero
}

func reduce(input []int, bit int, correctFunc func(input []int, bit int) int) int {
	if len(input) == 1 {
		return input[0]
	}

	correct := correctFunc(input, bit)

	var left []int
	for _, number := range input {
		value := (number >> bit) & 1
		if value == correct {
			left = append(left, number)
		}
	}
	return reduce(left, bit-1, correctFunc)
}

func main() {
	input := parse(utils.Input(3))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "Time", time.Since(start))
}

func parse(input string) []int {
	return utils.Binary(input)
}

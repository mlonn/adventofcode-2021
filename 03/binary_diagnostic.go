package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(input []string) int64 {

	var sums [12]int
	for _, line := range input {
		for i, bit := range line {
			if bit == '1' {
				sums[i] = sums[i] + 1
			}

		}
	}
	var gamma string
	var epsilon string
	for _, sum := range sums {
		if sum >= len(input)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	return e * g
}

// Part2 Part 2 of puzzle
func Part2(input []string) int64 {
	o2 := o2(input, 0)
	co2 := co2(input, 0)
	return o2 * co2
}

func o2(input []string, bit int) int64 {
	var correct uint8
	one, zero := 0, 0
	for _, line := range input {
		if line[bit] == '1' {
			one++
		} else {
			zero++
		}
	}
	if one >= zero {
		correct = '1'
	} else {
		correct = '0'
	}

	var left []string

	for _, line := range input {
		if line[bit] == correct {
			left = append(left, line)
		}
	}
	if len(left) > 1 {
		return o2(left, bit+1)
	}
	value, _ := strconv.ParseInt(left[0], 2, 64)
	return value
}
func co2(input []string, bit int) int64 {
	var correct uint8
	one, zero := 0, 0
	for _, line := range input {
		if line[bit] == '1' {
			one++
		} else {
			zero++
		}
	}
	if one >= zero {
		correct = '0'
	} else {
		correct = '1'
	}

	var left []string

	for _, line := range input {
		if line[bit] == correct {
			left = append(left, line)
		}
	}
	if len(left) > 1 {
		return co2(left, bit+1)
	}
	value, _ := strconv.ParseInt(left[0], 2, 64)
	return value
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

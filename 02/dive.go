package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Dive Part 1 of puzzle
func Dive(input []string) int {
	h, v := 0, 0
	for _, line := range input {
		s := strings.Split(line, " ")
		amountString, direction := s[1], s[0]
		amount, _ := strconv.Atoi(amountString)
		switch direction {
		case "down":
			v += amount
		case "up":
			v -= amount
		case "forward":
			h += amount
		}
	}
	return h * v
}

// DiveWithAim Part 2 of puzzle
func DiveWithAim(input []string) int {
	h, v, a := 0, 0, 0
	for _, line := range input {
		s := strings.Split(line, " ")
		amountString, direction := s[1], s[0]
		amount, _ := strconv.Atoi(amountString)
		switch direction {
		case "down":
			a += amount
		case "up":
			a -= amount
		case "forward":
			h += amount
			v += a * amount
		}
	}
	return h * v
}

func main() {
	stringInput := utils.Input(2)
	input := utils.Strings(stringInput)
	start := time.Now()
	fmt.Println("Part 1: ", Dive(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", DiveWithAim(input), "Time", time.Since(start))
}

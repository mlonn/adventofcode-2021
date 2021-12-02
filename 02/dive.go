package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Instruction struct {
	Amount    int
	Direction string
}

// Dive Part 1 of puzzle
func Dive(input []Instruction) int {
	h, v := 0, 0
	for _, instruction := range input {
		switch instruction.Direction {
		case "down":
			v += instruction.Amount
		case "up":
			v -= instruction.Amount
		case "forward":
			h += instruction.Amount
		}
	}
	return h * v
}

// DiveWithAim Part 2 of puzzle
func DiveWithAim(input []Instruction) int {
	h, v, a := 0, 0, 0
	for _, instruction := range input {
		switch instruction.Direction {
		case "down":
			a += instruction.Amount
		case "up":
			a -= instruction.Amount
		case "forward":
			h += instruction.Amount
			v += a * instruction.Amount
		}
	}
	return h * v
}

func parse(input string) []Instruction {
	list := utils.Strings(input)
	var instructions []Instruction
	for _, line := range list {
		s := strings.Split(line, " ")
		amount, _ := strconv.Atoi(s[1])
		instructions = append(instructions, Instruction{Amount: amount, Direction: s[0]})

	}
	return instructions
}

func main() {
	input := utils.Input(2)
	instructions := parse(input)
	start := time.Now()
	fmt.Println("Part 1: ", Dive(instructions), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", DiveWithAim(instructions), "Time", time.Since(start))
}

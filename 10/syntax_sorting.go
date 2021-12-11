package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"sort"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(input []string) int {

	points := 0
	for _, line := range input {
		var open string
	loop:
		for _, ch := range line {
			n := len(open) - 1
			switch ch {
			case '(', '{', '[', '<':
				open += string(ch)
			case ')':
				if open[n] != '(' {
					points += 3
					break loop
				} else {
					open = open[:n]
				}
			case ']':
				if open[n] != '[' {
					points += 57
					break loop
				} else {
					open = open[:n]
				}
			case '}':
				if open[n] != '{' {
					points += 1197
					break loop
				} else {
					open = open[:n]
				}
			case '>':
				if open[n] != '<' {
					points += 25137
					break loop
				} else {
					open = open[:n]
				}
			}
		}
	}

	return points
}

// Part2 Part 2 of puzzle
func Part2(input []string) int {
	var incomplete []string
	for _, line := range input {
		var open string
		var inc = true
	loop:
		for _, ch := range line {
			n := len(open) - 1
			switch ch {
			case '(', '{', '[', '<':
				open += string(ch)
			case ')':
				if open[n] != '(' {
					inc = false
					break loop
				} else {
					open = open[:n]
				}
			case ']':
				if open[n] != '[' {
					inc = false
					break loop
				} else {
					open = open[:n]
				}
			case '}':
				if open[n] != '{' {
					inc = false
					break loop
				} else {
					open = open[:n]
				}
			case '>':
				if open[n] != '<' {
					inc = false
					break loop
				} else {
					open = open[:n]
				}
			}
		}
		if inc {
			incomplete = append(incomplete, open)
		}
	}
	var scores []int
	for _, inc := range incomplete {
		sc := 0
		for i := range inc {
			sc *= 5
			switch inc[len(inc)-1-i] {
			case '(':
				sc += 1
			case '{':
				sc += 3
			case '[':
				sc += 2
			case '<':
				sc += 4
			}
		}
		scores = append(scores, sc)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	input := parse(utils.Input(10))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "Time", time.Since(start))
}

func parse(input string) []string {
	return utils.Strings(input)
}

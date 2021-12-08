package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(signals []Signal) int {
	v := 0
	for _, s := range signals {
		for _, n := range s.Output {
			switch len(n) {
			case 2:
				v++
			case 3:
				v++
			case 4:
				v++
			case 7:
				v++
			}
		}
	}
	return v
}

type Signal struct {
	Input  []string
	Output []string
}

// Part2 Part 2 of puzzle
func Part2(signals []Signal) int {
	v := 0

	for _, s := range signals {
		var numbers = map[int]string{}
		var segments = map[string]int{}

		for len(numbers) < 10 {
			for _, n := range s.Input {
				switch len(n) {
				case 2:
					numbers[1] = n
					segments[n] = 1
				case 3:
					numbers[7] = n
					segments[n] = 7
				case 4:
					numbers[4] = n
					segments[n] = 4
				case 5:
					if one, ok := numbers[1]; ok {
						if includeAll(n, one) {
							numbers[3] = n
							segments[n] = 3

						}
					}
					if four, ok := numbers[4]; ok {
						if includeN(n, four, 2) {
							numbers[2] = n
							segments[n] = 2
						}
					}
					three, tok := numbers[3]
					two, sok := numbers[2]
					if tok && sok {
						if n != three && n != two {
							numbers[5] = n
							segments[n] = 5
						}
					}
				case 6:
					if four, ok := numbers[4]; ok {
						if includeAll(n, four) {
							numbers[9] = n
							segments[n] = 9

						}
					}
					if one, ok := numbers[1]; ok {
						if !includeAll(n, one) {
							numbers[6] = n
							segments[n] = 6

						}
					}
					nine, nok := numbers[9]
					six, sok := numbers[6]
					if nok && sok {
						if n != nine && n != six {
							numbers[0] = n
							segments[n] = 0
						}
					}
				case 7:
					numbers[8] = n
					segments[n] = 8
				}
			}
		}
		var output string
		for _, n := range s.Output {
			for _, d := range numbers {
				if hasSame(n, d) {
					output += strconv.Itoa(segments[d])
				}
			}
		}
		i, _ := strconv.Atoi(output)
		v += i
	}

	return v
}
func includeAll(in, pattern string) bool {
	for _, p := range pattern {
		if !strings.Contains(in, string(p)) {
			return false
		}
	}
	return true
}
func hasSame(in, out string) bool {
	if len(in) != len(out) {
		return false
	}
	x := 0
	for _, p := range in {
		if strings.Contains(out, string(p)) {
			x++
		}
	}
	return x == len(in)
}
func includeN(in, pattern string, n int) bool {
	x := 0
	for _, p := range pattern {
		if strings.Contains(in, string(p)) {
			x++
		}
	}
	return n == x
}
func main() {
	input := parse(utils.Input(8))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "Time", time.Since(start))
}

func parse(input string) []Signal {
	lines := utils.Strings(input)
	var signals []Signal
	for _, s := range lines {
		split := strings.Split(s, " | ")
		in, out := split[0], split[1]
		outputs := strings.Split(out, " ")
		inputs := strings.Split(in, " ")
		signals = append(signals, Signal{
			Input:  inputs,
			Output: outputs,
		})
	}
	return signals
}

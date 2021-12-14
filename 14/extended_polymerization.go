package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"math"
	"strings"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(template string, rules map[string]string) int {
	used := make(map[string]int)
	for k := 0; k < 10; k++ {
		next := string(template[0])
		for i := 0; i < len(template)-1; i++ {
			instruction := string(template[i]) + string(template[i+1])
			next += rules[instruction] + string(template[i+1])
		}
		template = next
	}
	for _, t := range template {
		used[string(t)]++
	}

	min, max := math.MaxInt, 0
	for _, i := range used {
		if i < min {
			min = i

		}
		if i > max {
			max = i

		}
	}
	return max - min
}

// Part2 Part 2 of puzzle
func Part2(template string, rules map[string]string) int {
	used := make(map[string]int)
	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		instruction := string(template[i]) + string(template[i+1])
		pairs[instruction]++
		used[string(template[i])]++
	}
	// add the last character since it's not looped over above
	used[string(template[len(template)-1])]++
	for step := 0; step < 40; step++ {
		next := make(map[string]int)
		for k, v := range pairs {
			next[k] = v
		}
		for p, i := range pairs {
			result := rules[p]
			p1 := string(p[0]) + result
			p2 := result + string(p[1])
			used[result] += i
			next[p] -= i
			if next[p] < 1 {
				delete(next, p)
			}
			next[p1] += i
			next[p2] += i

		}
		pairs = next
	}

	min, max := math.MaxInt, 0
	for _, i := range used {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return max - min
}

func main() {
	template, rules := parse(utils.Input(14))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(template, rules), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(template, rules), "Time", time.Since(start))
}

func parse(input string) (string, map[string]string) {
	gropus := utils.Groups(input)
	m := make(map[string]string)
	for _, line := range utils.Strings(gropus[1]) {
		split := strings.Split(line, " -> ")
		m[split[0]] = split[1]
	}
	return gropus[0], m
}

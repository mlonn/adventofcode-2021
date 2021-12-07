package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(input []int) int {
	i := input
	sort.Ints(i)
	median := i[len(i)/2]
	sum := 0
	for _, n := range i {
		a := math.Abs(float64(median - n))
		sum += int(a)
	}

	return sum
}

// Part2 Part 2 of puzzle
func Part2(input []int) int {
	in := input
	sum := float64(0)
	for _, n := range in {
		sum += float64(n)
	}
	mean := sum / float64(len(in))
	f := int(math.Floor(mean))
	c := int(math.Ceil(mean))
	fuel1 := totalCost(in, f)
	fuel2 := totalCost(in, c)
	if fuel2 < fuel1 {
		return fuel2
	}
	return fuel1
}

func totalCost(in []int, p int) int {
	fuel := 0
	for _, n := range in {
		fuel += cost(n, p)
	}
	return fuel
}

func cost(n int, p int) int {
	a := int(math.Abs(float64(n - p)))
	return a * (1 + a) / 2
}

func main() {
	input := parse(utils.Input(7))
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

package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strings"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(input map[string]Cave) int {
	return traverse(input, "start", []string{})
}

// Part2 Part 2 of puzzle
func Part2(input map[string]Cave) int {
	return traverse2(input, "start", []string{}, false)
}

func traverse2(input map[string]Cave, name string, path []string, hasUsed bool) int {
	cave := input[name]

	if name == "end" {
		return 1
	}
	hasBeen := contains(path, name)
	if cave.Size == Small && hasUsed && hasBeen {
		return 0
	}

	newHasUsed := hasUsed
	if cave.Size == Small && hasBeen {
		newHasUsed = true
	}
	p := append(path, name)

	v := 0
	for _, c := range cave.Caves {
		v += traverse2(input, c.Name, p, newHasUsed)
	}
	return v
}

func traverse(input map[string]Cave, name string, path []string) int {
	cave := input[name]
	p := append(path, name)
	if name == "end" {
		return 1
	}

	if cave.Size == Small && contains(path, name) {
		return 0
	}
	v := 0
	for _, c := range cave.Caves {
		v += traverse(input, c.Name, p)
	}
	return v
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	input := parse(utils.Input(12))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "Time", time.Since(start))
}

type Size string

const (
	Large Size = "LARGE" // 0
	Small      = "SMALL"
)

type Cave struct {
	Size
	Caves   []Cave
	Visited bool
	Name    string
}

func parse(input string) map[string]Cave {
	m := make(map[string]Cave)
	for _, row := range utils.Strings(input) {
		sp := strings.Split(row, "-")
		startName := sp[0]
		start := Cave{
			Size:    Small,
			Visited: false,
			Name:    startName,
		}
		if start.Name == strings.ToUpper(start.Name) {
			start.Size = Large
		}
		endName := sp[1]
		end := Cave{
			Size:    Small,
			Visited: false,
			Name:    endName,
		}
		if end.Name == strings.ToUpper(end.Name) {
			end.Size = Large
		}

		if s, ok := m[startName]; ok {
			s.Caves = append(m[startName].Caves, end)
			m[startName] = s
		} else {
			start.Caves = []Cave{end}
			m[startName] = start
		}
		if startName != "start" {

			if s, ok := m[endName]; ok {
				s.Caves = append(m[endName].Caves, start)
				m[endName] = s
			} else {
				end.Caves = []Cave{start}
				m[endName] = end
			}
		}
	}

	return m
}

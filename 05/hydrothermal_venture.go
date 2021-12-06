package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(input []Line) int {
	coords := make(map[Point]int)
	for _, line := range input {
		if line.A.X == line.B.X {
			ay, by := line.A.Y, line.B.Y
			if ay > by {
				ay, by = by, ay
			}
			for i := ay; i <= by; i++ {
				p := Point{X: line.A.X, Y: i}
				coords[p]++
			}
		}
		if line.A.Y == line.B.Y {

			by, bx := line.A.X, line.B.X
			if by > bx {
				by, bx = bx, by
			}
			for i := by; i <= bx; i++ {
				p := Point{X: i, Y: line.A.Y}
				coords[p]++
			}
		}
	}
	totalt := 0
	for _, i := range coords {
		if i > 1 {
			totalt++
		}
	}
	return totalt
}

// Part2 Part 2 of puzzle
func Part2(input []Line) int {
	coords := make(map[Point]int)
	for _, line := range input {

		curX := line.A.X
		curY := line.A.Y
		for (curX != line.B.X || line.A.X == line.B.X) && (curY != line.B.Y || line.A.Y == line.B.Y) {
			p := Point{X: curX, Y: curY}
			coords[p]++
			if curY < line.B.Y {
				curY++
			} else if curY > line.B.Y {
				curY--
			}
			if curX < line.B.X {
				curX++
			} else if curX > line.B.X {
				curX--
			}

		}
		p := Point{X: curX, Y: curY}
		coords[p]++
	}

	total := 0

	for _, i := range coords {
		if i > 1 {
			total++
		}
	}
	return total
}

func main() {
	input := parse(utils.Input(5))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "Time", time.Since(start))
}

type Point struct {
	X int
	Y int
}

type Line struct {
	A Point
	B Point
}

func parse(input string) []Line {
	rows := utils.Strings(input)
	var lines []Line
	for _, row := range rows {
		split := strings.Split(row, " -> ")
		as := strings.Split(split[0], ",")
		bs := strings.Split(split[1], ",")
		ax, _ := strconv.Atoi(as[0])
		ay, _ := strconv.Atoi(as[1])
		bx, _ := strconv.Atoi(bs[0])
		by, _ := strconv.Atoi(bs[1])

		lines = append(lines, Line{
			A: Point{X: ax, Y: ay},
			B: Point{X: bx, Y: by},
		})

	}

	return lines
}

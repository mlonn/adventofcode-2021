package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(p map[point]bool, i []fold) int {

	f := i[0]
	nextp := p
	for pp := range p {
		if f.direction == UP && pp.y > f.value {
			delete(nextp, pp)
			distance := pp.y - f.value
			nextp[point{pp.x, f.value - distance}] = true
		} else if f.direction == LEFT && pp.x > f.value {
			delete(nextp, pp)
			distance := pp.x - f.value
			nextp[point{f.value - distance, pp.y}] = true
		}
	}
	p = nextp

	return len(p)
}

// Part2 Part 2 of puzzle
func Part2(p map[point]bool, i []fold) int {
	v := 0

	for _, f := range i {
		nextp := p
		for pp := range p {
			if f.direction == UP && pp.y > f.value {
				delete(nextp, pp)
				distance := pp.y - f.value
				nextp[point{pp.x, f.value - distance}] = true
			} else if f.direction == LEFT && pp.x > f.value {
				delete(nextp, pp)
				distance := pp.x - f.value
				nextp[point{f.value - distance, pp.y}] = true
			}
		}
		p = nextp
	}
	maxX, maxY := 0, 0

	for pp := range p {
		if pp.x > maxX {
			maxX = pp.x
		}
		if pp.y > maxY {
			maxY = pp.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if p[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

	return v
}

func main() {
	p, i := parse(utils.Input(13))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(p, i), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(p, i), "Time", time.Since(start))
}

type point struct {
	x int
	y int
}

type Direction string

const (
	UP   Direction = "UP"
	LEFT Direction = "LEFT"
)

type fold struct {
	direction Direction
	value     int
}

func parse(input string) (map[point]bool, []fold) {
	split := strings.Split(input, "\n\n")
	c, f := split[0], split[1]
	coordinates := utils.Strings(c)
	foldss := utils.Strings(f)
	points := make(map[point]bool)
	for _, coordinate := range coordinates {
		s := strings.Split(coordinate, ",")
		xs, ys := s[0], s[1]
		x, _ := strconv.Atoi(xs)
		y, _ := strconv.Atoi(ys)
		points[point{x, y}] = true

	}
	var folds []fold
	for _, line := range foldss {
		ff := strings.Split(line, "=")
		value, _ := strconv.Atoi(ff[1])
		if strings.Contains(ff[0], "x") {
			folds = append(folds, fold{
				direction: LEFT,
				value:     value,
			})
		} else {
			folds = append(folds, fold{
				direction: UP,
				value:     value,
			})
		}
	}
	return points, folds
}

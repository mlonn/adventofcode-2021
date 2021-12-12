package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Position struct {
	EnergyLevel int
	Flashed     bool
}

var flashes = 0

type Grid struct {
	Flashes int
	g       [][]Position
	w       int
	h       int
}

// Part1 Part 1 of puzzle
func Part1(input Grid) int {

	for i := 0; i < 100; i++ {
		input.step()
	}

	return flashes
}

// Part2 Part 2 of puzzle
func Part2(input Grid) int {

	var i = 0
	for {
		i++
		if input.step() {
			return i
		}
	}
}

func (grid Grid) flash(x, y int) {
	if grid.g[y][x].Flashed {
		return
	}
	flashes++
	grid.g[y][x].Flashed = true

	grid.increaseNeighbours(x, y)
}
func (grid Grid) step() bool {
	for y, row := range grid.g {
		for x := range row {
			grid.g[y][x].EnergyLevel++
			if grid.g[y][x].EnergyLevel > 9 {
				grid.flash(x, y)
				grid.Flashes++
			}
		}
	}

	var localFlashes = 0
	for y, row := range grid.g {
		for x := range row {
			if grid.g[y][x].Flashed {
				grid.g[y][x].EnergyLevel = 0
				localFlashes++
			}
		}
	}

	for y, row := range grid.g {
		for x := range row {
			if grid.g[y][x].Flashed {
				grid.g[y][x].EnergyLevel = 0
				grid.g[y][x].Flashed = false
			}
		}
	}
	if localFlashes == grid.h*grid.w {
		return true
	}
	return false
}

func (grid Grid) print() {
	for _, row := range grid.g {
		for _, v := range row {
			if v.EnergyLevel < 10 {
				fmt.Print(v.EnergyLevel)
			} else {
				fmt.Print(v.EnergyLevel)
			}

		}
		fmt.Print(" ")
		for _, v := range row {
			if v.Flashed {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
	fmt.Println(&grid.Flashes)
}

func (grid Grid) increaseNeighbours(x, y int) {

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Dont check our self
			if j == 0 && i == 0 {
				continue
			}
			// lower bound
			if x+j == -1 || y+i == -1 {
				continue
			}
			// upper bound
			if x+j >= grid.w || y+i >= grid.h {
				continue
			}
			nextValue := grid.g[y+i][x+j].EnergyLevel + 1
			grid.g[y+i][x+j].EnergyLevel = nextValue
			if nextValue > 9 {
				grid.flash(x+j, y+i)
			}

		}
	}

}

func main() {
	input := parse(utils.Input(11))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input), "Time", time.Since(start))
	input = parse(utils.Input(11))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input), "Time", time.Since(start))
}

func parse(input string) Grid {
	lines := utils.Strings(input)
	h := len(lines)
	w := len(lines[0])

	depth := make([][]Position, len(lines))
	for i, line := range lines {
		ns := strings.Split(line, "")
		depth[i] = make([]Position, len(line))
		for j, n := range ns {
			v, _ := strconv.Atoi(n)
			depth[i][j] = Position{
				EnergyLevel: v,
				Flashed:     false,
			}
		}

	}

	return Grid{
		g: depth,
		w: w,
		h: h,
	}
}

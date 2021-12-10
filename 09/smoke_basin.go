package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Position struct {
	Height  int
	Checked bool
}

type Grid struct {
	g [][]Position
	w int
	h int
}

// Part1 Part 1 of puzzle
func Part1(input Grid) int {
	v := 0

	for y, row := range input.g {
		for x := range row {
			if input.isLowPoint(x, y) {
				v += input.g[y][x].Height + 1
			}
		}
	}

	return v
}

// Part2 Part 2 of puzzle
func Part2(input Grid) int {

	var basins []int
	grid := input
	for y, row := range grid.g {
		for x := range row {
			if grid.isLowPoint(x, y) {
				basins = append(basins, grid.Basin(x, y))
			}
		}
	}
	sort.Ints(basins)

	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func (grid Grid) isLowPoint(x, y int) bool {
	val := grid.g[y][x].Height

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
			if j == 0 || i == 0 {
				if grid.g[y+i][x+j].Height <= val {
					return false
				}
			}
		}
	}
	return true
}
func (grid Grid) Basin(x, y int) int {

	// lower bound
	if x == -1 || y == -1 {
		return 0
	}
	// upper bound
	if x >= grid.w || y >= grid.h {
		return 0
	}
	val := grid.g[y][x]

	if val.Height == 9 {
		return 0
	}
	if val.Checked {
		return 0
	}

	grid.g[y][x].Checked = true
	return grid.Basin(x, y-1) + grid.Basin(x+1, y) + grid.Basin(x, y+1) + grid.Basin(x-1, y) + 1
}

func main() {
	input := parse(utils.Input(9))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input), "Time", time.Since(start))
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
				Height:  v,
				Checked: false,
			}
		}

	}
	return Grid{
		g: depth,
		w: w,
		h: h,
	}
}

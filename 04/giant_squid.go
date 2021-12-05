package main

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Part1 Part 1 of puzzle
func Part1(numbers []int, boards []Board) int {

	for _, number := range numbers {
		for _, board := range boards {
			board.Mark(number)
			if board.IsWinning() {
				return number * board.Score()
			}

		}
	}

	return -1
}

// Part2 Part 2 of puzzle
func Part2(numbers []int, boards []Board) int {
	winning := make(map[int]bool)
	for _, number := range numbers {
		for i, board := range boards {
			board.Mark(number)
			if board.IsWinning() {
				winning[i] = true
			}
			if len(winning) == len(boards) {
				return number * board.Score()

			}
		}
	}

	return -1
}

type Value struct {
	Value  int
	Makred bool
}

type Board [][]Value

func (b Board) Score() int {
	score := 0
	for _, rows := range b {
		for _, value := range rows {
			if !value.Makred {
				score += value.Value
			}

		}
	}
	return score
}

func (b Board) Mark(number int) {
	for i, rows := range b {
		for j, row := range rows {
			if number == row.Value {
				b[i][j].Makred = true
				return
			}
		}
	}
}

func (b Board) IsWinning() bool {
	for i, row := range b {
		rowWin := true
		colWin := true
		for j, value := range row {
			if !value.Makred {
				rowWin = false
			}
			if !b[j][i].Makred {
				colWin = false
			}
		}
		if rowWin {
			return true
		}
		if colWin {
			return true
		}
	}
	return false
}

func main() {
	numbers, boards := parse(utils.Input(04))
	start := time.Now()
	fmt.Println("Part 1: ", Part1(numbers, boards), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(numbers, boards), "Time", time.Since(start))
}

func parse(input string) ([]int, []Board) {
	groups := strings.Split(input, "\n\n")
	var numbers []int
	for _, s := range strings.Split(groups[0], ",") {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	var boards []Board
	for _, b := range groups[1:] {
		rows := strings.Split(b, "\n")
		board := make(Board, len(rows))
		for i, row := range rows {
			row = strings.ReplaceAll(row, "  ", " ")
			row = strings.TrimSpace(row)
			ns := strings.Split(row, " ")
			board[i] = make([]Value, len(ns))
			for j, s := range ns {
				n, _ := strconv.Atoi(s)
				board[i][j] = Value{
					Value:  n,
					Makred: false,
				}
			}
		}
		boards = append(boards, board)

	}

	return numbers, boards
}

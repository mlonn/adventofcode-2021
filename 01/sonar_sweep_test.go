package main

import (
	"advent-of-code-2021/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input := utils.Ints(utils.ReadInput("../inputs/01.txt"))
	solution := 1521
	answer := CountIncreasing(input)
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func TestPart2(t *testing.T) {
	input := utils.Ints(utils.ReadInput("../inputs/01.txt"))
	solution := 1543
	answer := CountIncresingSlidingWindow(input)
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func BenchmarkPart1(b *testing.B) {
	input := utils.Ints(utils.ReadInput("../inputs/01.txt"))
	solution := 1521
	for n := 0; n < b.N; n++ {
		answer := CountIncreasing(input)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	input := utils.Ints(utils.ReadInput("../inputs/01.txt"))
	solution := 1543
	for n := 0; n < b.N; n++ {
		answer := CountIncresingSlidingWindow(input)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

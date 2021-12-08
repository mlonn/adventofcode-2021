package main

import (
	"advent-of-code-2021/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input := utils.ReadInput("../inputs/08.txt")
	solution := 445
	answer := Part1(parse(input))
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func TestPart2(t *testing.T) {
	input := utils.ReadInput("../inputs/08.txt")
	solution := 1043101
	answer := Part2(parse(input))
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func BenchmarkPart1(b *testing.B) {
	input := utils.ReadInput("../inputs/08.txt")
	parsed := parse(input)
	solution := 445
	for n := 0; n < b.N; n++ {
		answer := Part1(parsed)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	input := utils.ReadInput("../inputs/08.txt")
	parsed := parse(input)
	solution := 1043101
	for n := 0; n < b.N; n++ {
		answer := Part2(parsed)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

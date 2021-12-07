package main

import (
	"advent-of-code-2021/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input := utils.ReadInput("../inputs/07.txt")
	solution := 329389
	answer := Part1(parse(input))
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func TestPart2(t *testing.T) {
	input := utils.ReadInput("../inputs/07.txt")
	solution := 86397080
	answer := Part2(parse(input))
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func BenchmarkPart1(b *testing.B) {
	input := utils.ReadInput("../inputs/07.txt")
	parsed := parse(input)
	solution := 329389
	for n := 0; n < b.N; n++ {
		answer := Part1(parsed)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	input := utils.ReadInput("../inputs/07.txt")
	parsed := parse(input)
	solution := 86397080
	for n := 0; n < b.N; n++ {
		answer := Part2(parsed)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

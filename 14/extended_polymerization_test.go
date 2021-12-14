package main

import (
	"advent-of-code-2021/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input := utils.ReadInput("../inputs/14.txt")
	solution := 2003
	answer := Part1(parse(input))
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func TestPart2(t *testing.T) {
	input := utils.ReadInput("../inputs/14.txt")
	solution := 2276644000111
	answer := Part2(parse(input))
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func BenchmarkPart1(b *testing.B) {
	input := utils.ReadInput("../inputs/14.txt")
	template, rules := parse(input)
	solution := 2003
	for n := 0; n < b.N; n++ {
		answer := Part1(template, rules)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	input := utils.ReadInput("../inputs/14.txt")
	template, rules := parse(input)
	solution := 2276644000111
	for n := 0; n < b.N; n++ {
		answer := Part2(template, rules)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

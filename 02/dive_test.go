package main

import (
	"advent-of-code-2021/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input := utils.ReadInput("../inputs/02.txt")
	solution := 1727835
	answer := Dive(parse(input))
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func TestPart2(t *testing.T) {
	input := utils.ReadInput("../inputs/02.txt")
	solution := 1544000595
	answer := DiveWithAim(parse(input))
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func BenchmarkPart1(b *testing.B) {
	input := utils.ReadInput("../inputs/02.txt")
	instructions := parse(input)
	solution := 1727835
	for n := 0; n < b.N; n++ {
		answer := Dive(instructions)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	input := utils.ReadInput("../inputs/02.txt")
	instructions := parse(input)
	solution := 1544000595
	for n := 0; n < b.N; n++ {
		answer := DiveWithAim(instructions)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

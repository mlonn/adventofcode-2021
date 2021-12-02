package main

import (
	"advent-of-code-2021/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input := utils.Strings(utils.ReadInput("../inputs/02.txt"))
	solution := 1727835
	answer := Dive(input)
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func TestPart2(t *testing.T) {
	input := utils.Strings(utils.ReadInput("../inputs/02.txt"))
	solution := 1544000595
	answer := DiveWithAim(input)
	if answer != solution {
		t.Fatalf("answer = %q; solution %q", answer, solution)
	}

}

func BenchmarkPart1(b *testing.B) {
	input := utils.Strings(utils.ReadInput("../inputs/02.txt"))
	solution := 1727835
	for n := 0; n < b.N; n++ {
		answer := Dive(input)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	input := utils.Strings(utils.ReadInput("../inputs/02.txt"))
	solution := 1544000595
	for n := 0; n < b.N; n++ {
		answer := DiveWithAim(input)
		if answer != solution {
			b.Fatalf("answer = %q; solution %q", answer, solution)
		}
	}
}

package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day2
// cpu: Apple M4 Pro
// BenchmarkPuzzle1-12    	      86	  11649330 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	input := read("../../input/day2/puzzle.txt")
	var result int
	for b.Loop() {
		result = puzzle(input, invalid1)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day2
// cpu: Apple M4 Pro
// BenchmarkPuzzle1Generator-12    	  477717	      2494 ns/op	    2808 B/op	     106 allocs/op
func BenchmarkPuzzle1Generator(b *testing.B) {
	input := read("../../input/day2/puzzle.txt")
	var result int
	for b.Loop() {
		result = puzzle_generate(input, generate1)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day2
// cpu: Apple M4 Pro
// BenchmarkPuzzle2-12    	      30	  40568160 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	input := read("../../input/day2/puzzle.txt")
	var result int
	for b.Loop() {
		result = puzzle(input, invalid2)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day2
// cpu: Apple M4 Pro
// BenchmarkPuzzle2Generator-12    	     571	   1898271 ns/op	   45936 B/op	     247 allocs/op
func BenchmarkPuzzle2Generator(b *testing.B) {
	input := read("../../input/day2/puzzle.txt")
	var result int
	for b.Loop() {
		result = puzzle_generate(input, generate2)
	}
	anchor = result
}

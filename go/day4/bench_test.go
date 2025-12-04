package main

import "testing"

var input = read("../../input/day4/puzzle.txt")

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day4
// cpu: Apple M4 Pro
// BenchmarkPuzzle1-12    	   11412	    105891 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for b.Loop() {
		result = puzzle1(input)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day4
// cpu: Apple M4 Pro
// BenchmarkPuzzle2-12    	   24738	     47442 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for b.Loop() {
		result = puzzle2(input)
	}
	anchor = result
}

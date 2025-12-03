package main

import "testing"

var input = read("../../input/day3/puzzle.txt")
var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day3
// cpu: Apple M4 Pro
// BenchmarkPuzzle1-12    	   76458	     13563 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for b.Loop() {
		result = puzzle(input, findLargestJoltage1)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day3
// cpu: Apple M4 Pro
// BenchmarkPuzzle2-12    	   32442	     35439 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for b.Loop() {
		result = puzzle(input, findLargestJoltage2)
	}
	anchor = result
}

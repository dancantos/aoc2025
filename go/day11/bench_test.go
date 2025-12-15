package main

import "testing"

var graph = read("../../input/day11/puzzle.txt")
var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day11
// cpu: Apple M4 Pro
// BenchmarkPuzzle1-12    	   20541	     57215 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for b.Loop() {
		result = puzzle1(graph)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day11
// cpu: Apple M4 Pro
// BenchmarkPuzzle2-12    	       4	 266640969 ns/op	19073392 B/op	  120508 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for b.Loop() {
		result = puzzle2(graph)
	}
	anchor = result
}

package main

import "testing"

var input = read("../../input/day1/puzzle.txt")
var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day1
// cpu: Apple M4 Pro
// BenchmarkPuzzle1-12    	   92149	     11294 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for b.Loop() {
		result = countZero(input)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day1
// cpu: Apple M4 Pro
// BenchmarkPuzzle2-12    	   85070	     11891 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for b.Loop() {
		result = countZeroPasses(input)
	}
	anchor = result
}

package main

import (
	"testing"
)

var ranges, available = read("../../input/day5/puzzle.txt")
var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day5
// cpu: Apple M4 Pro
// BenchmarkPuzzle1-12    	   16161	     73579 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for b.Loop() {
		result = puzzle1(ranges, available)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day5
// cpu: Apple M4 Pro
// BenchmarkPuzzle2-12    	   72616	     15888 ns/op	     192 B/op	       1 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	var cp = make([][2]int, len(ranges))
	for b.Loop() {
		// need a fresh copy as puzzle2 destroys original data
		b.StopTimer()
		for i, r := range ranges {
			cp[i][0], cp[i][1] = r[0], r[1]
		}
		b.StartTimer()
		result = puzzle2(cp)
	}
	anchor = result
}

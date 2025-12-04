package main

import "testing"

var input = read("../../input/day4/puzzle.txt")
var bonus = read("../../input/day4/bonus.txt")

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day4
// cpu: Apple M4 Pro
// BenchmarkPuzzle1/puzzle-12         	   11646	    100788 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPuzzle1/bonus-12          	    5491	    216855 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	b.Run("puzzle", func(b *testing.B) {
		for b.Loop() {
			result = puzzle1(input)
		}
	})
	b.Run("bonus", func(b *testing.B) {
		for b.Loop() {
			result = puzzle1(bonus)
		}
	})
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day4
// cpu: Apple M4 Pro
// BenchmarkPuzzle2/puzzle-12         	     661	   1565763 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPuzzle2/bonus-12          	      31	  37534947 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int

	b.Run("puzzle", func(b *testing.B) {
		for b.Loop() {
			b.StopTimer()
			input := input.Clone()
			b.StartTimer()
			result = puzzle2(input)
		}
	})

	b.Run("bonus", func(b *testing.B) {
		for b.Loop() {
			b.StopTimer()
			bonus := bonus.Clone()
			b.StartTimer()
			result = puzzle2(bonus)
		}
	})
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day4
// cpu: Apple M4 Pro
// BenchmarkPuzzle2Better/puzzle-12         	     648	   1785344 ns/op	  436156 B/op	      46 allocs/op
// BenchmarkPuzzle2Better/bonus-12          	     202	   5860214 ns/op	  127434 B/op	      21 allocs/op
func BenchmarkPuzzle2Better(b *testing.B) {
	var result int

	b.Run("puzzle", func(b *testing.B) {
		for b.Loop() {
			b.StopTimer()
			input := input.Clone()
			b.StartTimer()
			result = puzzle2Better(input)
		}
	})

	b.Run("bonus", func(b *testing.B) {
		for b.Loop() {
			b.StopTimer()
			bonus := bonus.Clone()
			b.StartTimer()
			result = puzzle2Better(bonus)
		}
	})
	anchor = result
}

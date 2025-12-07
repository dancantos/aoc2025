package main

import (
	"bufio"
	"bytes"
	"iter"
	"os"
	"testing"
)

const filename = "../../input/day7/puzzle.txt"

var content []byte = func() []byte {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}()

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day7
// cpu: Apple M4 Pro
// BenchmarkPuzzle1/solve_with_file_io-12         	   12342	     95970 ns/op	   26161 B/op	     156 allocs/op
// BenchmarkPuzzle1/solve_while_reading-12        	   65809	     18363 ns/op	   25880 B/op	     151 allocs/op
// BenchmarkPuzzle1/solve_after_read-12           	  530677	      2235 ns/op	    1336 B/op	       8 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	b.Run("solve_with_file_io", func(b *testing.B) {
		for b.Loop() {
			result = puzzle1(read(filename))
		}
	})

	b.Run("solve_while_reading", func(b *testing.B) {
		for b.Loop() {
			b.StopTimer()
			in := scanner(content)
			b.StartTimer()
			result = puzzle1(in)
		}
	})

	b.Run("solve_after_read", func(b *testing.B) {
		for b.Loop() {
			b.StopTimer()
			in := bufferAll(content)
			b.StartTimer()
			result = puzzle1(in)
		}
	})
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/aoc2025/go/day7
// cpu: Apple M4 Pro
// BenchmarkPuzzle2/solve_with_file_io-12         	   12177	     98405 ns/op	   27369 B/op	     158 allocs/op
// BenchmarkPuzzle2/solve_while_reading-12        	   59980	     20092 ns/op	   27088 B/op	     153 allocs/op
// BenchmarkPuzzle2/solve_after_read-12           	  482778	      2391 ns/op	    2544 B/op	      10 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int

	b.Run("solve_with_file_io", func(b *testing.B) {
		for b.Loop() {
			result = puzzle2(read(filename))
		}
	})

	b.Run("solve_while_reading", func(b *testing.B) {
		for b.Loop() {
			b.StopTimer()
			in := scanner(content)
			b.StartTimer()
			result = puzzle2(in)
		}
	})

	b.Run("solve_after_read", func(b *testing.B) {
		for b.Loop() {
			b.StopTimer()
			in := bufferAll(content)
			b.StartTimer()
			result = puzzle2(in)
		}
	})
	anchor = result
}

func bufferAll(content []byte) iter.Seq[string] {
	s := bufio.NewScanner(bytes.NewReader(content))
	result := []string{}
	for s.Scan() {
		result = append(result, s.Text())
	}
	return func(yield func(string) bool) {
		for _, line := range result {
			if !yield(line) {
				return
			}
		}
	}
}

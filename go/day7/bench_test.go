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
// BenchmarkPuzzle1/solve_with_file_io-12         	    9812	    118598 ns/op	   26097 B/op	     154 allocs/op
// BenchmarkPuzzle1/solve_while_reading-12        	   34195	     36183 ns/op	   25816 B/op	     149 allocs/op
// BenchmarkPuzzle1/solve_after_read-12           	   35368	     30763 ns/op	    1272 B/op	       6 allocs/op
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
// BenchmarkPuzzle2/solve_with_file_io-12         	   10092	    117921 ns/op	   27265 B/op	     155 allocs/op
// BenchmarkPuzzle2/solve_while_reading-12        	   33218	     35798 ns/op	   26984 B/op	     150 allocs/op
// BenchmarkPuzzle2/solve_after_read-12           	   38773	     28726 ns/op	    2440 B/op	       7 allocs/op
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

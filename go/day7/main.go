package main

import (
	"bufio"
	"bytes"
	"fmt"
	"iter"
	"os"
)

func main() {
	r1 := puzzle1(read("input/day7/puzzle.txt"))
	fmt.Println("Puzzle1:", r1)
	r2 := puzzle2(read("input/day7/puzzle.txt"))
	fmt.Println("Puzzle2:", r2)
}

const (
	EMPTY    = '.'
	ACTIVE   = '|'
	SPLITTER = '^'
	START    = 'S'
)

func puzzle1(lines iter.Seq[string]) int {
	// pop the first line, record the starting position
	var prev []rune
	var startPos int
	var l int
	lines(func(line string) bool {
		l = len(line)
		prev = make([]rune, l)
		for i, char := range line {
			if char == START {
				startPos = i
				prev[i] = ACTIVE
				return false
			}
		}
		return false
	})

	// Main solve
	runes := make([]rune, l)
	splitCount := 0
	minPos, maxPos := startPos, startPos // search boundary, increment as we find splitters
	for line := range lines {
		for i := minPos; i <= maxPos; i++ {
			// for i, char := range line {
			char := line[i]
			prevChar := prev[i]
			switch {
			case char == EMPTY && prevChar == ACTIVE:
				runes[i] = ACTIVE
			case char == SPLITTER && prevChar == ACTIVE:
				splitCount++
				runes[i-1] = ACTIVE
				runes[i] = SPLITTER
				runes[i+1] = ACTIVE

				// expand boundary if splits past it
				if i == minPos && minPos > 0 {
					minPos--
				}
				if i == maxPos && maxPos < l {
					maxPos++
				}
			}
		}
		copy(prev, runes)
	}
	return splitCount
}

func puzzle2(lines iter.Seq[string]) int {
	// pop the first line, record the starting position
	var prev []rune
	// tracks number of paths to a position
	var multiplicity []int
	var startPos int
	var l int
	lines(func(line string) bool {
		l = len(line)
		prev = make([]rune, l)
		multiplicity = make([]int, len(prev))
		for i, char := range line {
			if char == START {
				startPos = i
				prev[i] = ACTIVE
				multiplicity[i] = 1
				return false
			}
		}
		return false
	})

	// total path count
	pathCount := 1
	runes := make([]rune, len(prev))
	minPos, maxPos := startPos, startPos // search boundary, increment as we find splitters
	for line := range lines {
		for i := minPos; i <= maxPos; i++ {
			char := line[i]
			prevChar := prev[i]
			switch {
			case char == EMPTY && prevChar == ACTIVE:
				runes[i] = ACTIVE
			case char == SPLITTER && prevChar == ACTIVE:
				pathCount += multiplicity[i] // paths to i split, so we have mult[i]

				// update active and mult states
				runes[i-1] = ACTIVE
				runes[i] = SPLITTER
				runes[i+1] = ACTIVE
				multiplicity[i-1] += multiplicity[i]
				multiplicity[i+1] += multiplicity[i]
				multiplicity[i] = 0 // split, no longer active

				// expand boundary if splits past it
				if i == minPos && minPos > 0 {
					minPos--
				}
				if i == maxPos && maxPos < l {
					maxPos++
				}
			}
		}
		copy(prev, runes)
	}
	return pathCount
}

func scanner(content []byte) iter.Seq[string] {
	s := bufio.NewScanner(bytes.NewReader(content))
	return func(yield func(string) bool) {
		for s.Scan() {
			if !yield(s.Text()) {
				return
			}
		}
	}
}

func read(filename string) iter.Seq[string] {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	return func(yield func(string) bool) {
		for s.Scan() {
			if !yield(s.Text()) {
				return
			}
		}
	}
}

func printLines(lines iter.Seq[string]) {
	for line := range lines {
		fmt.Println(line)
	}
}

// 4509723641302
// 4509723641302

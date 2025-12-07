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
	// pop the first line
	var prev []rune
	lines(func(line string) bool { prev = []rune(line); return false })
	runes := make([]rune, len(prev))
	splitCount := 0
	for line := range lines {
		for j, char := range line {
			prevChar := prev[j]
			switch {
			case char == EMPTY && prevChar == ACTIVE:
				runes[j] = ACTIVE
			case char == SPLITTER && prevChar == ACTIVE:
				splitCount++
				runes[j-1] = ACTIVE
				runes[j] = SPLITTER
				runes[j+1] = ACTIVE
			case char == EMPTY && prevChar == START: // make this condition last, is will only run once at the start - reduce branching
				runes[j] = ACTIVE
			}
		}
		copy(prev, runes)
	}
	return splitCount
}

func puzzle2(lines iter.Seq[string]) int {
	// pop the first line
	var prev []rune
	lines(func(line string) bool { prev = []rune(line); return false })

	// total path count
	pathCount := 1
	// tracks number of paths to a position
	multiplicity := make([]int, len(prev))
	runes := make([]rune, len(prev))
	for line := range lines {
		for j, char := range line {
			prevChar := prev[j]
			switch {
			case char == EMPTY && prevChar == ACTIVE:
				runes[j] = ACTIVE
			case char == SPLITTER && prevChar == ACTIVE:
				pathCount += multiplicity[j] // paths to j split, so we have mult[j]
				// update active and mult states
				runes[j-1] = ACTIVE
				runes[j] = SPLITTER
				runes[j+1] = ACTIVE
				multiplicity[j-1] += multiplicity[j]
				multiplicity[j+1] += multiplicity[j]
				multiplicity[j] = 0 // split, no longer active

			case char == EMPTY && prevChar == START: // make this condition last, is will only run once at the start - reduce branching
				runes[j] = ACTIVE
				multiplicity[j] = 1
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

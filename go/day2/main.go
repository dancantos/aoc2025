package main

import (
	"fmt"
	"iter"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges := read("input/day2/sample.txt")
	fmt.Println("puzzle1          :", puzzle(ranges, invalid1))
	fmt.Println("puzzle1 generator:", puzzle_generate(ranges, generate1))
	fmt.Println()
	fmt.Println("puzzle2          :", puzzle(ranges, invalid2))
	fmt.Println("puzzle2 generator:", puzzle_generate(ranges, generate2))
}

func puzzle(ranges [][2]int, validFn func(int) bool) int {
	sum := 0
	for _, r := range ranges {
		for n := r[0]; n <= r[1]; n++ {
			if validFn(n) {
				sum += n
			}
		}
	}
	return sum
}

func puzzle_generate(ranges [][2]int, generator func(r [2]int) iter.Seq[int]) int {
	sum := 0
	for _, r := range ranges {
		for n := range generator(r) {
			sum += n
		}
	}
	return sum
}

func invalid1(n int) bool {
	digits := _digits(n)
	if digits%2 != 0 {
		return false
	}
	e := int(math.Pow10(digits / 2))
	return n/e == n%e
}

func invalid2(n int) bool {
	digits := _digits(n)
	for d := 1; d <= digits/2; d++ {
		if digits%d != 0 {
			continue
		}
		e := int(math.Pow10(d))
		if n == buildCandidate(n%e, d, digits, e) {
			return true
		}
	}
	return false
}

// generate1 generates invalid candidates for puzzle 1 based on the range
func generate1(r [2]int) iter.Seq[int] {
	lowDigits, highDigits := _digits(r[0]), _digits(r[1])

	return func(yield func(int) bool) {
		for d := lowDigits; d <= highDigits; d++ {
			// ignore digit counts not an even number
			if d%2 != 0 {
				continue
			}

			// compute low and high ranges for candidate halves
			e := int(math.Pow10(d / 2))
			low, high := r[0]/e, r[1]/e

			for n := max(low, 1); n <= min(high, e-1); n++ {
				candidate := n*e + n
				if inRange(r, candidate) {
					if !yield(candidate) {
						return
					}
				}
			}
		}
	}
}

// generate2 generates invalid candidates for puzzle 1 based on the range
func generate2(r [2]int) iter.Seq[int] {
	lowDigits, highDigits := _digits(r[0]), _digits(r[1])
	seen := make(map[int]bool)
	return func(yield func(int) bool) {
		for d := lowDigits; d <= highDigits; d++ {
			for dd := 1; dd <= d/2; dd++ {
				// ignore digit counts that do not divide the target digits
				if d%dd != 0 {
					continue
				}
				e := int(math.Pow10(dd))
				low, high := r[0]/e, r[1]/e
				for n := max(low, 1); n <= high; n++ {
					nn := n
					for nn >= e {
						nn /= e
					}
					candidate := buildCandidate(nn, dd, d, e)
					if !seen[candidate] && inRange(r, candidate) {
						if !yield(candidate) {
							return
						}
						seen[candidate] = true
					}
				}
			}
		}
	}
}

func read(file string) [][2]int {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), ",")
	result := make([][2]int, 0, len(split))
	for _, r := range split {
		lowHigh := strings.Index(r, "-")
		lowS, highS := r[:lowHigh], r[lowHigh+1:]
		low, _ := strconv.Atoi(lowS)
		high, _ := strconv.Atoi(highS)
		result = append(result, [2]int{low, high})
	}
	return result
}

func _digits(n int) int {
	return int(math.Ceil(math.Log10(float64(n))))
}

func buildCandidate(n, digits, totalDigits, e int) int {
	candidate := 0
	exp := 1
	for d := 0; d < totalDigits/digits; d++ {
		candidate += n * exp
		exp *= e
	}
	return candidate
}

func inRange(r [2]int, candidate int) bool {
	return r[0] <= candidate && candidate <= r[1]
}

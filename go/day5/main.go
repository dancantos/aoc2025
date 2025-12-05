package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges, available := read("input/day5/puzzle.txt")
	fmt.Println("Puzzle1:", puzzle1(ranges, available))
	fmt.Println("Puzzle2:", puzzle2(ranges))
}

func puzzle1(ranges [][2]int, available []int) int {
	valid := func(ranges [][2]int, item int) bool {
		for _, r := range ranges {
			if r[0] <= item && item <= r[1] {
				return true
			}
		}
		return false
	}
	count := 0
	for _, item := range available {
		if valid(ranges, item) {
			count++
		}
	}
	return count
}

func puzzle2(ranges [][2]int) int {
	merge := func(r1 *[2]int, r2 [2]int) bool {
		if r1[1] < r2[0] || r1[0] > r2[1] {
			return false
		}
		(*r1)[0] = min((*r1)[0], r2[0])
		(*r1)[1] = max((*r1)[1], r2[1])
		return true
	}

	merged := make([]bool, len(ranges))
	for i := 1; i < len(ranges); i++ {
		for j := 0; j < i; j++ {
			if merged[j] {
				continue
			}
			merged[j] = merge(&ranges[i], ranges[j])
		}
	}

	sum := 0
	for i, r := range ranges {
		if !merged[i] {
			sum += r[1] - r[0] + 1
		}
	}
	return sum
}

func read(filename string) ([][2]int, []int) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	sections := strings.Split(string(contents), "\n\n")
	section1, section2 := sections[0], sections[1]

	ranges := [][2]int{}
	for _, line := range strings.Split(section1, "\n") {
		dash := strings.Index(line, "-")
		i1, _ := strconv.Atoi(line[:dash])
		i2, _ := strconv.Atoi(line[dash+1:])
		ranges = append(ranges, [2]int{i1, i2})
	}

	available := []int{}
	for _, line := range strings.Split(section2, "\n") {
		if line == "" {
			continue
		}
		i, _ := strconv.Atoi(line)
		available = append(available, i)
	}

	return ranges, available
}

package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"

	"github.com/dancantos/aoc2025/go/grid"
)

func main() {
	grid := read("input/day4/bonus.txt")
	fmt.Println("puzzle1:", puzzle1(grid.Clone()))
	fmt.Println("puzzle2:", puzzle2(grid.Clone()))
	fmt.Println("puzzle2:", puzzle2Better(grid.Clone()))
}

func puzzle1(g grid.Grid[int]) int {
	result := 0
	for p, n := range g.All() {
		if n == 0 {
			continue
		}

		count := 0
		for _, n := range g.Neighbors(p.X, p.Y) {
			count += n
		}

		if count < 4 {
			result += 1
		}
	}
	return result
}

func puzzle2(g grid.Grid[int]) int {
	iter := func(g grid.Grid[int]) int {
		remCount := 0
		for p, n := range g.All() {
			if n == 0 {
				continue
			}

			count := 0
			for _, n := range g.Neighbors(p.X, p.Y) {
				count += n
			}

			if count < 4 {
				remCount += 1
				g.Set(p.X, p.Y, 0)
			}
		}
		return remCount
	}

	var remCount, totalRemoved int
	for {
		remCount = iter(g)
		if remCount == 0 {
			break
		}
		totalRemoved += remCount
	}

	return totalRemoved
}

func puzzle2Better(g grid.Grid[int]) int {
	horizon := map[grid.Vec2]struct{}{}
	remCount := 0

	// initial pass
	for p, n := range g.All() {
		if n == 0 {
			continue
		}
		count := 0
		for _, n := range g.Neighbors(p.X, p.Y) {
			count += n
		}

		if count < 4 {
			remCount += 1
			g.Set(p.X, p.Y, 0)
			for p, n := range g.Neighbors(p.X, p.Y) {
				if n != 0 {
					horizon[p] = struct{}{}
				}
			}
		}
	}

	// horizon loop
	var p grid.Vec2
	for len(horizon) > 0 {
		// convoluted way to get single elem from map
		maps.Keys(horizon)(func(extract grid.Vec2) bool {
			p = extract
			return false
		})
		delete(horizon, p)

		if g.Get(p.X, p.Y) == 0 {
			continue
		}

		count := 0
		for _, n := range g.Neighbors(p.X, p.Y) {
			count += n
		}
		if count < 4 {
			remCount += 1
			g.Set(p.X, p.Y, 0)
			for ph, n := range g.Neighbors(p.X, p.Y) {
				if n != 0 {
					horizon[ph] = struct{}{}
				}
			}
		}
	}
	return remCount
}

func read(filename string) grid.Grid[int] {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	result := [][]int{}
	for s.Scan() {
		row := []int{}
		for _, char := range s.Text() {
			if char == '.' {
				row = append(row, 0)
			} else {
				row = append(row, 1)
			}
		}
		result = append(result, row)
	}
	return grid.NewGrid(result)
}

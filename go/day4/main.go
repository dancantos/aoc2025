package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"maps"
	"os"
	"slices"
)

func main() {
	grid := read("input/day4/bonus.txt")
	fmt.Println("puzzle1:", puzzle1(grid.clone()))
	fmt.Println("puzzle2:", puzzle2(grid.clone()))
	fmt.Println("puzzle2:", puzzle2Better(grid.clone()))
}

func puzzle1(grid Grid[int]) int {
	result := 0
	for p, n := range grid.All() {
		if n == 0 {
			continue
		}

		count := 0
		for _, n := range grid.neighbors(p.x, p.y) {
			count += n
		}

		if count < 4 {
			result += 1
		}
	}
	return result
}

func puzzle2(grid Grid[int]) int {
	iter := func(g Grid[int]) int {
		remCount := 0
		for p, n := range g.All() {
			if n == 0 {
				continue
			}

			count := 0
			for _, n := range g.neighbors(p.x, p.y) {
				count += n
			}

			if count < 4 {
				remCount += 1
				g.set(p.x, p.y, 0)
			}
		}
		return remCount
	}

	var remCount, totalRemoved int
	for {
		remCount = iter(grid)
		if remCount == 0 {
			break
		}
		totalRemoved += remCount
	}

	return totalRemoved
}

func puzzle2Better(grid Grid[int]) int {
	horizon := map[Vec2]struct{}{}
	remCount := 0

	// initial pass
	for p, n := range grid.All() {
		if n == 0 {
			continue
		}
		count := 0
		for _, n := range grid.neighbors(p.x, p.y) {
			count += n
		}

		if count < 4 {
			remCount += 1
			grid.set(p.x, p.y, 0)
			for p, n := range grid.neighbors(p.x, p.y) {
				if n != 0 {
					horizon[p] = struct{}{}
				}
			}
		}
	}

	// horizon loop
	var p Vec2
	for len(horizon) > 0 {
		// convoluted way to get single elem from map
		maps.Keys(horizon)(func(extract Vec2) bool {
			p = extract
			return false
		})
		delete(horizon, p)

		if grid.get(p.x, p.y) == 0 {
			continue
		}

		count := 0
		for _, n := range grid.neighbors(p.x, p.y) {
			count += n
		}
		if count < 4 {
			remCount += 1
			grid.set(p.x, p.y, 0)
			for ph, n := range grid.neighbors(p.x, p.y) {
				if n != 0 {
					horizon[ph] = struct{}{}
				}
			}
		}
	}
	return remCount
}

func read(filename string) Grid[int] {
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
	return NewGrid(result)
}

type Vec2 struct{ x, y int }

type Grid[T any] struct {
	grid   [][]T
	bounds Vec2
}

func (g Grid[T]) All() iter.Seq2[Vec2, T] {
	return func(yield func(Vec2, T) bool) {
		for y := 0; y < g.bounds.y; y++ {
			for x := 0; x < g.bounds.x; x++ {
				if !yield(Vec2{x, y}, g.grid[y][x]) {
					return
				}
			}
		}
	}
}

func (g Grid[T]) get(x, y int) T {
	return g.grid[y][x]
}

func (g Grid[T]) set(x, y int, val T) {
	g.grid[y][x] = val
}

func (g Grid[T]) clone() Grid[T] {
	newGrid := make([][]T, 0, g.bounds.y)
	for _, row := range g.grid {
		newGrid = append(newGrid, slices.Clone(row))
	}
	return Grid[T]{newGrid, g.bounds}
}

func copyGrid[T any](dst, src Grid[T]) {
	for i, row := range src.grid {
		copy(dst.grid[i], row)
	}
}

func (g Grid[T]) Print(out io.Writer) {
	for _, row := range g.grid {
		for _, t := range row {
			fmt.Fprint(out, t)
		}
		fmt.Fprint(out, "\n")
	}
}

func (g Grid[T]) neighbors(x, y int) iter.Seq2[Vec2, T] {
	return func(yield func(Vec2, T) bool) {
		if x > 0 {
			if !yield(Vec2{x - 1, y}, g.get(x-1, y)) {
				return
			}
		}
		if x < g.bounds.x-1 {
			if !yield(Vec2{x + 1, y}, g.get(x+1, y)) {
				return
			}
		}
		if y > 0 {
			if !yield(Vec2{x, y - 1}, g.get(x, y-1)) {
				return
			}
		}
		if y < g.bounds.y-1 {
			if !yield(Vec2{x, y + 1}, g.get(x, y+1)) {
				return
			}
		}

		if x > 0 && y > 0 {
			if !yield(Vec2{x - 1, y - 1}, g.get(x-1, y-1)) {
				return
			}
		}
		if x < g.bounds.x-1 && y > 0 {
			if !yield(Vec2{x + 1, y - 1}, g.get(x+1, y-1)) {
				return
			}
		}
		if x > 0 && y < g.bounds.y-1 {
			if !yield(Vec2{x - 1, y + 1}, g.get(x-1, y+1)) {
				return
			}
		}
		if x < g.bounds.x-1 && y < g.bounds.y-1 {
			if !yield(Vec2{x + 1, y + 1}, g.get(x+1, y+1)) {
				return
			}
		}
	}

}

func NewGrid[T any](input [][]T) Grid[T] {
	return Grid[T]{
		grid:   input,
		bounds: Vec2{len(input[0]), len(input)},
	}
}

func NewGridxy[T any](x, y int) Grid[T] {
	grid := make([][]T, 0, y)
	for i := 0; i < y; i++ {
		grid = append(grid, make([]T, x))
	}
	return Grid[T]{grid, Vec2{x, y}}
}

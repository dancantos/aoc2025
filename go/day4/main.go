package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"os"
	"slices"
)

func main() {
	grid := read("input/day4/puzzle.txt")
	fmt.Println("puzzle1:", puzzle1(grid))
	fmt.Println("puzzle2:", puzzle2(grid))
}

func puzzle1(grid Grid[int]) int {
	result := 0
	for x, y := range grid.coords() {
		if grid.get(x, y) == 0 {
			continue
		}

		count := 0
		for n := range grid.neighbors(x, y) {
			count += n
		}

		if count < 4 {
			result += 1
		}
	}
	return result
}

func puzzle2(grid Grid[int]) int {
	// init all memory needed
	auxGrid := grid.clone()
	var tmpGrid Grid[int]

	iter := func(g Grid[int]) (int, Grid[int]) {
		copyGrid(auxGrid, grid)
		remCount := 0
		for x, y := range g.coords() {
			if g.get(x, y) == 0 {
				continue
			}

			count := 0
			for n := range g.neighbors(x, y) {
				count += n
			}

			if count < 4 {
				remCount += 1
				auxGrid.set(x, y, 0)
			}
		}
		return remCount, auxGrid
	}

	var remCount, totalRemoved int
	for {
		remCount, tmpGrid = iter(grid)
		copyGrid(grid, tmpGrid)
		if remCount == 0 {
			break
		}
		totalRemoved += remCount
	}

	return totalRemoved
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

type Grid[T any] struct {
	grid       [][]T
	xmax, ymax int
}

func (g Grid[T]) coords() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for y := 0; y < g.ymax; y++ {
			for x := 0; x < g.xmax; x++ {
				if !yield(x, y) {
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
	newGrid := make([][]T, 0, g.ymax)
	for _, row := range g.grid {
		newGrid = append(newGrid, slices.Clone(row))
	}
	return Grid[T]{newGrid, g.xmax, g.ymax}
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

func (g Grid[T]) neighbors(x, y int) iter.Seq[T] {
	return func(yield func(T) bool) {
		if x > 0 {
			if !yield(g.get(x-1, y)) {
				return
			}
		}
		if x < g.xmax-1 {
			if !yield(g.get(x+1, y)) {
				return
			}
		}
		if y > 0 {
			if !yield(g.get(x, y-1)) {
				return
			}
		}
		if y < g.ymax-1 {
			if !yield(g.get(x, y+1)) {
				return
			}
		}

		if x > 0 && y > 0 {
			if !yield(g.get(x-1, y-1)) {
				return
			}
		}
		if x < g.xmax-1 && y > 0 {
			if !yield(g.get(x+1, y-1)) {
				return
			}
		}
		if x > 0 && y < g.ymax-1 {
			if !yield(g.get(x-1, y+1)) {
				return
			}
		}
		if x < g.xmax-1 && y < g.ymax-1 {
			if !yield(g.get(x+1, y+1)) {
				return
			}
		}
	}

}

func NewGrid[T any](input [][]T) Grid[T] {
	return Grid[T]{
		grid: input,
		xmax: len(input[0]),
		ymax: len(input),
	}
}

func NewGridxy[T any](x, y int) Grid[T] {
	grid := make([][]T, 0, y)
	for i := 0; i < y; i++ {
		grid = append(grid, make([]T, x))
	}
	return Grid[T]{grid, x, y}
}

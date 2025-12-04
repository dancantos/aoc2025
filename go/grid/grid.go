package grid

import (
	"fmt"
	"io"
	"iter"
	"slices"
)

type Vec2 struct{ X, Y int }

type Grid[T any] struct {
	grid   [][]T
	bounds Vec2
}

func (g Grid[T]) All() iter.Seq2[Vec2, T] {
	return func(yield func(Vec2, T) bool) {
		for y := 0; y < g.bounds.Y; y++ {
			for x := 0; x < g.bounds.X; x++ {
				if !yield(Vec2{x, y}, g.grid[y][x]) {
					return
				}
			}
		}
	}
}

func (g Grid[T]) Get(x, y int) T {
	return g.grid[y][x]
}

func (g Grid[T]) Set(x, y int, val T) {
	g.grid[y][x] = val
}

func (g Grid[T]) Clone() Grid[T] {
	newGrid := make([][]T, 0, g.bounds.Y)
	for _, row := range g.grid {
		newGrid = append(newGrid, slices.Clone(row))
	}
	return Grid[T]{newGrid, g.bounds}
}

func CopyGrid[T any](dst, src Grid[T]) {
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

func (g Grid[T]) Neighbors(x, y int) iter.Seq2[Vec2, T] {
	return func(yield func(Vec2, T) bool) {
		if x > 0 {
			if !yield(Vec2{x - 1, y}, g.Get(x-1, y)) {
				return
			}
		}
		if x < g.bounds.X-1 {
			if !yield(Vec2{x + 1, y}, g.Get(x+1, y)) {
				return
			}
		}
		if y > 0 {
			if !yield(Vec2{x, y - 1}, g.Get(x, y-1)) {
				return
			}
		}
		if y < g.bounds.Y-1 {
			if !yield(Vec2{x, y + 1}, g.Get(x, y+1)) {
				return
			}
		}

		if x > 0 && y > 0 {
			if !yield(Vec2{x - 1, y - 1}, g.Get(x-1, y-1)) {
				return
			}
		}
		if x < g.bounds.X-1 && y > 0 {
			if !yield(Vec2{x + 1, y - 1}, g.Get(x+1, y-1)) {
				return
			}
		}
		if x > 0 && y < g.bounds.Y-1 {
			if !yield(Vec2{x - 1, y + 1}, g.Get(x-1, y+1)) {
				return
			}
		}
		if x < g.bounds.X-1 && y < g.bounds.Y-1 {
			if !yield(Vec2{x + 1, y + 1}, g.Get(x+1, y+1)) {
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

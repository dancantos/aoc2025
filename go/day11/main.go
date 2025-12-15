package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"maps"
	"os"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// graph := read("input/day11/sample.txt")
	// graph := read("input/day11/sample2.txt")
	graph := read("input/day11/puzzle.txt")
	// graph.Print(os.Stdout)
	// fmt.Println(puzzle1(graph))
	fmt.Println(puzzle2(graph))
}

func puzzle1(g Graph[string]) int {
	return _dfs(g, "you", "out", NewSet("you"))
}

func _dfs(g Graph[string], src, dst string, seen Set[string]) int {
	// fmt.Println(src)
	count := 0
	for vertex := range g[src].Values() {
		// fmt.Println(vertex)
		if seen.Has(vertex) {
			continue
		}
		if vertex == dst {
			return 1
		}
		seen.Add(vertex)
		count += _dfs(g, vertex, dst, seen)
		seen.Remove(vertex)
	}
	return count
}

func puzzle2(g Graph[string]) int {
	l := len(g) + 1 // make room for the out vertex
	base := mat.NewDense(l, l, nil)

	svrIndex := 0
	dacIndex := 0
	fftIndex := 0
	outIndex := l - 1
	vlist := make([]string, len(g))
	i := 0
	for v := range g.Vertices() {
		vlist[i] = v
		i++
	}
	for i, k := range vlist {
		// Record the indexes that represent our vertices of interest
		switch k {
		case "svr":
			svrIndex = i
		case "dac":
			dacIndex = i
		case "fft":
			fftIndex = i
		}

		// populate the outbound edges
		if g[k].Has("out") {
			base.Set(i, l-1, 1)
		}

		// populate the internal edges
		for j, v := range vlist {
			if g[k].Has(v) {
				base.Set(i, j, 1)
			}
		}
	}

	m := mat.DenseCopyOf(base)

	// store path counters for each section of the journey of interest
	var svr_dac_count float64
	var svr_fft_count float64
	var dac_fft_count float64
	var fft_dac_count float64
	var fft_out_count float64
	var dac_out_count float64
	for range l {
		svr_dac_count += m.At(svrIndex, dacIndex)
		svr_fft_count += m.At(svrIndex, fftIndex)
		dac_fft_count += m.At(dacIndex, fftIndex)
		fft_dac_count += m.At(fftIndex, dacIndex)
		fft_out_count += m.At(fftIndex, outIndex)
		dac_out_count += m.At(dacIndex, outIndex)
		m.Mul(m, base)
	}

	return int((svr_dac_count * dac_fft_count * fft_out_count) + (svr_fft_count * fft_dac_count * dac_out_count))
}

func read(filename string) Graph[string] {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	result := make(Graph[string])
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		k := parts[0]
		k = k[:len(k)-1]
		result[k] = NewSet(parts[1:]...)
	}
	return result
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func (s Set[T]) Has(val T) bool {
	_, has := s[val]
	return has
}

func (s Set[T]) Add(val T) {
	s[val] = struct{}{}
}

func (s Set[T]) Remove(val T) {
	delete(s, val)
}

func (s Set[T]) Values() iter.Seq[T] {
	return maps.Keys(s)
}

func (s Set[T]) String() string {
	b := strings.Builder{}
	for k := range s {
		fmt.Fprintf(&b, "%v, ", k)
	}
	str := b.String()
	return "{" + str[:len(str)-2] + "}"
}

type Graph[V comparable] map[V]Set[V]

func (g Graph[V]) Print(out io.Writer) {
	for k, v := range g {
		fmt.Fprintf(out, "%v: %v\n", k, v)
	}
}

func (g Graph[V]) All() iter.Seq2[V, Set[V]] {
	return maps.All(g)
}

func (g Graph[V]) Vertices() iter.Seq[V] {
	return maps.Keys(g)
}

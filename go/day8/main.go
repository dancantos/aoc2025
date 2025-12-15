package main

import "os"

func main() {
	read("input/day8/sample.txt")
	// read("input/day8/puzzle.txt")
}

func read(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	_ = f
}

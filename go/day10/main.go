package day10

import "os"

func read(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	_ = file
}

func main() {
	read("input/day10/sample.txt")
	// read("input/day10/puzzle.txt")
}

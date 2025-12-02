package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// import sys

// def parse_input(file):
//     sign = lambda char : 1 if char == 'R' else -1
//     split = lambda instruction : (sign(instruction[0]), int(instruction[1:]))
//     with open(file, "r") as f:
//         return [ -1*sign*num for sign, num in [ split(line.strip()) for line in f if len(line) > 0 ] ]

// def count_zero(moves, start=50, count=0):
//     for turn in moves:
//         start = (start+turn) % 100
//         count += start == 0
//     return count

// def count_zero_passes(moves, start=50, count=0):
//     for turn in moves:
//         next = start + turn
//                       # full turns past 100       landed on 0    started from 0 and went negative (overcounted)
//         count +=    abs((next-(next<0))//100)  +  (next == 0)   -   (start == 0 and turn < 0)
//         start = next%100
//     return count

// args = sys.argv[1:]
// if len(args) == 0:
//     file = "input/day1/puzzle.txt"
// else:
//     file = args[0]

// moves = parse_input(file)
// print("part 1:", count_zero(moves))
// print("part 2:", count_zero_passes(moves))

func main() {
	moves := read("input/day1/sample.txt")
	fmt.Println(moves)
	fmt.Println("puzzle1:", countZero(moves))
	fmt.Println("puzzle2:", countZeroPasses(moves))
}

func countZero(moves []int) int {
	count := 0
	current := 50
	for _, turn := range moves {
		current = (current + turn) % 100
		if current == 0 {
			count++
		}
	}
	return count
}

func countZeroPasses(moves []int) int {
	count := 0
	current := 50
	for _, turn := range moves {
		next := current + turn

		switch {
		case next > 0:
			count += next / 100
		case next == 0:
			count++
		case next < 0:
			count += -(next / 100) + 1
		}

		if current == 0 && turn < 0 {
			count--
		}
		current = next % 100
	}
	return count
}

func read(file string) []int {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	result := []int{}
	for s.Scan() {
		instruction := strings.TrimSpace(s.Text())
		if len(instruction) == 0 {
			continue
		}
		dir, amountS := instruction[0], instruction[1:]
		amount, _ := strconv.Atoi(amountS)
		if dir == 'R' {
			result = append(result, amount)
		} else {
			result = append(result, -amount)
		}
	}
	return result
}

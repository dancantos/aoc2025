package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	input := read("input/day3/sample.txt")
	fmt.Println(input)
	fmt.Println(puzzle(input, findLargestJoltage1))
	fmt.Println(puzzle(input, findLargestJoltage2))
}

func puzzle(banks [][]int, fn func([]int) int) int {
	sum := 0
	for _, bank := range banks {
		sum += fn(bank)
	}
	return sum
}

func findLargestJoltage1(bank []int) int {
	high, low := bank[0], bank[1]
	for i := 1; i < len(bank); i++ {
		if i < len(bank)-1 && bank[i] > high {
			high, low = bank[i], bank[i+1]
		} else if bank[i] > low {
			low = bank[i]
		}
	}
	return high*10 + low
}

func findLargestJoltage2(bank []int) int {
	digits := _findLargestJoltage2(bank, 12, make([]int, 0, 12))
	sum := 0
	for pos, n := range digits {
		sum += n * int(math.Pow10(11-pos))
	}
	return sum
}

func _findLargestJoltage2(bank []int, size int, current []int) []int {
	if size == 0 {
		return current
	}
	high, index := bank[0], 0
	for i := 1; i <= len(bank)-size; i++ {
		if bank[i] > high {
			high, index = bank[i], i
		}
	}
	return _findLargestJoltage2(bank[index+1:], size-1, append(current, high))
}

func _max(digits []int) int {
	high := digits[0]
	for _, n := range digits[1:] {
		if n > high {
			high = n
		}
	}
	return high
}

func read(file string) [][]int {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	result := [][]int{}
	for s.Scan() {
		row := []int{}
		for _, char := range s.Text() {
			row = append(row, (int(char - '0')))
		}
		result = append(result, row)
	}
	return result
}

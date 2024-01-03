package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	sequences [][]int
}

func readInput(file string) Puzzle {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sequences := [][]int{}
	for fileScanner.Scan() {
		parts := strings.Fields(fileScanner.Text())
		sequence := []int{}
		for _, v := range parts {
			value, _ := strconv.Atoi(v)
			sequence = append(sequence, value)
		}
		sequences = append(sequences, sequence)
	}
	readFile.Close()

	return Puzzle{
		sequences: sequences,
	}
}

func diff(sequence []int) []int {
	result := []int{}
	for i := 0; i < len(sequence)-1; i++ {
		result = append(result, sequence[i+1]-sequence[i])
	}
	return result
}

func finished(sequence []int) bool {
	for _, v := range sequence {
		if v != 0 {
			return false
		}
	}
	return true
}

func part1(puzzle Puzzle) int {
	result := 0
	for _, sequence := range puzzle.sequences {
		diffs := [][]int{sequence}
		for !finished(diffs[len(diffs)-1]) {
			v := diff(diffs[len(diffs)-1])
			diffs = append(diffs, v)
		}
		for i := len(diffs) - 2; i >= 0; i-- {
			// fmt.Printf("i: %v\n", i)
			// fmt.Printf("diffs[i]: %v\n", diffs[i])
			// fmt.Printf("len(diffs[i])-2: %v\n", len(diffs[i])-2)
			// fmt.Printf("len(diffs[i+1])-1: %v\n\n", len(diffs[i+1])-1)
			left := diffs[i][len(diffs[i])-1]
			bottom := diffs[i+1][len(diffs[i+1])-1]
			// fmt.Printf("left: %v\n", left)
			// fmt.Printf("bottom: %v\n\n", bottom)
			newEl := left + bottom
			diffs[i] = append(diffs[i], newEl)
		}
		result += diffs[0][len(diffs[0])-1]

	}
	return result
}

func part2(puzzle Puzzle) int {

	result := 0
	for _, sequence := range puzzle.sequences {
		diffs := [][]int{sequence}
		for !finished(diffs[len(diffs)-1]) {
			v := diff(diffs[len(diffs)-1])
			diffs = append(diffs, v)
		}
		for i := len(diffs) - 2; i >= 0; i-- {
			// fmt.Printf("i: %v\n", i)
			// fmt.Printf("diffs[i]: %v\n", diffs[i])
			// fmt.Printf("len(diffs[i])-2: %v\n", len(diffs[i])-2)
			// fmt.Printf("len(diffs[i+1])-1: %v\n\n", len(diffs[i+1])-1)
			left := diffs[i][0]
			bottom := diffs[i+1][0]
			// fmt.Printf("left: %v\n", left)
			// fmt.Printf("bottom: %v\n\n", bottom)
			newEl := left - bottom
			diffs[i] = append([]int{newEl}, diffs[i]...)
		}
		result += diffs[0][0]

	}
	return result
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle := readInput(filename)
	return part1(puzzle), part2(puzzle)
}

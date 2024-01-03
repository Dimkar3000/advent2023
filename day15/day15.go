package day15

import (
	"bytes"
	"os"
)

type Puzzle struct {
	instructions [][]byte
}
type Lence struct {
	label     string
	operation byte
	power     int
}

type Puzzle2 struct {
	lences []Lence
}

func calcHash(input []byte) int {
	result := 0
	for _, v := range input {
		result += int(v)
		result *= 17
		result %= 256
	}
	return result
}

func readInput(file string) Puzzle {
	input, _ := os.ReadFile(file)
	instructions := bytes.Split(input, []byte{','})
	return Puzzle{instructions: instructions}
}

func readInput2(file string) Puzzle2 {
	input, _ := os.ReadFile(file)
	instructions := bytes.Split(input, []byte{','})
	lences := []Lence{}
	for _, instruction := range instructions {
		if instruction[len(instruction)-1] == '-' {
			label := string(instruction[:len(instruction)-1])
			lences = append(lences, Lence{label, '-', 0})
			continue
		}
		label := string(instruction[:len(instruction)-2])
		operation := instruction[len(instruction)-2]
		power := int(instruction[len(instruction)-1] - '0')
		lences = append(lences, Lence{label, operation, power})
	}
	return Puzzle2{lences: lences}
}

func part1(puzzle Puzzle) int {
	result := 0
	for _, item := range puzzle.instructions {
		val := calcHash(item)
		result += val
		// fmt.Printf("%v becomes %v\n", string(item), calcHash(item))

	}
	return result
}

func addItem(box []Lence, l Lence) []Lence {
	result := []Lence{}
	contains := false
	for _, lence := range box {
		if lence.label == l.label {
			result = append(result, l)
			contains = true
		} else {
			result = append(result, lence)
		}
	}
	if len(result) == 0 || !contains {
		result = append(result, l)
	}
	return result
}

func removeItem(box []Lence, label string) []Lence {
	result := []Lence{}
	for _, lence := range box {
		if lence.label != label {
			result = append(result, lence)
		}
	}
	return result
}

func part2(puzzle Puzzle2) int {
	boxes := make([][]Lence, 256)
	for _, instruction := range puzzle.lences {
		hash := calcHash([]byte(instruction.label))
		if instruction.operation == '=' {
			boxes[hash] = addItem(boxes[hash], instruction)
		} else {
			boxes[hash] = removeItem(boxes[hash], instruction.label)
		}
	}

	result := 0
	for i, box := range boxes {
		for j, lence := range box {
			// fmt.Printf("label %v, Box %d, Slot %d, Power %d\n", lence.label, (i + 1), (j + 1), lence.power)
			result += (i + 1) * (j + 1) * lence.power
		}
	}
	return result
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle := readInput(filename)
	puzzle2 := readInput2(filename)
	return part1(puzzle), part2(puzzle2)
}

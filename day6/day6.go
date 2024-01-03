package day6

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) Puzzle {
	buf, _ := os.ReadFile(file)
	s := string(buf)
	lines := strings.Split(s, "\n")
	line1 := strings.TrimSpace(strings.Split(lines[0], ":")[1])
	line2 := strings.TrimSpace(strings.Split(lines[1], ":")[1])

	inputs1 := strings.Fields(line1)
	inputs2 := strings.Fields(line2)

	result := Puzzle{}
	for i := 0; i < len(inputs1); i++ {
		time, _ := strconv.Atoi(inputs1[i])
		distance, _ := strconv.Atoi(inputs2[i])
		result.games = append(result.games, [2]int{time, distance})
	}

	return result
}

func readInput2(file string) Puzzle {
	buf, _ := os.ReadFile(file)
	s := string(buf)
	lines := strings.Split(s, "\n")
	line1 := strings.TrimSpace(strings.Split(lines[0], ":")[1])
	line2 := strings.TrimSpace(strings.Split(lines[1], ":")[1])

	line1 = strings.ReplaceAll(line1, " ", "")
	line2 = strings.ReplaceAll(line2, " ", "")

	inputs1 := strings.Fields(line1)
	inputs2 := strings.Fields(line2)

	result := Puzzle{}
	for i := 0; i < len(inputs1); i++ {
		time, _ := strconv.Atoi(inputs1[i])
		distance, _ := strconv.Atoi(inputs2[i])
		result.games = append(result.games, [2]int{time, distance})
	}

	return result
}

type Puzzle struct {
	games [][2]int
}

func isSolution(num int, a, c float64) bool {
	return num*num-int(a)*num+int(c) == 0
}

func part1(inputs Puzzle) int {
	result := 1

	for _, game := range inputs.games {
		a := float64(game[0])
		c := float64(game[1])
		d := math.Sqrt(a*a-4*c) / 2
		lower := int(math.Ceil(a/2 - d))
		higher := int(math.Floor(a/2 + d))

		edges := 0
		if isSolution(lower, a, c) {
			edges += 1
		}
		if isSolution(lower, a, c) {
			edges += 1
		}

		result *= (higher - lower + 1 - edges)
	}

	return result
}

func part2(inputs Puzzle) int {
	result := 1

	for _, game := range inputs.games {
		a := float64(game[0])
		c := float64(game[1])
		d := math.Sqrt(a*a-4*c) / 2
		lower := int(math.Ceil(a/2 - d))
		higher := int(math.Floor(a/2 + d))

		edges := 0
		if isSolution(lower, a, c) {
			edges += 1
		}
		if isSolution(lower, a, c) {
			edges += 1
		}

		result *= (higher - lower + 1 - edges)
	}

	return result
}

func Solve(filename string) (interface{}, interface{}) {
	input1 := readInput(filename)
	input2 := readInput2(filename)
	return part1(input1), part2(input2)
}

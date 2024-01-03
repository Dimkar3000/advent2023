package day11

import (
	"os"
	"strings"
)

type Node struct {
	x, y int64
}
type Puzzle struct {
	nodes []Node
}

func readInput(file string, age int) Puzzle {

	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	nodes := []Node{}
	paddingY := 0
	for y, line := range lines {
		foundGalaxy := false
		paddingX := 0
		for x, v := range line {
			if v == '#' {
				foundGalaxy = true
				nodes = append(nodes, Node{int64(x + paddingX*(age-1)), int64(y + paddingY*(age-1))})
			} else {
				foundGalaxyX := false
				for i := 0; i < len(lines); i++ {
					if lines[i][x] == '#' {
						foundGalaxyX = true
						break
					}
				}
				if !foundGalaxyX {
					paddingX++
				}
			}
		}
		if !foundGalaxy {
			paddingY++
		}
	}
	return Puzzle{nodes: nodes}
}

func part1(puzzle Puzzle) int64 {
	sum := int64(0)
	for i := 0; i < len(puzzle.nodes)-1; i++ {
		for j := i + 1; j < len(puzzle.nodes); j++ {
			xDiff := puzzle.nodes[i].x - puzzle.nodes[j].x
			yDiff := puzzle.nodes[i].y - puzzle.nodes[j].y
			if xDiff < 0 {
				xDiff = -xDiff
			}
			if yDiff < 0 {
				yDiff = -yDiff
			}
			sum += int64(xDiff + yDiff)
		}
	}
	return sum
}

func part2(puzzle Puzzle) int64 {
	sum := int64(0)
	for i := 0; i < len(puzzle.nodes)-1; i++ {
		for j := i + 1; j < len(puzzle.nodes); j++ {
			xDiff := puzzle.nodes[i].x - puzzle.nodes[j].x
			yDiff := puzzle.nodes[i].y - puzzle.nodes[j].y
			if xDiff < 0 {
				xDiff = -xDiff
			}
			if yDiff < 0 {
				yDiff = -yDiff
			}
			sum += int64(xDiff + yDiff)
		}
	}
	return sum
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle1 := readInput(filename, 2)
	puzzle2 := readInput(filename, 1000000)
	return part1(puzzle1), part2(puzzle2)
}

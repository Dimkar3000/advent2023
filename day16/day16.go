package day16

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	North Direction = 0
	South Direction = 1
	West  Direction = 2
	East  Direction = 3
)

type Position struct {
	x, y      int
	direction Direction
}

type Puzzle struct {
	grid [][]byte
}

func readInput(file string) Puzzle {
	input, _ := os.ReadFile(file)

	grid := bytes.Split(input, []byte{'\n'})
	return Puzzle{grid}
}

func keepGoing(position Position, grid [][]byte, keys map[string]int) []Position {
	result := []Position{}
	if position.direction == North {
		result = append(result, Position{position.x + 1, position.y, North})
	} else if position.direction == South {
		result = append(result, Position{position.x - 1, position.y, South})
	} else if position.direction == West {
		result = append(result, Position{position.x, position.y + 1, West})
	} else if position.direction == East {
		result = append(result, Position{position.x, position.y - 1, East})
	}
	return result
}

func newPositions(position Position, grid [][]byte, keys map[string]int) []Position {
	result := []Position{}
	key := fmt.Sprintf("%v %v,%v", position.x, position.y, position.direction)
	if _, ok := keys[key]; ok {
		return result
	} else {
		keys[key] = 1
	}
	char := grid[position.x][position.y]

	// Keep going that way
	if char == '.' ||
		(char == '-' && (position.direction == West || position.direction == East)) ||
		(char == '|' && (position.direction == North || position.direction == South)) {
		result = append(result, keepGoing(position, grid, keys)...)
	} else if char == '-' {
		result = append(result, Position{position.x, position.y + 1, West})
		result = append(result, Position{position.x, position.y - 1, East})
	} else if char == '|' {
		result = append(result, Position{position.x + 1, position.y, North})
		result = append(result, Position{position.x - 1, position.y, South})
	} else if char == '/' {
		if position.direction == South {
			result = append(result, Position{position.x, position.y + 1, West})
		} else if position.direction == East {
			result = append(result, Position{position.x + 1, position.y, North})
		} else if position.direction == West {
			result = append(result, Position{position.x - 1, position.y, South})
		} else if position.direction == North {
			result = append(result, Position{position.x, position.y - 1, East})
		}
	} else if char == '\\' {
		if position.direction == South {
			result = append(result, Position{position.x, position.y - 1, East})
		} else if position.direction == East {
			result = append(result, Position{position.x - 1, position.y, South})
		} else if position.direction == West {
			result = append(result, Position{position.x + 1, position.y, North})
		} else if position.direction == North {
			result = append(result, Position{position.x, position.y + 1, West})
		}
	}
	final := []Position{}
	for _, position := range result {
		if position.x >= 0 && position.x < len(grid) && position.y >= 0 && position.y < len(grid[0]) {
			final = append(final, position)
		}
	}
	return final
}
func calculatePower(startingPos Position, grid [][]byte) int {
	keys := map[string]int{}
	positions := []Position{startingPos}
	for {
		newStart := positions[0]
		positions = positions[1:]
		positions = append(positions, newPositions(newStart, grid, keys)...)
		if len(positions) == 0 {
			break
		}
	}

	uniquePositions := map[string]int{}
	for key := range keys {
		nKey := strings.Split(key, ",")[0]
		uniquePositions[nKey] = 1
	}

	return len(uniquePositions)
}

func part1(puzzle Puzzle) int {
	return calculatePower(Position{0, 0, West}, puzzle.grid)
}

func part2(puzzle Puzzle) int {
	max := 0
	for row := 0; row < len(puzzle.grid); row++ {
		leftMax := calculatePower(Position{row, 0, West}, puzzle.grid)
		rightMax := calculatePower(Position{row, len(puzzle.grid[0]) - 1, East}, puzzle.grid)
		if max < leftMax {
			max = leftMax
		}
		if max < rightMax {
			max = rightMax
		}
	}

	for col := 0; col < len(puzzle.grid[0]); col++ {
		topMax := calculatePower(Position{0, col, North}, puzzle.grid)
		bottomMax := calculatePower(Position{len(puzzle.grid) - 1, col, South}, puzzle.grid)
		if max < topMax {
			max = topMax
		}
		if max < bottomMax {
			max = bottomMax
		}
	}

	return max
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle := readInput(filename)
	return part1(puzzle), part2(puzzle)
}

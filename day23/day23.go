package day23

import (
	"advent2023/common"
	"bytes"
	"fmt"
	"os"
)

type Puzzle struct {
	grid common.Graph
}

func createKey(row, col int) string {
	return fmt.Sprintf("%v,%v", row, col)
}

func readInput(filename string, part2 bool) Puzzle {
	input, _ := os.ReadFile(filename)
	if part2 {
		input = bytes.ReplaceAll(input, []byte(">"), []byte("."))
		input = bytes.ReplaceAll(input, []byte("<"), []byte("."))
		input = bytes.ReplaceAll(input, []byte("^"), []byte("."))
		input = bytes.ReplaceAll(input, []byte("v"), []byte("."))
	}
	lines := bytes.Fields(input)
	counter := 0
	visited := map[string]int{}
	for row, line := range lines {
		for col, char := range line {
			if char != '#' {
				key := createKey(row, col)
				visited[key] = counter

				counter++
			}
		}
	}
	g := common.CreateGraph(counter)
	counter = 0
	for row, line := range lines {
		for col, char := range line {
			if char == '>' {
				key := createKey(row, col+1)
				k, ok := visited[key]
				if ok {
					g.AddEdge(counter, k, 1)
				}
			} else if char == '<' {
				key := createKey(row, col-1)
				k, ok := visited[key]
				if ok {
					g.AddEdge(counter, k, 1)
				}
			} else if char == '^' {
				key := createKey(row-1, col)
				k, ok := visited[key]
				if ok {
					g.AddEdge(counter, k, 1)
				}
			} else if char == 'v' {
				key := createKey(row+1, col)
				k, ok := visited[key]
				if ok {
					g.AddEdge(counter, k, 1)
				}
			} else if char == '.' {
				nKey := createKey(row-1, col)
				sKey := createKey(row+1, col)
				eKey := createKey(row, col+1)
				wKey := createKey(row, col-1)
				n, ok := visited[nKey]
				if ok && lines[row-1][col] != 'v' {
					g.AddEdge(counter, n, 1)
				}
				s, ok := visited[sKey]
				if ok && lines[row+1][col] != '^' {
					g.AddEdge(counter, s, 1)
				}
				w, ok := visited[wKey]
				if ok && lines[row][col-1] != '>' {
					g.AddEdge(counter, w, 1)
				}
				e, ok := visited[eKey]
				if ok && lines[row][col+1] != '<' {
					g.AddEdge(counter, e, 1)
				}
			}

			if char != '#' {
				counter++
			}
		}
	}
	return Puzzle{g}
}

func part1(puzzle Puzzle) int {
	end := puzzle.grid.Size() - 1
	return puzzle.grid.LongestPath(0, end)

}

func Solve(filename string) (interface{}, interface{}) {
	puzzle1 := readInput(filename, false)
	return part1(puzzle1), "(won't do)"
}

package day10

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct{ x, y int }
type Edge struct{ start, end Node }
type Puzzle struct {
	nodes     []Node
	edges     []Edge
	startNode Node
}

var (
	hasUp    = "|LJ"
	hasDown  = "|F7"
	hasLeft  = "-J7"
	hasRight = "-FL"
)

func runeInString(char rune, input string) bool {
	for _, v := range input {
		if v == char {
			return true
		}
	}
	return false
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readInput(input string) Puzzle {
	file, err := readLines(input)

	if err != nil {
		fmt.Println(err)
	}

	nodes := []Node{}
	edges := []Edge{}
	var startNode Node
	for y, line := range file {
		for x, letter := range line {

			if letter == 'S' {
				startNode = Node{x, y}
			} else if letter == '|' {
				nodes = append(nodes, Node{x, y})
				if y > 0 && runeInString(rune(file[y-1][x]), hasDown) {
					edges = append(edges, Edge{Node{x, y - 1}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x, y - 1}})
				}
				if y < len(file)-2 && runeInString(rune(file[y+1][x]), hasUp) {
					edges = append(edges, Edge{Node{x, y + 1}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x, y + 1}})
				}
			} else if letter == '-' {
				nodes = append(nodes, Node{x, y})
				if x > 0 && runeInString(rune(file[y][x-1]), hasRight) {
					edges = append(edges, Edge{Node{x - 1, y}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x - 1, y}})
				}
				if x < len(line)-2 && runeInString(rune(file[y][x+1]), hasLeft) {
					edges = append(edges, Edge{Node{x + 1, y}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x + 1, y}})
				}
			} else if letter == 'L' {
				nodes = append(nodes, Node{x, y})
				if y > 0 && runeInString(rune(file[y][y-1]), hasDown) {
					edges = append(edges, Edge{Node{x, y - 1}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x, y - 1}})
				}
				if x < len(line)-2 && runeInString(rune(file[y][x+1]), hasLeft) {
					edges = append(edges, Edge{Node{x + 1, y}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x + 1, y}})
				}
			} else if letter == 'J' {
				nodes = append(nodes, Node{x, y})
				if y > 0 && runeInString(rune(file[y-1][x]), hasDown) {
					edges = append(edges, Edge{Node{x, y - 1}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x, y - 1}})
				}
				if x > 0 && runeInString(rune(file[y][x-1]), hasRight) {
					edges = append(edges, Edge{Node{x - 1, y}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x - 1, y}})
				}
			} else if letter == '7' {
				nodes = append(nodes, Node{x, y})
				if x > 0 && runeInString(rune(file[y][x-1]), hasRight) {
					edges = append(edges, Edge{Node{x - 1, y}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x - 1, y}})
				}
				if y < len(file)-2 && runeInString(rune(file[y+1][x]), hasUp) {
					edges = append(edges, Edge{Node{x, y + 1}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x, y + 1}})
				}
			} else if letter == 'F' {
				nodes = append(nodes, Node{x, y})
				if x < len(line)-2 && runeInString(rune(file[y][x+1]), hasLeft) {
					edges = append(edges, Edge{Node{x + 1, y}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x + 1, y}})
				}
				if y < len(file)-2 && runeInString(rune(file[y+1][x]), hasUp) {
					edges = append(edges, Edge{Node{x, y + 1}, Node{x, y}})
					edges = append(edges, Edge{Node{x, y}, Node{x, y + 1}})
				}
			}
		}
	}

	return Puzzle{
		nodes:     nodes,
		edges:     edges,
		startNode: startNode,
	}
}

func findNextNode(current, previous, startNode Node, edges []Edge) (Node, bool) {
	for _, edge := range edges {
		if edge.start.x == current.x && edge.start.y == current.y {
			if startNode.x != previous.x && startNode.y != previous.y && edge.end.x == startNode.x && edge.end.y == startNode.y {
				return startNode, true
			}

			if edge.end.x != previous.x || edge.end.y != previous.y {
				return edge.end, false
			}
		}
	}

	return Node{0, 0}, true
}

func createPath(current Node, puzzle *Puzzle) []Node {
	previous := puzzle.startNode
	path := []Node{}
	for {
		v, exit := findNextNode(current, previous, puzzle.startNode, puzzle.edges)
		if exit {
			path = append(path, current)
			break
		}
		previous = current
		current = v
		path = append(path, previous)
		if current.x == puzzle.startNode.x && current.y == puzzle.startNode.y {
			break
		}
	}
	return path
}

func part1(puzzle Puzzle) int {
	currents := []Node{
		{puzzle.startNode.x, puzzle.startNode.y + 1},
		{puzzle.startNode.x, puzzle.startNode.y - 2},
		{puzzle.startNode.x - 1, puzzle.startNode.y},
		{puzzle.startNode.x + 1, puzzle.startNode.y}}
	for _, current := range currents {
		path := createPath(current, &puzzle)
		if len(path) > 1 {
			return len(path)/2 + 1
		}
	}
	return -1
}

func part2(puzzle Puzzle) int {
	currents := []Node{
		{puzzle.startNode.x, puzzle.startNode.y + 1},
		{puzzle.startNode.x, puzzle.startNode.y - 2},
		{puzzle.startNode.x - 1, puzzle.startNode.y},
		{puzzle.startNode.x + 1, puzzle.startNode.y}}
	var loop []Node
	for _, current := range currents {
		path := createPath(current, &puzzle)
		if len(path) > 1 {
			loop = path
			break
		}
	}
	// https://en.wikipedia.org/wiki/Shoelace_formula
	polygonArea := 0
	for i := 0; i < len(loop); i++ {
		cur := loop[i]
		next := loop[(i+1)%len(loop)]

		polygonArea += cur.x*next.y - cur.y*next.x
	}

	if polygonArea < 0 {
		polygonArea = -polygonArea
	}
	polygonArea /= 2

	// https://en.wikipedia.org/wiki/Pick%27s_theorem

	return polygonArea - (len(loop))/2
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle := readInput(filename)
	return part1(puzzle), part2(puzzle)
}

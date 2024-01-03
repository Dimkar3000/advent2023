package day8

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

type Puzzle struct {
	instructions []rune
	nodes        []Node
}

type Node struct {
	name  string
	left  string
	right string
}

func readInput(file string) Puzzle {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	path := []rune(fileScanner.Text())
	fileScanner.Scan()

	nodes := []Node{}
	for fileScanner.Scan() {
		parts := strings.Fields(fileScanner.Text())
		name := parts[0]
		left := parts[2][1:4]
		right := parts[3][:3]
		nodes = append(nodes, Node{
			name:  name,
			left:  left,
			right: right,
		})
	}

	readFile.Close()

	return Puzzle{
		instructions: path,
		nodes:        nodes,
	}
}

func findLenght(start string, endFunc func(current string) bool, instructions []rune, lookup map[string]int, nodes []Node) int {
	counter := 0

	// Create a dictionary for fast lookups
	currentNode := start
	for {
		for _, instr := range instructions {
			counter++
			if instr == 'R' {
				currentNode = nodes[lookup[currentNode]].right
			}
			if instr == 'L' {
				currentNode = nodes[lookup[currentNode]].left
			}
			if endFunc(currentNode) {
				break
			}
		}

		if endFunc(currentNode) {
			break
		}
	}
	return counter
}

func lcm(a, b *big.Int) *big.Int {
	g := new(big.Int).Set(a)
	g.Mul(g, b)
	g.Div(g, new(big.Int).GCD(nil, nil, a, b))
	return g
}

func calculateLCM(numbers []int) *big.Int {
	result := big.NewInt(int64(numbers[0]))

	for i := 1; i < len(numbers); i++ {
		num := big.NewInt(int64(numbers[i]))
		result = lcm(result, num)
	}

	return result
}

func part1(puzzle Puzzle) int {
	lookup := map[string]int{}
	for i, v := range puzzle.nodes {
		lookup[v.name] = i
	}

	return findLenght("AAA", func(current string) bool { return current == "ZZZ" }, puzzle.instructions, lookup, puzzle.nodes)
}

func part2(puzzle Puzzle) int64 {
	// Create a dictionary for fast lookups
	lookup := map[string]int{}
	currentNodes := []string{}
	for i, v := range puzzle.nodes {
		lookup[v.name] = i
		if v.name[2] == 'A' {
			currentNodes = append(currentNodes, v.name)
		}
	}

	lengths := []int{}
	for _, currentNode := range currentNodes {
		lengths = append(lengths, findLenght(
			currentNode,
			func(current string) bool { return current[2] == 'Z' },
			puzzle.instructions,
			lookup,
			puzzle.nodes))
	}
	lcm := calculateLCM(lengths)
	// 	gcd := new(big.Int).GCD(nil, nil, &big.Int(int64(clm)), i)
	// 	lcm = lcm * i / int(gcd.Int64())
	// }

	return lcm.Int64()
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle := readInput(filename)
	return part1(puzzle), part2(puzzle)
}

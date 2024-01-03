package day25

import (
	"bytes"
	"os"
)

// Based On: https://topaz.github.io/paste/#XQAAAQAvAQAAAAAAAAA0m0pnuFI8c82uPD0wiI6r5tRTRja96TZwDRjCPYGCOqExbvZ5+0jns153Ad69VqBjBqQaUspo7NrNRpDc/+ZGFdNSun/wnVOT1qXxnAamzkwKLhxHiCZtWhGd8B/ZVlpxob4CxZmqu+ni/or+caAnmcMG4xTqEH9UT10sSC20Jtq3yetoZEkBm6TqysbKJlV/9noYXwMQIZPn9RpkyjgtIf4yVG1f/rH45+y+/ZBb3VCSDZhg7WGjPNptygpe3fRdyDtq8WnADgtT5icHWOo5yexJky3ebI/tBCnO7+3k4rLE6hL9Eg7kkfrsksYD/3WZAgA=
type Puzzle struct {
	graph map[string][]string
}

func readInput(filename string) Puzzle {
	input, _ := os.ReadFile(filename)
	lines := bytes.Split(input, []byte("\n"))
	graph := map[string][]string{}
	for _, line := range lines {
		parts := bytes.Split(line, []byte(":"))
		label := parts[0]
		fields := bytes.Fields(bytes.TrimSpace(parts[1]))
		for _, field := range fields {
			l := string(label)
			f := string(field)
			graph[l] = append(graph[l], f)
			graph[f] = append(graph[f], l)
		}
	}
	return Puzzle{graph}
}
func subtractSets(a, b []string) []string {
	result := []string{}
	for _, i := range a {
		inside := true
		for _, j := range b {
			if i == j {
				inside = false
				break
			}
		}
		if inside {
			result = append(result, i)
		}
	}
	return result
}

func createSet(graph map[string][]string) []string {
	result := []string{}
	for key := range graph {
		result = append(result, key)
	}
	return result
}

func count(v string, g map[string][]string, s []string) int {
	M := subtractSets(g[v], s)
	return len(M)
}

func apply_count(g map[string][]string, s []string) []int {
	result := []int{}

	for _, v := range s {
		result = append(result, count(v, g, s))
	}
	return result
}

func sum(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func remove_max(g map[string][]string, s []string) []string {
	max := count(s[0], g, s)
	max_key := s[0]
	for _, v := range s {
		nMax := count(v, g, s)
		if nMax > max {
			max = nMax
			max_key = v
		}
	}
	result := []string{}
	for _, v := range s {
		if v != max_key {
			result = append(result, v)
		}
	}
	return result
}

func part(puzzle Puzzle) int {
	Sog := createSet(puzzle.graph)
	S := createSet(puzzle.graph)
	for sum(apply_count(puzzle.graph, S)) != 3 {
		S = remove_max(puzzle.graph, S)
	}
	return len(S) * len(subtractSets(Sog, S))
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle := readInput(filename)
	return part(puzzle), "DONE!!!"
}

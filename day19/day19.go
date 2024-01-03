package day19

import (
	"bytes"
	"os"
	"strconv"
)

type Present struct {
	data map[string]int
}

func apply(present Present, condition Condition) string {
	if condition.operator == "" {
		return condition.result
	}
	if condition.operator == "<" && present.data[condition.label] < condition.limit {
		return condition.result
	}
	if condition.operator == ">" && present.data[condition.label] > condition.limit {
		return condition.result
	}
	return ""
}

func apply_workflow(present Present, workflow Workflow) string {
	for _, condition := range workflow.condtions {
		re := apply(present, condition)
		if re != "" {
			return re
		}
	}

	return ""
}

type Condition struct {
	label    string
	result   string
	operator string
	limit    int
}
type Workflow struct {
	name      string
	condtions []Condition
}

type Puzzle struct {
	workflows map[string]Workflow
	presents  []Present
}

func readCondition(input []byte) Condition {
	result := Condition{}
	if !bytes.ContainsAny(input, "<>") {
		result.operator = ""
		result.result = string(input)
	}
	if bytes.Contains(input, []byte("<")) {
		parts := bytes.Split(input, []byte("<"))
		pp := bytes.Split(parts[1], []byte(":"))
		result.operator = "<"
		result.label = string(parts[0])
		result.result = string(pp[1])
		limit, _ := strconv.Atoi(string(pp[0]))
		result.limit = limit
	}
	if bytes.Contains(input, []byte(">")) {
		parts := bytes.Split(input, []byte(">"))
		pp := bytes.Split(parts[1], []byte(":"))
		result.operator = ">"
		result.label = string(parts[0])
		result.result = string(pp[1])
		limit, _ := strconv.Atoi(string(pp[0]))
		result.limit = limit
	}
	return result
}

func readWorkflow(input []byte) Workflow {
	result := Workflow{}
	parts := bytes.Split(input, []byte("{"))
	result.name = string(parts[0])
	ins := parts[1][:len(parts[1])-1]
	cond_parts := bytes.Split(ins, []byte(","))
	for _, cond := range cond_parts {
		result.condtions = append(result.condtions, readCondition(cond))
	}
	return result
}

func readWorkflows(input []byte) map[string]Workflow {
	result := map[string]Workflow{}
	lines := bytes.Split(input, []byte("\n"))
	for _, line := range lines {
		w := readWorkflow(line)
		result[w.name] = w
	}
	return result
}

func readPresent(input []byte) Present {
	trimmed := input[1 : len(input)-1]
	result := Present{}
	data := map[string]int{}
	parts := bytes.Split(trimmed, []byte(","))
	for _, part := range parts {
		pp := bytes.Split(part, []byte("="))
		v, _ := strconv.Atoi(string(pp[1]))
		data[string(pp[0])] = v
	}
	result.data = data
	return result
}

func readPresents(input []byte) []Present {
	result := []Present{}
	lines := bytes.Split(input, []byte("\n"))
	for _, line := range lines {
		result = append(result, readPresent(line))
	}

	return result
}

func readInput(file string) Puzzle {
	input, _ := os.ReadFile(file)
	parts := bytes.Split(input, []byte("\n\n"))
	return Puzzle{
		workflows: readWorkflows(parts[0]),
		presents:  readPresents(parts[1]),
	}
}

func isAccepted(present Present, workflows map[string]Workflow) bool {
	currentWorkflowName := "in"
	for {
		if currentWorkflowName == "A" || currentWorkflowName == "R" || currentWorkflowName == "" {
			break
		}
		workflow := workflows[currentWorkflowName]
		currentWorkflowName = apply_workflow(present, workflow)
	}
	return currentWorkflowName == "A"
}
func getScore(present Present) int {
	return present.data["x"] + present.data["m"] + present.data["a"] + present.data["s"]
}

func part1(puzzle Puzzle) int {
	sum := 0
	for _, present := range puzzle.presents {
		if isAccepted(present, puzzle.workflows) {
			sum += getScore(present)
		}
	}
	return sum
}

type Range struct {
	low, high int64
}

func (state *State) copy() State {
	newData := map[string]Range{}
	for key, data := range state.data {
		newData[key] = data
	}
	return State{data: newData}
}

func (r *Range) size() int64 {
	return r.high - r.low + 1
}

func (state *State) value() int64 {
	x := state.data["x"]
	m := state.data["m"]
	a := state.data["a"]
	s := state.data["s"]
	return int64(x.size()) * int64(m.size()) * int64(a.size()) * int64(s.size())
}

func apply_cond2(cond Condition, state State) (State, *State) {
	passing_state := state
	not_passing_state := state.copy()
	if cond.operator == "" {
		return passing_state, nil
	}

	ps := passing_state.data[cond.label]
	nps := not_passing_state.data[cond.label]
	if cond.operator == ">" {
		ps.low = int64(cond.limit) + 1
		nps.high = int64(cond.limit)
	} else if cond.operator == "<" {
		ps.high = int64(cond.limit) - 1
		nps.low = int64(cond.limit)
	} else {
		panic("")
	}
	passing_state.data[cond.label] = ps
	not_passing_state.data[cond.label] = nps

	return passing_state, &not_passing_state
}

func apply_workflow2(workflow Workflow, state State) (map[string][]State, int64) {
	result := map[string][]State{}
	sum := int64(0)

	not_passing := state.copy()
	for _, cond := range workflow.condtions {
		r1, r2 := apply_cond2(cond, not_passing)
		if cond.result == "A" {
			sum += r1.value()
		} else if cond.result != "R" {
			result[cond.result] = append(result[cond.result], r1)
		}
		if r2 != nil {
			not_passing = *r2
		}
	}

	return result, sum
}

type State struct {
	data map[string]Range
}

func part2(puzzle Puzzle) int64 {
	state := State{
		data: map[string]Range{
			"x": {1, 4000},
			"m": {1, 4000},
			"a": {1, 4000},
			"s": {1, 4000},
		},
	}
	sum := int64(0)
	workflow := puzzle.workflows["in"]
	r1, r2 := apply_workflow2(workflow, state)
	sum += r2
	for {
		if len(r1) == 0 {
			break
		}
		nr1 := map[string][]State{}
		for wName, states := range r1 {
			workflow = puzzle.workflows[wName]
			for _, state := range states {
				r11, r22 := apply_workflow2(workflow, state)
				sum += r22
				for key, value := range r11 {
					nr1[key] = append(nr1[key], value...)
				}
			}
		}
		r1 = nr1
	}

	return sum
}

func Solve(path string) (interface{}, interface{}) {
	puzzle := readInput(path)
	return part1(puzzle), part2(puzzle)
}

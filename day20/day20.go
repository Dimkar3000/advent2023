package day20

import (
	"bytes"
	"os"
)

type Event struct {
	destination string
	source      string
	state       bool // high = true
}

type Component struct {
	name    string
	ctype   string
	inputs  []string
	outputs []string
	state   []bool
}

var (
	flipflop    = "flipflop"
	conjunction = "conjunction"
	undefined   = "undefined"
)

func createComponent(input []byte) Component {
	cType := undefined
	state := []bool{}
	if input[0] == '%' {
		cType = flipflop
		input = input[1:]
		state = append(state, false)
	} else if input[0] == '&' {
		cType = conjunction
		input = input[1:]
	}
	parts := bytes.Split(input, []byte(" -> "))
	name := parts[0]
	pp := bytes.Split(parts[1], []byte(", "))
	outputs := []string{}
	for _, p := range pp {
		outputs = append(outputs, string(p))
	}
	return Component{string(name), cType, []string{}, outputs, state}
}

type Puzzle struct {
	system map[string]Component
}

func configureInputs(system map[string]Component) {
	for key, value := range system {
		outputs := value.outputs
		for _, output := range outputs {
			v, ok := system[output]
			if !ok {
				continue
			}
			v.inputs = append(system[output].inputs, key)
			if v.ctype == conjunction {
				v.state = append(v.state, false)
			}
			system[output] = v
		}

	}
}

func readIput(filename string) Puzzle {
	input, _ := os.ReadFile(filename)
	lines := bytes.Split(input, []byte("\n"))
	system := map[string]Component{}
	for _, line := range lines {
		c := createComponent(line)
		system[c.name] = c
	}

	configureInputs(system)

	return Puzzle{system}
}

// Function to calculate the greatest common divisor (GCD)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate the least common multiple (LCM)
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

var kb *int
var jt *int
var ks *int
var sx *int

var (
	sumkb = 0
	sumjt = 0
	sumks = 0
	sumsx = 0
)

func processEvents(i int, system map[string]Component, events []Event) (int, int, bool) {
	highCount := 0
	lowCount := 0

	for {
		if len(events) == 0 {
			break
		}
		nEvents := []Event{}

		for _, event := range events {
			if event.state {
				highCount++
			} else {
				lowCount++
			}
			dest := system[event.destination]
			if i >= 0 {
				if event.source == "kb" && event.state {
					if kb != nil {
						sumkb = i - *kb
					}
					kb = &i
				}
				if event.source == "jt" && event.state {
					if jt != nil {
						sumjt = i - *jt
					}
					jt = &i
				}
				if event.source == "ks" && event.state {
					if ks != nil {
						sumks = i - *ks
					}
					ks = &i
				}
				if event.source == "sx" && event.state {
					if sx != nil {
						sumsx = i - *sx
					}
					sx = &i
				}
				if sumkb > 0 && sumjt > 0 && sumks > 0 && sumsx > 0 {
					ll := lcm(lcm(lcm(sumkb, sumjt), sumks), sumsx)
					return ll, 0, true
				}
			}
			if dest.name == "broadcaster" {
				for _, out := range system["broadcaster"].outputs {
					nEvents = append(nEvents, Event{out, dest.name, event.state})
				}
			} else if dest.ctype == flipflop {
				if !event.state {
					dest.state[0] = !dest.state[0]
					for _, out := range dest.outputs {
						nEvents = append(nEvents, Event{out, dest.name, dest.state[0]})
					}
				}
			} else if dest.ctype == conjunction {
				idx := -1
				for i, v := range dest.inputs {
					if v == event.source {
						idx = i
					}
				}
				if idx < 0 {
					panic("")
				}
				dest.state[idx] = event.state
				allHigh := true
				for _, v := range dest.state {
					if !v {
						allHigh = false
						break
					}
				}
				for _, out := range dest.outputs {
					nEvents = append(nEvents, Event{out, dest.name, !allHigh})
				}
			}

		}

		events = nEvents
	}

	return highCount, lowCount, false
}

func part1(puzzle Puzzle) int {
	high, low := 0, 0
	for i := 0; i < 1000; i++ {
		nhigh, nlow, _ := processEvents(-1, puzzle.system, []Event{{"broadcaster", "button", false}})
		high += nhigh
		low += nlow
	}
	return high * low
}

func part2(puzzle Puzzle) int {
	for i := 0; true; i++ {
		r, _, finished := processEvents(i, puzzle.system, []Event{{"broadcaster", "button", false}})
		if finished {
			return r
		}
	}
	return -1
}

func Solve(input string) (interface{}, interface{}) {
	puzzle1 := readIput(input)
	puzzle2 := readIput(input)
	return part1(puzzle1), part2(puzzle2)
}

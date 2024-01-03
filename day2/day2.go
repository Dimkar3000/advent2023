package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Roll struct {
	color string
	count int
}

type Game struct {
	id    int
	rolls [][]Roll
}

func parseSet(set string) []Roll {
	items := strings.Split(strings.TrimSpace(set), ",")
	var results []Roll
	for _, v := range items {
		parts := strings.Split(strings.TrimSpace(v), " ")
		c, _ := strconv.Atoi(parts[0])
		item := Roll{
			count: c,
			color: parts[1],
		}
		results = append(results, item)
	}

	return results
}

func parseGame(line string) Game {
	items := strings.Split(line, ":")
	game := Game{}
	i, _ := strconv.Atoi(strings.Split(items[0], " ")[1])
	game.id = i
	sets := strings.Split(strings.TrimSpace(items[1]), ";")
	for _, v := range sets {
		items := parseSet(v)
		game.rolls = append(game.rolls, items)
	}
	return game
}

func parseInput(filename string) []Game {
	readFile, err := os.Open(filename)
	var games []Game
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		game := parseGame(line)
		games = append(games, game)
	}
	readFile.Close()
	return games
}

func gameValid(game *Game) bool {
	for _, set := range game.rolls {
		for _, roll := range set {
			if roll.color == "red" && roll.count > 12 {
				return false
			}
			if roll.color == "green" && roll.count > 13 {
				return false
			}
			if roll.color == "blue" && roll.count > 14 {
				return false
			}
		}
	}
	return true
}

func minGame(game *Game) int {
	greenMin := 0
	blueMin := 0
	redMin := 0

	for _, set := range game.rolls {
		for _, roll := range set {
			if roll.color == "red" && roll.count > redMin {
				redMin = roll.count
			}
			if roll.color == "blue" && roll.count > blueMin {
				blueMin = roll.count
			}
			if roll.color == "green" && roll.count > greenMin {
				greenMin = roll.count
			}
		}

	}

	return greenMin * blueMin * redMin
}

func part1(games []Game) int {
	sum := 0
	for _, game := range games {
		if gameValid(&game) {
			sum += game.id
		}
	}
	return sum
}
func part2(games []Game) int {
	sum := 0
	for _, game := range games {
		sum += minGame(&game)
	}
	return sum
}

func Solve(filename string) (interface{}, interface{}) {
	games := parseInput(filename)
	p1 := part1(games)
	p2 := part2(games)
	return p1, p2
}

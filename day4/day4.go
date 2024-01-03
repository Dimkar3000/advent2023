package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/juliangruber/go-intersect"
)

type Game struct {
	index   int
	winning []int
	got     []int
}

func readInput(filename string) []Game {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var sum []Game
	i := 1
	for fileScanner.Scan() {
		input := fileScanner.Text()
		game := Game{
			index: i,
		}
		i++
		games := strings.Split(strings.Split(input, ":")[1], "|")

		wining := []int{}
		for _, w := range strings.Split(strings.TrimSpace(games[0]), " ") {
			v, _ := strconv.Atoi(w)
			if v == 0 {
				continue
			}
			wining = append(wining, v)
		}
		game.winning = wining

		got := []int{}
		for _, g := range strings.Split(strings.TrimSpace(games[1]), " ") {
			v, _ := strconv.Atoi(g)
			if v == 0 {
				continue
			}
			got = append(got, v)
		}
		game.got = got
		sum = append(sum, game)
	}
	readFile.Close()
	return sum
}

func part1(games []Game) int {
	sum := 0

	for _, game := range games {
		r := intersect.Simple(game.winning, game.got)
		power := len(r)
		if power > 0 {
			sum += 1 << (power - 1)
		}
	}

	return sum
}

func part2(games []Game) int {
	cards := []int{}
	for i := 0; i < len(games); i++ {
		cards = append(cards, 1)
	}
	for i := 0; i < len(games); i++ {
		r := len(intersect.Simple(games[i].winning, games[i].got))
		for j := i + 1; j <= i+r; j++ {
			cards[j] += cards[i]
		}
	}

	sum := 0
	for _, v := range cards {
		sum += v
	}

	return sum
}

func Solve(filename string) (interface{}, interface{}) {
	lines := readInput(filename)
	return part1(lines), part2(lines)

}

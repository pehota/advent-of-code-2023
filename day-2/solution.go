package day2

import (
	"regexp"
	"strconv"
	"strings"
)

func Run() (int, error) {
	input := getInput()
	games := parseInput(input)
	constraints := Draw{red: 12, green: 13, blue: 14}
	result := 0

	for _, game := range games {
		if gameWithinConstraints(game, constraints) {
			result += game.id
		}
	}
	return result, nil
}

type Draw struct {
	red   int
	green int
	blue  int
}

type Game struct {
	draws []Draw
	id    int
}

func extractGameId(line string) int {
	re := regexp.MustCompile(`Game ([\d]{1,})`)
	matches := strings.Split(re.FindString(line), " ")
	id, _ := strconv.Atoi(matches[1])
	return id
}

func extractDraws(line string) []Draw {
	outcomesInput := strings.Split(line, ":")[1]
	outcomes := strings.Split(outcomesInput, ";")
	result := make([]Draw, 0)

	for _, outcome := range outcomes {
		re := regexp.MustCompile(`([\d]{1,}) (red|green|blue)`)
		matches := re.FindAllString(outcome, -1)
		tmp := map[string]int{"red": 0, "green": 0, "blue": 0}

		for _, match := range matches {
			keyVal := strings.Split(match, " ")
			val, _ := strconv.Atoi(keyVal[0])
			tmp[keyVal[1]] = val
		}

		result = append(result, Draw{red: tmp["red"], green: tmp["green"], blue: tmp["blue"]})
	}

	return result
}

func parseGame(line string) (Game, error) {
	gameId := extractGameId(line)
	draws := extractDraws(line)

	return Game{id: gameId, draws: draws}, nil
}

func parseInput(input []string) []Game {
	result := make([]Game, 0)

	for _, line := range input {
		game, _ := parseGame(line)
		result = append(result, game)
	}

	return result
}

func gameWithinConstraints(game Game, constraints Draw) bool {
	for _, draw := range game.draws {
		if draw.blue > constraints.blue || draw.red > constraints.red || draw.green > constraints.green {
			return false
		}
	}
	return true
}

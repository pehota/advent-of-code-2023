package day2

import (
	"reflect"
	"testing"
)

func TestGameParsing(t *testing.T) {
	input := map[string]parseResult{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green": {1, []gameOutcome{
			{red: 4, green: 0, blue: 3}, {red: 1, green: 2, blue: 6}, {red: 0, green: 2, blue: 0},
		}},
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue": {2, []gameOutcome{
			{red: 0, green: 2, blue: 1}, {red: 1, green: 3, blue: 4}, {red: 0, green: 1, blue: 1},
		}},
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": {3, []gameOutcome{
			{red: 20, green: 8, blue: 6}, {red: 4, green: 13, blue: 5}, {red: 1, green: 5, blue: 0},
		}},
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": {4, []gameOutcome{
			{red: 3, green: 1, blue: 6}, {red: 6, green: 3, blue: 0}, {red: 14, green: 3, blue: 15},
		}},
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green": {5, []gameOutcome{
			{red: 6, green: 3, blue: 1}, {red: 1, green: 2, blue: 2},
		}},
	}

	for gameInput, expectedResult := range input {
		result, _ := parseGame(gameInput)
		if !reflect.DeepEqual(result, expectedResult) {
			t.Fatalf("Expected: %o; received: %o", result, expectedResult)
		}
	}
}

func TestGameWithinConstraints(t *testing.T) {
	input := map[string]bool{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green":                   true,
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue":         true,
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": false,
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": false,
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green":                   true,
	}

	for gameInput, expectedResult := range input {
		game, _ := parseGame(gameInput)
		constraints := gameOutcome{red: 12, green: 13, blue: 14}
		result := gameWithinConstraints(game, constraints)
		if result != expectedResult {
			t.Fatalf("Expected: %t; received: %t; game: %o; constraints: %o", result, expectedResult, game, constraints)
		}
	}
}

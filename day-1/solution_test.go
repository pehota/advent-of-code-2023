package solution

import (
	"testing"
)

func TestExtractNumber(t *testing.T) {
	inputExpectationsMap := map[string]int{
		"1abc2":       12,
		"pqr3stu8vwx": 38,
		"a1b2c3d4e5f": 15,
		"treb7uchet":  77,
	}

	for input, expectedResult := range inputExpectationsMap {
		result := extractNumber(input)
		if result != expectedResult {
			t.Fatalf("Expected %d, received: %d", expectedResult, result)
		}
	}
}

func TestExtractNumberWithWords(t *testing.T) {
	inputExpectationsMap := map[string]int{
		"two1nine":         29,
		"eightwothree":     83,
		"abcone2threexyz":  13,
		"xtwone3four":      24,
		"4nineeightseven2": 42,
		"zoneight234":      14,
		"7pqrstsixteen":    76,
		"eighthree":        83,
		"sevenine":         79,
	}

	for input, expectedResult := range inputExpectationsMap {
		result := extractNumber(input)
		if result != expectedResult {
			t.Fatalf("Expected %d, received: %d", expectedResult, result)
		}
	}
}

func TestCalculateSum(t *testing.T) {
	input := []int{1, 0, 2}
	expectedResult := 3
	result := calculateSum(input)
	if result != expectedResult {
		t.Fatalf("Expected %d, received: %d", expectedResult, result)
	}
}

package solution

import (
	"testing"
)

func TestExtractNumber(t *testing.T) {
	input := "treb7uchet"
	expectedResult := 77
	// pqr3stu8vwx
	// a1b2c3d4e5f
	// treb7uchet"
	result := extractNumber(input)

	if result != expectedResult {
		t.Fatalf("Expected %d, received: %d", expectedResult, result)
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

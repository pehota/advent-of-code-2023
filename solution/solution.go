package solution

import (
	"errors"
	"regexp"
	"strconv"
)

func calculateSum(input []int) int {
	sum := 0
	for _, val := range input {
		sum += val
	}
	return sum
}

func prepareInput(input []string) []int {
	mapped := make([]int, len(input))
	for i, val := range input {
		mapped[i] = extractNumber(val)
	}
	return mapped
}

func Run(input []string) (int, error) {
	if len(input) < 1 {
		return -1, errors.New("empty input")
	}
	mapped := prepareInput(input)
	sum := calculateSum(mapped)
	return sum, nil
}

func extractNumber(line string) int {
	re, _ := regexp.Compile(`\d`)
	matches := re.FindAllString(line, -1)
	result := matches[0] + matches[len(matches)-1]
	conv, _ := strconv.Atoi(result)
	return conv
}

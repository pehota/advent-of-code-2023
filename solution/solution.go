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

func normaliseMatch(input string) string {
	wordToIntMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	val, keyExists := wordToIntMap[input]

	if keyExists {
		return val
	}
	return input
}

func reverseString(s string) string {
	reversed := ""
	for _, v := range s {
		reversed = string(v) + reversed
	}
	return reversed
}

func extractNumber(line string) int {
	re, _ := regexp.Compile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	endRe, _ := regexp.Compile(`(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|\d)`)
	firstMatch := re.FindString(line)
	lastMatch := reverseString(endRe.FindString(reverseString(line)))

	result := normaliseMatch(firstMatch) + normaliseMatch(lastMatch)
	conv, _ := strconv.Atoi(result)
	return conv
}

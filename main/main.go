package main

import (
	"fmt"
	"log"

	"github.com/pehota/advent-of-code/day-1"
)

func main() {
	input := getInput()
	result, err := solution.Run(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %d", result)
}

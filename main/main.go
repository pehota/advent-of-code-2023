package main

import (
	"fmt"
	"log"

	"github.com/pehota/advent-of-code/day-2"
)

func main() {
	result, err := day2.RunPart2()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %d", result)
}

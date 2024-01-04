package main

import (
	"fmt"
	"log"

	"github.com/pehota/advent-of-code/day-1"
)

func main() {
	result, err := day1.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %d", result)
}

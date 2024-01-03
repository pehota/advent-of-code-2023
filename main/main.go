package main

import (
	"fmt"
	"log"

	"example.com/advent-of-code/day-1"
)

func main() {
	message, err := day1.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

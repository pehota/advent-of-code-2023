package main

import (
	"fmt"
	"log"

	"github.com/pehota/advent-of-code/solution"
)

func main() {
	input := getInput()
	message, err := solution.Run(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

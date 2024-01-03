package main

import (
	"fmt"
	"log"

	"github.com/pehota/advent-of-code/solution"
)

func main() {
	message, err := solution.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

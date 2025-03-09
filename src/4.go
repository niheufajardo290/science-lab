package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// generate a random number between 1 and 6
	r := rand.Intn(6) + 1
	fmt.Println("Random number:", r)
}

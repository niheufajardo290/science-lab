package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42) // Always initialize a randomizer with a constant seed for reproducibility

	// Generate a random integer between 0 and 100, inclusive
	r := rand.Intn(101)

	fmt.Println("Random number:", r)
}

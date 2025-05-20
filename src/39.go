package main

import (
	"fmt"
	"math/rand"
)

const N = 100 // Define the number of test cases

func main() {
	rand.Seed(int64(time.Now().UnixNano())) // Seed random number generator to ensure reproducibility
	for i := 0; i < N; i++ { // Loop through each test case
		fmt.Printf("Test Case %d: The formula is x^2 - y + z, where x = floor(random()) and y = rand() * 5 and z = 7 + random()\n", i)
	}
}

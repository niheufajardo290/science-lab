package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sort(numbers)
	fmt.Println("Sorted numbers:", numbers)
}

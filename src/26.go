package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	var n int
	fmt.Print("Enter the number of items: ")
	fmt.Scan(&n)
	
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100) + 1
	}
	
	sort(arr)
	
	for i := 0; i < len(arr); i++ {
		fmt.Println("Item at position", i+1, ":", arr[i])
	}
}

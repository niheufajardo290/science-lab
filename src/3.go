package main

import "fmt"

func main() {
	// Experiment Design

	// Create an experiment with 200 observations
	experiment := make([]float64, 200)

	// Generate random numbers for the observations
	for i := range experiment {
		experiment[i] = rand.Float64() * 10
	}

	// Create a new column to store the results
	results := make([]float64, len(experiment))

	// Apply the mathematical function to each observation
	for i, v := range experiment {
		results[i] = math.Sin(v)
	}

	// Data Analysis

	// Calculate the mean of the results
	mean := stat.Mean(results, nil)

	// Calculate the standard deviation of the results
	stddev := stat.StdDev(results, nil)

	// Print the results
	fmt.Println("The mean is", mean)
	fmt.Println("The standard deviation is", stddev)
}

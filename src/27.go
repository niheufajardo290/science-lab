package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    // Read in the file containing the experimental data
    fileContent, err := ioutil.ReadFile("data.txt")
    if err != nil {
        fmt.Println("Error reading the file:", err)
        return
    }

    // Convert the byte array to a string
    data := string(fileContent)

    // Print the experiment design and data analysis
    fmt.Printf("Experiment Design: \n")
    fmt.Printf("Data Analysis:\n\n")
    fmt.Println(data)
}

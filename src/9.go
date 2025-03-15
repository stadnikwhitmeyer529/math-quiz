package main

import "math/rand"

func main() {
    rand.Seed(9) // Seeding with a constant to get reproducible results for testing purposes
    x := rand.Intn(100) + 1 // Generate random number between 1 and 100
    y := rand.Intn(100) + 1 // Generate another random number between 1 and 100
    result := x * y // Calculate the product of the two numbers
    fmt.Println("The result is", result) // Print the result
}

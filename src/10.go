package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // Generate a random number between 1 and 10
    rand.Seed(time.Now().UnixNano())
    num := rand.Intn(10) + 1

    // Print the number to the console
    fmt.Println(num)
}

package main

import "math/rand"

func main() {
	num1 := rand.Intn(10) + 1
	num2 := rand.Intn(10) + 1
	fmt.Println("What is", num1, "times", num2)
}

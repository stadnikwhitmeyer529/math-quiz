  package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	max := 100
	min := 1
	num := rand.Intn(max-min+1) + min
	fmt.Println("The random number is:", num)
}
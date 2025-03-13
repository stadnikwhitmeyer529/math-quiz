package main

import "math/rand"

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

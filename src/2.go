package main

import "math"

func GenerateRandomMathQuiz() string {
	operation := randInt(2) // 1 or 2
	number1 := randInt(10) + 1
	number2 := randInt(10) + 1

	switch operation {
	case 1:
		return strconv.Itoa(number1) + "+" + strconv.Itoa(number2)
	case 2:
		return strconv.Itoa(number1) + "-" + strconv.Itoa(number2)
	}

	return ""
}

func randInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

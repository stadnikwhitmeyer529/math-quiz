package main

import "fmt"

func main() {
    var s1, s2 int
    fmt.Println("Please enter two numbers:")
    fmt.Scan(&s1, &s2)
    if s1 > s2 {
        fmt.Println("The first number is greater than the second number.")
    } else {
        fmt.Println("The first number is less than or equal to the second number.")
    }
}

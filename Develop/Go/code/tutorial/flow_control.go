package main

import (
    "fmt"
    "math"
    "time"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func sum(num int) int {
    sum := 0
    for i := 0; i < num; i++ {
        sum += i
    }
    return sum
}

func now() string {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		return "Good morning!"
	case t.Hour() < 17:
		return "Good afternoon."
	default:
		return "Good evening."
	}
}

func main() {
	defer fmt.Println("world") // stack, lazy evaluation

	fmt.Println(sum(10))
    fmt.Println(sqrt(-4))
    fmt.Println(now())
    fmt.Println("hello")
}


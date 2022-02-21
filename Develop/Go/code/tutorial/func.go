package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func compute_sqrt(x float64, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func compute_sqrt_3_4(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(swap("hello", "world"))

	hypot := compute_sqrt
	fmt.Println(hypot(5, 12))
	fmt.Println(compute_sqrt_3_4(hypot))
}

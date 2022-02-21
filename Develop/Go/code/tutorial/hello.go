package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
)

func main() {
    fmt.Println("The time is", time.Now())
    fmt.Println("My favorite number is ",rand.Intn(10))
    fmt.Println(math.Pi) // Exported value
}


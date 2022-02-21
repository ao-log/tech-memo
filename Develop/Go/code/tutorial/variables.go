package main

import (
    "fmt"
    "math"
)

const Pi = 3.14

func main() {
    fmt.Println(Pi)

    var i int // zero value
    var j int = 1
    k := 2 // Short variable declarations
    fmt.Println(i, j, k)

    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f) // Type conversions
    fmt.Println(x, y, z)
}


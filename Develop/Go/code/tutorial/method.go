package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer receiver。参照(メモリのアドレス)を渡しているので Struct の中身を更新できる。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ポインタを受け取る関数
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.X)

	fmt.Println(v.Abs())
	v.Scale(10) // (&v).Scale(5) として解釈される。
	fmt.Println(v.Abs())

	Scale(&v, 10)
	fmt.Println(v.Abs())
}

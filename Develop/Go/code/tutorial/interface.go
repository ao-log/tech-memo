package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// インタフェース M() を実装。
func (t *T) M() {
	// nil を扱う
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

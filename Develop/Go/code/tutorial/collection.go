package main

import (
	"fmt"
)

func array2(s1 string, s2 string) [2]string {
	var a [2]string
	a[0] = s1
	a[1] = s2
	return a
}

func slice() {
	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	b := names[1:3] // [Paul George]
	b[0] = "XXX"    // override names[1]

	printSlice(names[:])  // len=4 cap=4 [John XXX George Ringo]
	printSlice(names[:2]) // len=2 cap=4 [John XXX]
	printSlice(names[1:]) // len=3 cap=3 [XXX George Ringo]

	names = append(names, "Alice")
	printSlice(names) // len=5 cap=8 [John XXX George Ringo Alice]

	// range
	for i, v := range names {
		fmt.Println(i, v)
	}

}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func map_temperature() {
	// map[type of key]type of value
	tempelature := map[string]int{
		"Tokyo":    22,
		"Hokkaido": 10,
	}
	fmt.Printf("Temperature(Tokyo)=%d\n", tempelature["Tokyo"])
}

func main() {
	fmt.Println(array2("Hello", "world"))

	slice()

	map_temperature()
}

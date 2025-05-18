package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// Naked return
func nakedReturn(x, y int) (a, b int) {
	a = x
	b = y
	return
}

var c, python, java, goLang = false, false, false, true

func Module1() {
	fmt.Println("===MODULE 1===")

	fmt.Println(("Hello, World!"))
	fmt.Println(add(1, 2))
	fmt.Println(swap("Hello", "World"))

	// Short variable declaration
	a, b := nakedReturn(1, 2)
	fmt.Println("nakedReturn:", a, b)

	var i int
	fmt.Println(i, c, python, java, goLang)

	x, y := 3, 4
	z := math.Sqrt(float64(x*x) + float64(y*y))
	fmt.Println("Hypotenuse:", z)

	const Truthy = true
	fmt.Println("Truthy:", Truthy)
}

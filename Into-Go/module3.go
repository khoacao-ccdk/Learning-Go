package main

import (
	"fmt"
	"math"
	"strings"
)

func pointer() {
	x, y := 1, 20
	p := &x
	fmt.Println("p:", *p)

	*p = 21 // Setting the value of x through the pointer
	fmt.Println(x)

	p = &y
	*p = *p / 2 // Setting the value of y through the pointer
	fmt.Println(y)

	fmt.Println("p:", p) // Printing the address of y
}

func array() {
	prime := [10]int{1, 2, 3, 5, 7}
	fmt.Println(prime)

	helloArr := [2]string{"hello", "world"}
	fmt.Println(helloArr[0], helloArr[1])

	slide := prime[1:4] // Slicing the array
	fmt.Println(slide)

	prime[2] = 100
	fmt.Println(slide) // Slicing reflects the change in the original array
}

func slice() {
	prime := [10]int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	slice := prime[:4]
	printSlice("Initial slice", slice) // Initial slice

	slice = slice[2:]
	printSlice("Re-sliced", slice) // Reslicing the slice
	slice = prime[:0]
	printSlice("Zero-length slice", slice) // Zero-length slice
}

func nilSlice() {
	var s []int                    // nil slice
	fmt.Println(s, len(s), cap(s)) // nil slice has length and capacity of 0

	if s == nil {
		fmt.Println("s is nil")
	}
}

func makeArray() {
	a := make([]int, 5)
	printSlice("a", a)

	// Third argument specifies the capacity
	b := make([]int, 5, 10)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	// Capacity is defined by the starting point
	// In this case, d starts at index 2 of b
	// which means it has a capacity of 8
	d := b[2:5]
	printSlice("d", d)
}

func appendSlice() {
	var s []int
	printSlice("s", s)

	s = append(s, 0)
	printSlice("s after append", s)

	// This will allocate a new array
	// and copy the old elements to it
	// Resulting slice will have a new capacity
	// and a new address
	s = append(s, 1, 2, 3, 4)
	printSlice("s after multiple append", s)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func rangeLoop() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Println("Index:", i, "Value:", v)
	}

	fmt.Println(("Ignoring the value in the loop"))
	for i := range pow {
		fmt.Println("Index:", i) // Ignoring the value
	}

	fmt.Println("Ignoring the index in the loop")
	for _, v := range pow {
		fmt.Println("Value:", v) // Ignoring the index
	}
}

func mapIntro() {
	var m map[string]Vertex // Declare a map variable
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{40, -74}
	fmt.Println(m["Bell Labs"])

	var m2 = map[string]Vertex{
		"Google":   {37, -122},
		"Facebook": {37, -122},
	}
	fmt.Println(m2)

	m["Microsoft"] = Vertex{47, -122}
	fmt.Println(m)
	elem := m["Microsoft"]
	fmt.Println("Microsoft:", elem)
	delete(m, "Microsoft") // Deleting an element from the map
	fmt.Println("After deletion:", m)
	elem, ok := m["Microsoft"] // Checking if an element exists

	// If the element does not exist, ok will be false
	// and elem will be the zero value of Vertex
	// In this case, it will be Vertex{0, 0}
	fmt.Println("Microsoft exists?", ok, "Value:", elem)
}

func WordCount(s string) map[string]int {
	countMap := make(map[string]int)

	for _, word := range strings.Fields(s) {
		var _, ok = countMap[word]

		if !ok {
			countMap[word] = 1
		} else {
			countMap[word]++
		}
	}

	return countMap
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type Vertex struct {
	x, y int
}

func Module3() {
	fmt.Println("===MODULE 3===")
	pointer()

	v := Vertex{1, 2}
	p := &v
	p.x = 1e9 // Setting the value of x through the pointer
	fmt.Println(p.x, p.y)
	fmt.Println(Vertex{x: 1})
	fmt.Println(Vertex{})

	p = &Vertex{1, 2} // Has type *Vertex
	fmt.Println(p)

	array()
	slice()
	nilSlice()
	makeArray()
	appendSlice()
	rangeLoop()
	mapIntro()

	fmt.Println("Word Count:", WordCount("I am learning Go. Go is fun! Go is powerful."))

	hypotenuse := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("Hypotenuse alone:", hypotenuse(5, 6))
	fmt.Println("Hypotenuse using compute:", compute(hypotenuse))
	fmt.Println("Math.Pow using compute:", compute(math.Pow))

	fmt.Println("Adder function:")
	pos, neg := adder(), adder()

	for i := range 10 {
		fmt.Println("Positive adder:", pos(i))
		fmt.Println("Negative adder:", neg(-2*i))
	}
}

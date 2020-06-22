package main

import (
	"math"
	"fmt"
)

func main() {
	// hello world
	fmt.Println("hello world")

	// arithmetic
	fmt.Println(1 + 1)
	fmt.Println(1 * 2)
	fmt.Println(2 - 1)
	fmt.Println(3 / 2)
	fmt.Println(53 % 7)
	fmt.Println(math.Sqrt(5))

	// string
	a := "geleia"
	fmt.Println(a + " talk")
	fmt.Println(a[3])
	fmt.Println(string(a[3]))

	// array's
	l := []int{1, 2, 3, 4, 5}
	for i, v := range l {
		fmt.Println(i, v)
	}

	m := []([]float64){
		[]float64{0.1, 0.2, 0.3},
		[]float64{0.4, 0.5, 0.6},
		[]float64{0.7, 0.8, 0.9},
	}
	fmt.Println(m)

	// struct
	type point struct {
		x, y float64
	}
	p1 := point{x: 0, y: 0}
	p2 := point{
		x: 1,
		y: 1,
	}
	fmt.Println(p1, p2)
}

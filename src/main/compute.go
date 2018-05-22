package main

import "fmt"

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func Compute() {
	hypot := func(x, y float64) float64 {
		return x * y
	}

	fmt.Println(hypot(2, 3))
	fmt.Println(compute(hypot))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	a := 27.0
	b := Solver(a)
	c := math.Cbrt(a)
	fmt.Println("Solver: ", b)
	fmt.Println("math.Cbrt: ", c)
	fmt.Println("The difference: ", b-c)
}

func Solver(x float64) float64 {
	z := x / 2

	for {
		a := z
		z -= (z*z*z - x) / (3 * z * z)
		fmt.Println(z)
		if a >= z && (a-z) < 1e-8 {
			return z
		}
	}
}

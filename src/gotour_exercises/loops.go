package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := x / 2

	for z*z-x > math.Abs(0.00000001) {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

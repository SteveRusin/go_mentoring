package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	input := 25.0
	fmt.Println(Sqrt(input))
	fmt.Println(math.Sqrt(input))
}

package main

import (
	"fmt"
	"math"
)


func is_small(n float64) bool {
	return math.Abs(n) < 0.0000000001
}


func NewtonRaphson(x int) (float64, int) {
	z := 1.0
	iter := 0
	for {
		dz := (math.Pow(z,2) - float64(x)) / (2*z)
		if is_small(dz) {
			break
		}
		z -= dz
		iter++
		fmt.Printf("Iteration %d: z = %.6f, dz = %.6f\n", iter, z, dz)
	}

	return z, iter
}


func main() {
	x := 150
	z, iter := NewtonRaphson(x)

	fmt.Println(z, iter)

}



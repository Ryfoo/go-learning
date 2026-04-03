package main

import (
	"errors"
	"fmt"
)

// Exercise: Functions
//
// Functions in Go are declared with the func keyword.
// Go supports multiple return values, which is commonly used to return
// a result alongside an error.
//
// Syntax:
//   func name(param type) returnType { ... }
//   func name(a, b int) (int, error) { ... }
//
// A variadic function accepts a variable number of arguments using ...
//   func sum(nums ...int) int { ... }
//
// Functions are first-class values and can be assigned to variables.

// add returns the sum of two integers.
// TODO: Implement this function.
func add(a, b int) int {
	// TODO: return a + b
	return 0
}

// divide returns the result of a / b.
// If b is zero, return 0 and an error with the message "division by zero".
// Use errors.New("division by zero") to create the error.
// TODO: Implement this function.
func divide(a, b float64) (float64, error) {
	// TODO: check for zero divisor, then return the result
	return 0, nil
}

// sumAll returns the sum of all provided integers.
// It must accept any number of arguments.
// TODO: Implement using a variadic parameter.
func sumAll(nums ...int) int {
	// TODO: iterate over nums and accumulate the total
	return 0
}

// applyTwice applies the function f to x twice: f(f(x)).
// Functions are values in Go and can be passed as arguments.
// TODO: Implement this function.
func applyTwice(f func(int) int, x int) int {
	// TODO: return f(f(x))
	return 0
}

func functionsExercise() {
	fmt.Println(add(3, 4)) // 7

	result, err := divide(10, 3)
	fmt.Printf("%.4f %v\n", result, err) // 3.3333 <nil>

	_, err = divide(5, 0)
	fmt.Println(err) // division by zero

	fmt.Println(sumAll(1, 2, 3, 4, 5)) // 15

	double := func(n int) int { return n * 2 }
	fmt.Println(applyTwice(double, 3)) // 12

	_ = errors.New // remove when you use errors.New
}

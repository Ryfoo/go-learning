package main

import "fmt"

// Exercise: Pointers
//
// A pointer holds the memory address of a value.
//   var p *int        // pointer to int, zero value is nil
//   p = &x            // & gives the address of x
//   *p = 10           // * dereferences: reads or writes the value at p
//
// Passing a pointer to a function lets the function modify the original value.
// new(T) allocates a zeroed value of type T and returns a pointer to it.

// increment increments the integer pointed to by p by 1.
// TODO: Implement by dereferencing p and adding 1.
func increment(p *int) {
	// TODO: *p = ...
}

// swap exchanges the values of the two integers pointed to by a and b.
// TODO: Implement using a temporary variable.
func swap(a, b *int) {
	// TODO: swap *a and *b
}

// newInt allocates a new int with the given value and returns a pointer to it.
// Use the new built-in or take the address of a local variable.
// TODO: Implement this function.
func newInt(val int) *int {
	// TODO: return a pointer to a new int set to val
	return nil
}

// zeroOut sets every element of the slice to 0.
// Note: slices are already reference types, so a pointer to the slice is
// not needed to modify its elements.
// TODO: Implement using a range loop.
func zeroOut(s []int) {
	// TODO: set each element to 0
}

func pointersExercise() {
	x := 5
	increment(&x)
	fmt.Println(x) // 6

	a, b := 10, 20
	swap(&a, &b)
	fmt.Println(a, b) // 20 10

	p := newInt(42)
	fmt.Println(*p) // 42

	nums := []int{1, 2, 3, 4, 5}
	zeroOut(nums)
	fmt.Println(nums) // [0 0 0 0 0]
}

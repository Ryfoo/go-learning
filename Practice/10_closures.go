package main

import "fmt"

// Exercise: Closures and Higher-Order Functions
//
// A closure is a function value that captures variables from its surrounding scope.
// The captured variables are shared between the closure and the outer function.
//
// Higher-order functions accept functions as arguments or return them.
// This enables patterns like map, filter, and reduce over slices.
//
// Example closure:
//   func counter() func() int {
//       n := 0
//       return func() int { n++; return n }
//   }

// makeCounter returns a closure that increments and returns a private counter
// each time it is called. Each call to makeCounter produces an independent counter.
// TODO: Implement makeCounter.
func makeCounter() func() int {
	// TODO: declare a local counter variable, return a closure that increments it
	return nil
}

// makeAdder returns a closure that adds x to any value passed to it.
// Example: add5 := makeAdder(5); add5(3) == 8
// TODO: Implement makeAdder.
func makeAdder(x int) func(int) int {
	// TODO: return a closure that captures x
	return nil
}

// mapInts applies f to every element of s and returns the results in a new slice.
// TODO: Implement this higher-order function.
func mapInts(s []int, f func(int) int) []int {
	// TODO: allocate result slice, apply f to each element, return
	return nil
}

// filter returns a new slice containing only elements of s for which keep returns true.
// TODO: Implement this higher-order function.
func filter(s []int, keep func(int) bool) []int {
	// TODO: append elements that satisfy keep, return the slice
	return nil
}

// reduce folds s into a single value using f, starting from initial.
// f receives the accumulated value and the current element.
// Example: reduce([]int{1,2,3}, 0, func(acc, v int) int { return acc + v }) == 6
// TODO: Implement this higher-order function.
func reduce(s []int, initial int, f func(int, int) int) int {
	// TODO: accumulate and return the result
	return 0
}

func closuresExercise() {
	c1 := makeCounter()
	c2 := makeCounter()
	fmt.Println(c1(), c1(), c1()) // 1 2 3
	fmt.Println(c2(), c2())       // 1 2 (independent counter)

	add10 := makeAdder(10)
	fmt.Println(add10(5))  // 15
	fmt.Println(add10(20)) // 30

	nums := []int{1, 2, 3, 4, 5}
	doubled := mapInts(nums, func(n int) int { return n * 2 })
	fmt.Println(doubled) // [2 4 6 8 10]

	evens := filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println(evens) // [2 4]

	sum := reduce(nums, 0, func(acc, v int) int { return acc + v })
	fmt.Println(sum) // 15
}

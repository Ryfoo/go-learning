package main

import "fmt"

// Exercise: Arrays, Slices, and Maps
//
// Array: fixed-size sequence of elements of the same type.
//   var a [3]int = [3]int{1, 2, 3}
//
// Slice: a dynamic, flexible view over an array. Preferred over arrays.
//   s := []int{1, 2, 3}
//   s = append(s, 4)
//
// Map: an unordered collection of key-value pairs.
//   m := map[string]int{"a": 1}
//   m["b"] = 2
//   val, ok := m["a"]  // ok is false if key does not exist

// sumArray returns the sum of all elements in a fixed-size array of 5 integers.
// TODO: Implement using a for loop or range.
func sumArray(arr [5]int) int {
	// TODO: iterate over the array and return the sum
	return 0
}

// filterPositive returns a new slice containing only the positive numbers
// from the input slice (values strictly greater than 0).
// TODO: Implement using append and a range loop.
func filterPositive(nums []int) []int {
	// TODO: build and return the filtered slice
	return nil
}

// wordCount returns a map where each key is a word from the input slice
// and each value is the number of times that word appears.
// TODO: Implement using a map and a range loop.
func wordCount(words []string) map[string]int {
	// TODO: count occurrences and return the map
	return nil
}

func collectionsExercise() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sumArray(arr)) // 15

	nums := []int{-3, 0, 1, 4, -1, 7}
	fmt.Println(filterPositive(nums)) // [1 4 7]

	words := []string{"go", "is", "go", "fun", "go"}
	counts := wordCount(words)
	fmt.Println(counts["go"])  // 3
	fmt.Println(counts["fun"]) // 1
	fmt.Println(counts["is"])  // 1
}

package main

import (
	"fmt"
	"sync"
)

// Exercise: Goroutines and Channels
//
// A goroutine is a lightweight thread managed by the Go runtime.
// Start one with the go keyword: go myFunc()
//
// Channels are the primary way goroutines communicate.
//   ch := make(chan int)       // unbuffered: send blocks until received
//   ch := make(chan int, 10)   // buffered: send blocks only when full
//   ch <- value               // send
//   value := <-ch             // receive
//   close(ch)                 // signal no more values will be sent
//
// WaitGroup coordinates waiting for a set of goroutines to finish.
//   var wg sync.WaitGroup
//   wg.Add(1); go func() { defer wg.Done(); ... }(); wg.Wait()
//
// Mutex protects a shared value from concurrent writes.
//   var mu sync.Mutex
//   mu.Lock(); sharedVar++; mu.Unlock()

// generateSquares sends the squares of integers 1..n to a channel,
// then closes the channel.
// TODO: Run this in a goroutine (the caller does that), send n squares, close.
func generateSquares(n int, ch chan<- int) {
	// TODO: loop from 1 to n, send i*i on ch, close ch when done
}

// parallelSum splits nums into two halves, computes the sum of each half
// in a separate goroutine, and returns the total sum.
// Use a channel (or WaitGroup + shared variable) to collect partial results.
// TODO: Implement using goroutines and a channel.
func parallelSum(nums []int) int {
	// TODO: split, launch goroutines, collect results, return total
	return 0
}

// safeCounter uses a Mutex to protect a shared counter that is incremented
// concurrently by n goroutines. It returns the final counter value.
// TODO: Implement using sync.Mutex and sync.WaitGroup.
func safeCounter(n int) int {
	var mu sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	// TODO: launch n goroutines; each must lock, increment counter, unlock
	// Wait for all goroutines to finish, then return counter.

	wg.Wait()
	return counter
}

func goroutinesExercise() {
	// generateSquares
	ch := make(chan int)
	go generateSquares(5, ch)
	for v := range ch {
		fmt.Print(v, " ") // 1 4 9 16 25
	}
	fmt.Println()

	// parallelSum
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(parallelSum(nums)) // 36

	// safeCounter
	fmt.Println(safeCounter(1000)) // 1000
}

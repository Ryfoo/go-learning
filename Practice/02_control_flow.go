package main

import "fmt"

// Exercise: Control Flow
//
// Go provides the standard control flow constructs:
//   - if / else if / else
//   - for (Go's only loop — covers while and do-while patterns too)
//   - switch
//
// Unlike many languages, conditions do not need parentheses.
// A for loop with a single condition behaves like a while loop.
//
// Examples:
//   for i := 0; i < 5; i++ { ... }
//   for condition { ... }        // while-style
//   for { ... }                  // infinite loop

// grade returns a letter grade for the given score.
// Grading scale: A >= 90, B >= 80, C >= 70, D >= 60, F otherwise.
// TODO: Implement this function using if/else if/else.
func grade(score int) string {
	// TODO: return the correct letter grade
	return ""
}

// fizzBuzz prints numbers from 1 to n (inclusive).
// For multiples of 3, print "Fizz" instead of the number.
// For multiples of 5, print "Buzz" instead of the number.
// For multiples of both 3 and 5, print "FizzBuzz".
// TODO: Implement using a for loop and if/else.
func fizzBuzz(n int) {
	// TODO: loop from 1 to n and print the correct output
}

// dayName returns the name of the weekday for the given number (1=Monday ... 7=Sunday).
// Return "Unknown" for any other value.
// TODO: Implement using a switch statement.
func dayName(day int) string {
	// TODO: use switch to return the day name
	return ""
}

func controlFlowExercise() {
	fmt.Println(grade(95))  // A
	fmt.Println(grade(82))  // B
	fmt.Println(grade(55))  // F

	fizzBuzz(15)

	fmt.Println(dayName(1)) // Monday
	fmt.Println(dayName(7)) // Sunday
	fmt.Println(dayName(9)) // Unknown
}

package main

import (
	"errors"
	"fmt"
)

// Exercise: Error Handling
//
// Go treats errors as values. Functions signal failure by returning an error
// as the last return value. Callers must check this value explicitly.
//
// Creating errors:
//   errors.New("message")
//   fmt.Errorf("context: %w", err)   // wraps an underlying error
//
// Sentinel errors are package-level variables used for comparison:
//   var ErrNotFound = errors.New("not found")
//
// errors.Is checks whether a target error appears anywhere in the chain:
//   errors.Is(err, ErrNotFound)
//
// Custom error types implement the error interface (Error() string).
// Use a pointer receiver so the type satisfies the interface correctly.

// ErrNotFound is returned when a requested key does not exist in the store.
// TODO: Declare ErrNotFound as a sentinel error using errors.New.
var ErrNotFound = errors.New("") // TODO: fill in the message "not found"

// ValidationError is returned when input fails a validation check.
// It carries the field name and a human-readable reason.
// TODO: Define the ValidationError struct with Field and Reason string fields.
// TODO: Implement the error interface: Error() should return
//   "validation error on <Field>: <Reason>"
type ValidationError struct {
	// TODO: add Field and Reason fields
}

func (e *ValidationError) Error() string {
	// TODO: return the formatted message
	return ""
}

// store is a simple in-memory key-value store used for the exercises below.
var store = map[string]string{
	"name": "Alice",
	"city": "Paris",
}

// lookup returns the value for key in store.
// If the key does not exist, return ErrNotFound.
// TODO: Implement this function.
func lookup(key string) (string, error) {
	// TODO: check the map, return the value or ErrNotFound
	return "", nil
}

// validateAge checks that age is between 0 and 120 (inclusive).
// If age is out of range, return a *ValidationError with Field "age" and
// an appropriate Reason.
// TODO: Implement this function.
func validateAge(age int) error {
	// TODO: validate and return nil or a *ValidationError
	return nil
}

func errorHandlingExercise() {
	// lookup
	val, err := lookup("name")
	fmt.Println(val, err) // Alice <nil>

	_, err = lookup("email")
	fmt.Println(errors.Is(err, ErrNotFound)) // true

	// validateAge
	fmt.Println(validateAge(25))  // <nil>
	fmt.Println(validateAge(-1))  // validation error on age: ...
	fmt.Println(validateAge(200)) // validation error on age: ...

	// Type assertion to inspect the custom error
	err = validateAge(-1)
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println(ve.Field) // age
	}
}

package main

import (
	"fmt"
	"math"
)

// Exercise: Interfaces
//
// An interface defines a set of method signatures.
// A type satisfies an interface by implementing all of its methods.
// This is implicit — no "implements" keyword is needed.
//
// The empty interface (interface{} or any) is satisfied by all types.
// Type assertions extract a concrete value from an interface:
//   val, ok := i.(ConcreteType)
//
// A type switch dispatches on the dynamic type of an interface value:
//   switch v := i.(type) {
//   case int:    ...
//   case string: ...
//   }

// Shape is an interface for geometric shapes.
// TODO: Define the Shape interface with two methods:
//   Area() float64
//   Perimeter() float64

// Circle represents a circle with a Radius.
// TODO: Define the Circle struct with a float64 Radius field.
// TODO: Implement Area and Perimeter for Circle.
//   Area      = math.Pi * r * r
//   Perimeter = 2 * math.Pi * r

type Circle struct {
	// TODO: add Radius field
}

// Triangle represents a right triangle with legs A and B.
// TODO: Define the Triangle struct with float64 fields A and B.
// TODO: Implement Area and Perimeter for Triangle.
//   Area      = 0.5 * A * B
//   Hypotenuse = math.Sqrt(A*A + B*B)
//   Perimeter = A + B + Hypotenuse

type Triangle struct {
	// TODO: add A and B fields
}

// totalArea returns the sum of Area() for all shapes in the slice.
// TODO: Implement by ranging over the slice and calling Area() on each shape.
func totalArea(shapes []Shape) float64 {
	// TODO: accumulate and return the total
	return 0
}

// describe prints a type-specific description for each value in the slice.
// For int:    print "Integer: <value>"
// For string: print "String: <value>"
// For Shape:  print "Shape area: <Area()>"
// For other:  print "Unknown type"
// TODO: Implement using a type switch.
func describe(values []any) {
	// TODO: range over values and switch on type
}

func interfacesExercise() {
	c := Circle{Radius: 5}
	fmt.Printf("%.4f\n", c.Area())      // 78.5398
	fmt.Printf("%.4f\n", c.Perimeter()) // 31.4159

	t := Triangle{A: 3, B: 4}
	fmt.Printf("%.4f\n", t.Area())      // 6.0000
	fmt.Printf("%.4f\n", t.Perimeter()) // 12.0000

	shapes := []Shape{c, t}
	fmt.Printf("%.4f\n", totalArea(shapes)) // 84.5398

	describe([]any{42, "hello", c, 3.14})
	// Integer: 42
	// String: hello
	// Shape area: 78.5398
	// Unknown type

	_ = math.Pi // remove when you use math
}

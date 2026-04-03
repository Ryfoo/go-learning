package main

import "fmt"

// Exercise: Structs
//
// A struct groups related fields together under a single type.
// Methods are functions with a receiver argument that binds them to a type.
//
// Value receiver: func (r Rectangle) Area() float64 { ... }
//   - operates on a copy; use when the method does not modify the struct.
//
// Pointer receiver: func (r *Rectangle) Scale(factor float64) { ... }
//   - operates on the original; use when the method modifies the struct.
//
// Struct embedding lets one struct include another, inheriting its fields.

// Rectangle represents a rectangle with a width and height.
// TODO: Define the Rectangle struct with two float64 fields: Width and Height.
type Rectangle struct {
	// TODO: add Width and Height fields
}

// Area returns the area of the rectangle (Width * Height).
// TODO: Add this method to Rectangle using a value receiver.

// Perimeter returns the perimeter of the rectangle (2 * (Width + Height)).
// TODO: Add this method to Rectangle using a value receiver.

// Scale multiplies both Width and Height by the given factor.
// TODO: Add this method to Rectangle using a pointer receiver.

// ColoredRectangle embeds Rectangle and adds a Color field.
// TODO: Define this struct. Embed Rectangle (not as a named field) and add a
// Color string field.
type ColoredRectangle struct {
	// TODO: embed Rectangle and add Color
}

// Describe returns a string describing the colored rectangle.
// Format: "Color: <Color>, Area: <Area>"
// Example: "Color: red, Area: 12.00"
// Use fmt.Sprintf for formatting.
// TODO: Add this method to ColoredRectangle using a value receiver.
func (cr ColoredRectangle) Describe() string {
	// TODO: return the formatted string
	return ""
}

func structsExercise() {
	r := Rectangle{Width: 4, Height: 3}
	fmt.Println(r.Area())      // 12
	fmt.Println(r.Perimeter()) // 14

	r.Scale(2)
	fmt.Println(r.Width, r.Height) // 8 6

	cr := ColoredRectangle{
		Rectangle: Rectangle{Width: 4, Height: 3},
		Color:     "red",
	}
	fmt.Println(cr.Describe()) // Color: red, Area: 12.00
	// Embedded fields are promoted, so this also works:
	fmt.Println(cr.Area()) // 12
}

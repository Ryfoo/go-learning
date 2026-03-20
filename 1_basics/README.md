## **Golang Basics**


Every Go program is made up of packages.

Programs start running in package main. 

### Type Convention 
any variable name that starts with capital letter means thatit can be exported such as, Pi
while names that do not start with capital letters such as value and result cannot be exported outside the package

### Functions defintions
Go follows the following format for functions declaration/definition
func func_name(parameters : type) return_type { body } 
not that Go is statically-typed meaning that a variable cannot refer to str while it was declared as int for instance
functions in Go can return multiple values (tuples as return a, b, c,...)

### Go basic types
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

### Type Conversions
Go uses the format of T(v) to the convert the value v to type T
such as, var x int = 10, y = float64(x)


### Type Inference
When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable type is inferred from the value on the right hand side.
var i int
j := i (j is an int)
a := 42 (a is an int)


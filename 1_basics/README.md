# Go (Golang) Basics

Welcome! This guide walks you through the fundamentals of Go in a friendly, approachable way. Whether you are completely new or just need a refresher, you are in the right place.

---

## 🔴 Table of Contents

- [Packages](#packages)
- [Naming Conventions](#naming-conventions)
- [Variables & Declarations](#variables--declarations)
- [Basic Types](#basic-types)
- [Type Conversions](#type-conversions)
- [Type Inference](#type-inference)
- [Functions](#functions)
- [Constants](#constants)
- [Control Flow](#control-flow)
- [Loops](#loops)
- [Pointers](#pointers)

---

## 🟢 Packages

Every Go program is made up of **packages**. Think of a package as a folder of related code that you can share and reuse.

All programs start running in the special `main` package — that is where Go looks for the entry point of your application.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

You can import multiple packages at once using a grouped import, which is the preferred style:

```go
import (
    "fmt"
    "math"
)
```

---

## 🟡 Naming Conventions

Go uses a simple but important rule for controlling what gets shared between packages:

- 🟢 **Exported** (public) — names that **start with a capital letter** are accessible from other packages. Example: `Pi`, `Println`, `MyFunction`
- 🔴 **Unexported** (private) — names that **start with a lowercase letter** stay inside their own package. Example: `value`, `result`, `helperFunc`

This is Go's way of doing access control — no `public` or `private` keywords needed.

---

## 🟢 Variables & Declarations

There are a few ways to declare variables in Go.

**Using `var`** — explicit and clear, works anywhere:

```go
var name string = "Alice"
var age int = 30
```

**Short declaration with `:=`** — quick and concise, only works inside functions:

```go
name := "Alice"
age := 30
```

**Declaring multiple variables at once:**

```go
var (
    x int    = 10
    y float64 = 3.14
    z bool   = true
)
```

> 🟡 Note: Go does not allow unused variables. If you declare a variable and never use it, the compiler will throw an error. This keeps your code clean.

**Zero Values** — if you declare a variable without assigning a value, Go gives it a default "zero value":

| Type | Zero Value |
|------|-----------|
| `int` | `0` |
| `float64` | `0.0` |
| `bool` | `false` |
| `string` | `""` (empty string) |

```go
var count int    // count is 0
var active bool  // active is false
var label string // label is ""
```

---

## 🟡 Basic Types

Go is a **statically typed** language, which means every variable has a fixed type that cannot change.

**Boolean:**
```go
bool
```

**String:**
```go
string
```

**Integer types** (pick the size you need):
```go
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
```

**Handy aliases:**
```go
byte   // same as uint8  — often used for raw binary data
rune   // same as int32  — represents a single Unicode character
```

**Floating point:**
```go
float32  float64
```

**Complex numbers:**
```go
complex64  complex128
```

> 🟢 In practice, you will use `int`, `float64`, `string`, `bool`, and `rune` most of the time. The sized variants matter when you care about memory layout or are working with binary data.

---

## 🟢 Type Conversions

Go will never automatically convert between types — you have to do it explicitly. The format is `T(value)`, where `T` is the type you want to convert to.

```go
var x int = 10
var y float64 = float64(x)  // convert int to float64
var z int = int(y)           // convert float64 back to int (decimal is dropped)
```

```go
var r rune = 'A'
var b byte = byte(r)  // rune to byte
```

> 🔴 Be careful when converting from a larger type to a smaller one — you can lose data. For example, converting `float64(3.99)` to `int` gives you `3`, not `4`.

---

## 🟡 Type Inference

When you use `:=` or `var =` without specifying a type, Go figures out the type from the value on the right-hand side.

```go
var i int
j := i       // j is inferred as int

a := 42      // a is int
b := 3.14    // b is float64
c := "hello" // c is string
d := true    // d is bool
```

This makes your code shorter without sacrificing type safety.

---

## 🟢 Functions

Functions in Go follow this structure:

```
func functionName(parameterName type) returnType {
    // body
}
```

A simple example:

```go
func add(x int, y int) int {
    return x + y
}
```

**Shorthand for same-type parameters** — if multiple parameters share the same type, you only need to write the type once:

```go
func add(x, y int) int {
    return x + y
}
```

**Multiple return values** — Go functions can return more than one value, which is very useful for returning a result alongside an error:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

**Named return values** — you can name the return values and use a bare `return`:

```go
func minMax(arr []int) (min, max int) {
    min, max = arr[0], arr[0]
    for _, v := range arr {
        if v < min { min = v }
        if v > max { max = v }
    }
    return
}
```

> 🟡 Named returns are great for short functions but can make longer functions harder to read — use them with care.

---

## 🟢 Constants

Constants are declared with `const` and cannot be changed after they are set.

```go
const Pi = 3.14159
const Greeting = "Hello"
```

Go also has a special tool for creating sequences of related constants called `iota`:

```go
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
    Wednesday      // 3
)
```

`iota` starts at 0 and increments by 1 for each constant in the block. Very handy for enums.

---

## 🟡 Control Flow

**If / Else:**

```go
if x > 10 {
    fmt.Println("big number")
} else if x == 10 {
    fmt.Println("exactly ten")
} else {
    fmt.Println("small number")
}
```

Go also lets you run a short statement right before the condition — the variable is scoped to the if block:

```go
if result, err := someFunc(); err != nil {
    fmt.Println("error:", err)
} else {
    fmt.Println("result:", result)
}
```

**Switch:**

Go's `switch` is like the one in C, Java, or JavaScript, with two key differences:

- 🟢 Go automatically breaks after each case — you do not need to write `break`
- 🟡 Cases do not have to be constants or integers

```go
switch day {
case "Monday":
    fmt.Println("Start of the work week")
case "Friday":
    fmt.Println("Almost the weekend!")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Midweek grind")
}
```

You can also use a switch with no condition — it works like a clean chain of if/else:

```go
switch {
case age < 18:
    fmt.Println("Minor")
case age < 65:
    fmt.Println("Adult")
default:
    fmt.Println("Senior")
}
```

---

## 🟢 Loops

Go has only one looping keyword: `for`. But it is flexible enough to cover all your looping needs.

**Basic loop (like a C-style for):**

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

**While-style loop:**

```go
count := 0
for count < 10 {
    count++
}
```

**Infinite loop:**

```go
for {
    // runs forever until a break or return
}
```

**Loop over a collection with `range`:**

```go
nums := []int{2, 4, 6, 8}
for index, value := range nums {
    fmt.Println(index, value)
}
```

Use `_` to ignore either the index or the value:

```go
for _, value := range nums {
    fmt.Println(value)
}
```

---

## 🟡 Pointers

A pointer holds the **memory address** of a value rather than the value itself. Go uses pointers but keeps them simpler than C — there is no pointer arithmetic.

```go
x := 42
p := &x       // p is a pointer to x (& gives you the address)
fmt.Println(*p) // dereference p to get the value: prints 42

*p = 100      // change the value at that address
fmt.Println(x)  // x is now 100
```

- `&variable` — gives you the pointer (address of the variable)
- `*pointer` — gives you the value at that address (dereference)

> 🟢 Pointers are commonly used when you want a function to modify a variable, or when passing large structs without copying them.

---

That covers the core building blocks of Go! From here, explore **structs**, **interfaces**, **goroutines**, and **channels** to unlock the real power of the language. Happy coding!

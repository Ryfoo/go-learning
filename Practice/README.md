# Go Syntax Exercises

A set of skeleton files to practice Go's core features by completing missing implementations.

## Structure

Each file covers one concept:

| File | Concept |
|---|---|
| `01_variables.go` | Variables, constants, and basic types |
| `02_control_flow.go` | if/else, for loops, switch |
| `03_functions.go` | Multiple return values, variadic functions, first-class functions |
| `04_collections.go` | Arrays, slices, and maps |
| `05_structs.go` | Structs, methods, and embedding |
| `06_interfaces.go` | Interfaces, type assertions, and type switches |
| `07_errors.go` | Error handling, sentinel errors, and custom error types |
| `08_goroutines.go` | Goroutines, channels, WaitGroup, and Mutex |
| `09_pointers.go` | Pointers, dereferencing, and pass-by-pointer |
| `10_closures.go` | Closures and higher-order functions |

## How to Work Through the Exercises

1. Open a file and read the documentation at the top — it explains the concept and relevant syntax.
2. Find every `// TODO` comment and replace it with working code.
3. In `main.go`, uncomment the corresponding exercise function call.
4. Run the program and compare your output to the expected values shown in the comments.

```bash
go run .
```

## Requirements

- Go 1.18 or later.
- All files are in `package main`, so run them together from the same directory.

## Tips

- The expected output for each function call is written as a comment on the same line, for example: `fmt.Println(add(3, 4)) // 7`.
- Complete one file at a time and keep the others commented out in `main.go` to avoid compilation errors from unimplemented stubs.
- If the compiler reports "declared and not used", make sure you are using every variable you declare.

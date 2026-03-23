# Go (Golang) — Advanced Concepts

This guide picks up where the basics left off. Here you will find the concepts that make Go genuinely powerful and a joy to work with — methods, interfaces, generics, and concurrency. Take your time with each section; these ideas build on one another.

---

## 🔴 Table of Contents

- [Structs](#structs)
- [Methods](#methods)
- [Pointer vs Value Receivers](#pointer-vs-value-receivers)
- [Method & Pointer Indirection](#method--pointer-indirection)
- [Interfaces](#interfaces)
- [The Empty Interface](#the-empty-interface)
- [Type Assertions & Type Switches](#type-assertions--type-switches)
- [Stringer — Implementing fmt.Stringer](#stringer--implementing-fmtstringer)
- [Errors](#errors)
- [Generics & Type Parameters](#generics--type-parameters)
- [Type Constraints](#type-constraints)
- [Goroutines](#goroutines)
- [Channels](#channels)
- [Buffered Channels](#buffered-channels)
- [Select](#select)
- [sync.WaitGroup](#syncwaitgroup)
- [sync.Mutex — Protecting Shared Data](#syncmutex--protecting-shared-data)
- [defer, panic & recover](#defer-panic--recover)

---

## 🟢 Structs

Before diving into methods, you need to know about structs. A **struct** is a collection of fields grouped together under one name — Go's way of modeling a real-world object.

```go
type Vertex struct {
    X float64
    Y float64
}
```

Creating and using a struct:

```go
v := Vertex{X: 3.0, Y: 4.0}
fmt.Println(v.X, v.Y)  // access fields with a dot
```

You can also use a pointer to a struct. Go lets you access fields through the pointer without explicitly dereferencing it:

```go
p := &v
p.X = 10  // same as (*p).X = 10 — Go handles this for you
```

> 🟡 Structs are value types in Go. When you assign one struct to another variable, you get a full copy. Use a pointer if you want both variables to point at the same data.

---

## 🟢 Methods

Go does not have classes, but you can define **methods** on any type you define. A method is just a function with a special extra argument called the **receiver**, which appears between the `func` keyword and the method name.

```go
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

v := Vertex{3, 4}
fmt.Println(v.Abs())  // 5
```

Here, `v Vertex` is the receiver — it says "this method belongs to the `Vertex` type, and inside the method, the value is called `v`."

You can define methods on any named type, not just structs:

```go
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
    return float64(c)*9/5 + 32
}

temp := Celsius(100)
fmt.Println(temp.ToFahrenheit())  // 212
```

> 🟡 You can only define methods on types that are declared in the same package. You cannot add methods to built-in types like `int` or `string` directly — but you can create a named type based on them, as shown above.

---

## 🟡 Pointer vs Value Receivers

There are two kinds of receivers: **value receivers** and **pointer receivers**.

**Value receiver** — the method gets a copy of the value. Changes inside the method do not affect the original:

```go
func (v Vertex) Scale(factor float64) {
    v.X *= factor  // only changes the local copy
    v.Y *= factor
}
```

**Pointer receiver** — the method gets a pointer to the original. Changes inside the method affect the original:

```go
func (v *Vertex) Scale(factor float64) {
    v.X *= factor  // changes the actual Vertex
    v.Y *= factor
}
```

Use a pointer receiver when:

- 🟢 You need the method to modify the value
- 🟢 The struct is large and you want to avoid copying it on every call

Use a value receiver when:

- 🟢 The method only reads data and does not need to change anything
- 🟢 The type is small (like a simple number type)

> 🔴 For consistency, if any method on a type uses a pointer receiver, it is good practice to make all methods on that type use pointer receivers too.

---

## 🟢 Method & Pointer Indirection

One of the nicest things Go does for you is handle the conversion between values and pointers automatically when calling methods.

If a method requires a **pointer receiver** but you call it on a plain value, Go automatically takes the address for you:

```go
type MyType struct{ Val int }

func (m *MyType) Double() { m.Val *= 2 }

v := MyType{Val: 5}
v.Double()  // Go rewrites this as (&v).Double() automatically
```

The reverse works too — if a method requires a **value receiver** but you have a pointer, Go automatically dereferences it:

```go
func (m MyType) Print() { fmt.Println(m.Val) }

p := &MyType{Val: 10}
p.Print()  // Go rewrites this as (*p).Print() automatically
```

> 🟡 This convenience only works with methods. For regular function parameters, you still have to pass the right type explicitly — Go will not auto-convert for you there.

---

## 🟢 Interfaces

An **interface** is a set of method signatures. Any type that implements all those methods automatically satisfies the interface — no explicit declaration needed. This is called **implicit implementation** and it is one of Go's best features.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

Now any type that has both `Area()` and `Perimeter()` methods automatically is a `Shape`:

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }
```

Both `Rectangle` and `Circle` satisfy `Shape` without ever saying so. You can now write functions that work with any `Shape`:

```go
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

printShapeInfo(Rectangle{Width: 5, Height: 3})
printShapeInfo(Circle{Radius: 4})
```

> 🟢 Interfaces are Go's primary tool for writing flexible, reusable code. They let you write functions and data structures that work with anything that behaves the right way — not just specific types.

---

## 🟡 The Empty Interface

The empty interface `interface{}` (or `any` in modern Go) has no methods, so every type satisfies it. It is Go's way of saying "I can hold any value."

```go
func printAnything(v interface{}) {
    fmt.Println(v)
}

printAnything(42)
printAnything("hello")
printAnything(true)
```

In Go 1.18+, `any` is an alias for `interface{}` and is preferred for readability:

```go
func printAnything(v any) {
    fmt.Println(v)
}
```

> 🔴 Use the empty interface sparingly. When you use it, you lose type safety and need type assertions to get a useful value back. Prefer concrete types or typed interfaces whenever possible.

---

## 🟢 Type Assertions & Type Switches

Since an interface can hold any type, you sometimes need to reach inside and get the actual concrete value back. That is what a **type assertion** does.

```go
var i interface{} = "hello"

s := i.(string)        // assert that i holds a string
fmt.Println(s)         // hello

n, ok := i.(int)       // safe assertion — ok is false if wrong type
fmt.Println(n, ok)     // 0 false
```

> 🔴 Without the `, ok` pattern, a wrong type assertion causes a panic. Always use the safe form when you are not certain of the type.

When you have multiple possible types, a **type switch** is cleaner than chaining assertions:

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("int: %d\n", v)
    case string:
        fmt.Printf("string: %q\n", v)
    case bool:
        fmt.Printf("bool: %t\n", v)
    default:
        fmt.Printf("unknown type: %T\n", v)
    }
}
```

---

## 🟡 Stringer — Implementing fmt.Stringer

One of the most commonly implemented interfaces in Go is `fmt.Stringer`. If your type has a `String() string` method, Go's `fmt` package will use it whenever it needs to print your value.

```go
type Point struct {
    X, Y int
}

func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

p := Point{3, 7}
fmt.Println(p)  // prints: (3, 7)
```

This is the Go equivalent of `toString()` in other languages.

---

## 🟢 Errors

Go does not have exceptions. Instead, functions return an `error` value as their last return value. A `nil` error means success; anything else means something went wrong.

```go
import "errors"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

You can create custom error types by implementing the `error` interface, which only requires one method — `Error() string`:

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}
```

> 🟢 The pattern `if err != nil` will feel repetitive at first, but it makes error handling explicit and easy to trace — you always know exactly where something could go wrong.

---

## 🟢 Generics & Type Parameters

Before Go 1.18, if you wanted a function to work with different types, you had two options: copy and paste it for each type, or use `interface{}` and lose type safety. **Generics** solve this properly.

A **type parameter** is a placeholder for a type, written in square brackets before the function arguments:

```go
func Map[T, U any](slice []T, f func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = f(v)
    }
    return result
}
```

You can call it with any types and Go will figure them out:

```go
nums := []int{1, 2, 3, 4}
doubled := Map(nums, func(n int) int { return n * 2 })
// doubled = [2 4 6 8]

words := []string{"go", "is", "fun"}
lengths := Map(words, func(s string) int { return len(s) })
// lengths = [2 2 3]
```

Generic functions also work for data structures. Here is a simple generic stack:

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    var zero T
    if len(s.items) == 0 {
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}
```

---

## 🟡 Type Constraints

Type parameters can be constrained to only allow certain types. A **constraint** is just an interface that lists which types are acceptable.

```go
type Number interface {
    int | int32 | int64 | float32 | float64
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

fmt.Println(Sum([]int{1, 2, 3}))         // 6
fmt.Println(Sum([]float64{1.1, 2.2}))    // 3.3
```

The `~` prefix means "this type or any type whose underlying type is this":

```go
type Ordered interface {
    ~int | ~float64 | ~string
}
```

The `golang.org/x/exp/constraints` package (and the built-in `cmp` package in Go 1.21+) provides pre-built constraints like `constraints.Ordered` so you do not have to define them yourself.

> 🟡 Generics are best for truly reusable algorithms and data structures. Do not reach for them just to avoid writing two similar functions — sometimes two clear functions are more readable than one generic one.

---

## 🟢 Goroutines

A **goroutine** is a lightweight thread managed by the Go runtime. You launch one by putting the `go` keyword in front of a function call:

```go
func sayHello(name string) {
    fmt.Println("Hello,", name)
}

go sayHello("Alice")  // runs concurrently
go sayHello("Bob")    // also runs concurrently
```

Goroutines are extremely cheap — you can run thousands of them without breaking a sweat. The Go runtime multiplexes them onto a small pool of OS threads for you.

> 🔴 One thing to watch out for: if `main()` returns, all goroutines are killed immediately, even if they have not finished. You need a way to wait for them — that is where channels and `WaitGroup` come in.

Anonymous functions work great with goroutines:

```go
for i := 0; i < 5; i++ {
    i := i  // capture loop variable (important!)
    go func() {
        fmt.Println(i)
    }()
}
```

> 🔴 Notice the `i := i` line. Without it, all goroutines would share the same loop variable and likely print the same value. Capturing it as a new variable in each iteration is the standard fix.

---

## 🟢 Channels

A **channel** is a pipe that lets goroutines communicate with each other safely. You send values in one end and receive them from the other.

```go
ch := make(chan int)  // create a channel that carries int values

go func() {
    ch <- 42  // send 42 into the channel
}()

value := <-ch  // receive from the channel (waits until a value arrives)
fmt.Println(value)  // 42
```

The `<-` arrow shows the direction of data flow — always pointing toward the receiver.

A practical example — running work in the background and collecting results:

```go
func square(n int, ch chan int) {
    ch <- n * n
}

ch := make(chan int)
go square(3, ch)
go square(5, ch)
go square(7, ch)

fmt.Println(<-ch)  // 9  (order may vary)
fmt.Println(<-ch)  // 25
fmt.Println(<-ch)  // 49
```

**Directional channels** — you can restrict a channel to only sending or only receiving, which makes APIs clearer:

```go
func producer(out chan<- int) {  // can only send
    for i := 0; i < 5; i++ {
        out <- i
    }
    close(out)
}

func consumer(in <-chan int) {  // can only receive
    for v := range in {
        fmt.Println(v)
    }
}
```

> 🟢 Use `close(ch)` to signal that no more values will be sent. Receivers can then use `for v := range ch` to loop until the channel is closed. Never close a channel from the receiver side — only the sender should close it.

---

## 🟡 Buffered Channels

By default, a channel blocks the sender until someone receives. A **buffered channel** has a capacity — it lets the sender continue without waiting, up to that limit.

```go
ch := make(chan int, 3)  // buffer of 3

ch <- 1  // does not block
ch <- 2  // does not block
ch <- 3  // does not block
// ch <- 4  // this would block — buffer is full

fmt.Println(<-ch)  // 1
fmt.Println(<-ch)  // 2
fmt.Println(<-ch)  // 3
```

Think of it like a queue with a limited size. The sender blocks only when the buffer is full; the receiver blocks only when the buffer is empty.

> 🟡 Buffered channels are useful for smoothing out bursts of work, but do not use them as a workaround for race conditions. If you are not sure whether to buffer, start with an unbuffered channel.

---

## 🟢 Select

`select` lets a goroutine wait on multiple channel operations at once. It picks whichever case is ready. If multiple are ready, it picks one at random.

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() { ch1 <- "one" }()
go func() { ch2 <- "two" }()

select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
}
```

Adding a `default` case makes `select` non-blocking — it runs immediately if no channel is ready:

```go
select {
case msg := <-ch:
    fmt.Println("Got:", msg)
default:
    fmt.Println("No message ready, moving on")
}
```

A common pattern is using `select` with a `done` channel to gracefully cancel a goroutine:

```go
func worker(jobs <-chan int, done <-chan struct{}) {
    for {
        select {
        case j := <-jobs:
            fmt.Println("Processing job", j)
        case <-done:
            fmt.Println("Worker shutting down")
            return
        }
    }
}
```

---

## 🟡 sync.WaitGroup

A `WaitGroup` is the standard way to wait for a collection of goroutines to finish before moving on.

```go
import "sync"

var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)  // tell the WaitGroup to expect one more goroutine
    go func(id int) {
        defer wg.Done()  // signal when this goroutine is done
        fmt.Printf("Worker %d finished\n", id)
    }(i)
}

wg.Wait()  // block until all goroutines have called Done()
fmt.Println("All workers finished")
```

The three key methods are:

| Method | What it does |
|--------|-------------|
| `wg.Add(n)` | Register `n` goroutines to wait for |
| `wg.Done()` | Mark one goroutine as finished (usually via `defer`) |
| `wg.Wait()` | Block until the counter reaches zero |

> 🔴 Call `wg.Add(1)` before starting the goroutine, not inside it. If you add inside, there is a race condition where `Wait()` might return before `Add()` is called.

---

## 🟢 sync.Mutex — Protecting Shared Data

When multiple goroutines access the same variable, you have a **race condition** — they can overwrite each other's changes unpredictably. A `Mutex` (mutual exclusion lock) prevents that by allowing only one goroutine at a time to access the critical section.

```go
import "sync"

type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()         // acquire the lock
    defer c.mu.Unlock() // release when this function returns
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

Using it with goroutines:

```go
counter := &SafeCounter{}
var wg sync.WaitGroup

for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        counter.Increment()
    }()
}

wg.Wait()
fmt.Println(counter.Value())  // always 1000
```

> 🟡 Go also has a `sync.RWMutex` for situations where many goroutines read but only a few write. Readers can hold the lock simultaneously; writers get exclusive access. Use `RLock()`/`RUnlock()` for reads and `Lock()`/`Unlock()` for writes.

> 🟢 A good mental model: channels are for communication between goroutines ("send this data over there"), while mutexes are for protecting state shared in memory ("nobody else touch this while I am working on it"). Prefer channels when you can.

---

## 🟡 defer, panic & recover

**defer** schedules a function call to run just before the surrounding function returns — no matter how it returns. It is perfect for cleanup:

```go
func readFile(path string) {
    f, err := os.Open(path)
    if err != nil {
        return
    }
    defer f.Close()  // this runs when readFile returns, guaranteed

    // ... read the file
}
```

Multiple defers run in **last-in, first-out** order:

```go
func example() {
    defer fmt.Println("first deferred, runs last")
    defer fmt.Println("second deferred, runs second")
    defer fmt.Println("third deferred, runs first")
    fmt.Println("function body")
}
// Output:
// function body
// third deferred, runs first
// second deferred, runs second
// first deferred, runs last
```

**panic** stops normal execution and unwinds the stack (running all deferred functions along the way). It is like an exception, but should only be used for truly unrecoverable situations:

```go
func mustPositive(n int) int {
    if n <= 0 {
        panic("n must be positive")
    }
    return n
}
```

**recover** can catch a panic inside a deferred function and prevent the program from crashing. This is useful in servers where you do not want one bad request to take down everything:

```go
func safeCall(f func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    f()
    return nil
}
```

> 🔴 `panic` and `recover` are not a substitute for proper error handling. Use them only as a last resort. For expected failure cases, always return an `error` instead.

---

You now have a solid understanding of Go's advanced building blocks. The next natural steps are exploring the **standard library** (especially `net/http`, `encoding/json`, and `os`), writing **tests** with the `testing` package, and building real programs. Good luck!

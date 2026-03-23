package main

import "fmt"

type enhanced_var[T any] struct {
    Value   T
    Address *T
    Type    string
}

func (e *enhanced_var[T]) Def() {
    e.Type = fmt.Sprintf("%T", e.Value)
}

func (e *enhanced_var[T]) Locate() {
    e.Address = &e.Value
}

func (e enhanced_var[T]) Describe() {
    fmt.Printf("Address: %p, Type: %s\n", e.Address, e.Type)
}

func main() {
    e_var := enhanced_var[string]{Value: "hello"}

    e_var.Def()
    e_var.Locate()
    e_var.Describe()
}

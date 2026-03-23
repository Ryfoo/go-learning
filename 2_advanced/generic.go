package main 


import (
	"fmt"
)

var (
	arr []int = []int{4, 89, 101, 2, -1}
	x int = 101
)

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1 
}


func main() {
	if i := Index(arr, x); i > 0 {
		fmt.Printf("Elemenet has been found at the index %d\n", i)
	}else {
		fmt.Println("the element hasn't been found")
	}
}

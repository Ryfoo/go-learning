package main


import "fmt"



//to declare dynamic array, make() should be used as :
//make([]T, length, capacity)

func main() {
	b := make([]int, 10, 10) //creates a dynamic array of integers with length 10 and capacity 10, returns reference of type slice to b
	fmt.Println(b)
}


package main 

import "fmt"


func linear_transform(x int, a int, b int) int {
	return a * x + b
}
// nake return is : naming the returned values such as a,c with using return without any arguments
//to add code clarity (not for long and complicated functions
func adjacents(b int) (a, c int){
	a = b - 1
	c = b + 1
	return 
}
//var can be used to declare a list of variables
var debug, open_source, availabe bool
//initializers can be used to declared varibles
//note that the variable will take the type of the initializer by default

var i, j int = 1, 2
//or simply var i, j = 1, 2

//the warlus operator := can be used for implicit var definition 
//it can be only used inside a function 
//as the package level requires keywords such as, func and var

func multiplier(x int) int{
	k := 10 //here k is assigned as int
	return k * x
}
func main() {
	fmt.Println(linear_transform(10, 10, 10))
}



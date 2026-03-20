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

func multiplier(x int) int{
	k := 10 //here k is assigned as int
	return k * x
}
func main() {
	fmt.Println(linear_transform(10, 10, 10))
}



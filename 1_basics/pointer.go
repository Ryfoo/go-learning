package main


import "fmt"


//Pointers are like C/C++ Pointers except Go doesn't include pointer arithmetic

//showcasing of the difference between passing-by-value and
//passing-by-reference

func pass_value(x int) {
	x = 100
	fmt.Println("the value x has been changed")
}
func pass_reference(p *int) {
	*p = 100
	fmt.Println("the value x has been chabged")
}

func main() {
	x := 10
	p := &x
	pass_value(x)
	fmt.Println("changing the value of x with the first function : %d", x)
	pass_reference(p)
	fmt.Println("changing the value of x with the second functuon : %d", x)
	//only the second function changes the values stored at the memory address of x. 

}

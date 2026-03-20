package main

import (
	"fmt"
	"math"
)

//Go uses 3 statements for (for) loops
//for statement1;statement2;statement3{}
//note that Go doesn't use () for the statements

//Go's only loop is the for loop 
//the while loop can be implemented with omitting the first and third statements
func loop(x int){
	i := 0
	for ;i < x; i++{
		fmt.Println(i)
	}
}

//Go's control flow can be controlled with the if statements
//further more, a variable can be defined only on the internal if-else block

//the following functions checks if 2^n is greater than 100
//for a given n
func exp_greater(n float64) bool {
	if exp:=math.Pow(2,n); exp > 100{
		return true
	} else{
		return false
	}

}

//in the function exp_greater the else-block can be dropped (omitted)
//but sometimes it adds clarity
//note that exp is only defined inside the if-else block
//as calling it outside that block will lead to an error


func main(){
	var x, n = 10, float64(10)
	loop(x)

	if check:=exp_greater(n); check{
		fmt.Println("True")
	} else{
		fmt.Println("False")
	}

}



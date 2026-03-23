package main
import "fmt"


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
	k := 10
	return k * x
}
//constants are declared with the const keyword
//constants cannot be declared using :=

const Pi = 3.14


func main(){


}

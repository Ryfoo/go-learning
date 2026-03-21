package main


import (
	"fmt"
	"math"
)



func main() {
	//arrays are declared as [n]T where n : is the number of elements
	//T is the type of elements
	//Go doesn't accept mixed types elements in one array
	//and the number of elements should be declared, since it defines the type of of the array
	//arrays cannot be resized unless using special Go methods and utils
	// [i]string isn't the same type as [j]string where i != j

	var a [10]int

	for i := 0; i < 10 ; i++ {
		a[i] = int(math.Pow(float64(i), 2))
	}

	fmt.Println(a)

	//slices are defined as var slice []T = array[low : high]
	//note that slices aren't independent copies, they are just a reference to the array from low to (but not including) high
	//sizes don't matter with slices, since they just hold the address of the according memory chunk (low -> high)
	//but they don't hold a segment of memory itself

	var s []int = a[0:8]
	//or simply s := a[0:8]

	fmt.Println(s)
	//low or high can be omitted, and the lowest or highest values will be assumed, [:] -> [lowest:highest], [low:] -> [low:highest]

}

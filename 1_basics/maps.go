package main


import "fmt"


//maps can be initialized as var map[T1]T2, where T1 is the type of the key, while T2 is the type of the element


//map literals, map[T1]T2 := {key1 : elem2, ...}

//can be declared with make() such as, m := make(map[T1]T2) 


func main() {
	m := make(map[string]int)

	//insert or modify an element
	m["version"] = 12
	m["debug"] = 1
	m["users"] = 10000
	//delete an element
	delete(m, "version")


	for i := range m {
		fmt.Println(m[i])
	}
	
}

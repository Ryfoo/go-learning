package main


import "fmt"


//Notice one of the most known Go's type conventions
//struct's fields that starts with capital letters are public while others are private

//pointers to struct are a bit speical
//when modifying or referring to a field, the pointer doesn't have to be dereferenced
//one of Go's conventions for simplicity
//meaning (*pointer_to_struct).field = pointer_to_struct.field

type Vertex struct {
	X int
	Y int
	private string 
}


func main() {
	a := Vertex{4, 3, "no one has to know"}
	p := &a
	//notice how both the next lines do the same job
	p.X = 10
	fmt.Println(p.X)
	(*p).X = 20
	fmt.Println(p.X)

	//that line will raise an error
	//fmt.Print(a.private)


	//non-initialized values will take zero-values (0, "", false)
	b := Vertex { X : 1, Y: 2} // private has value ""
	fmt.Print(b)
}

package main 



import "fmt"
var Debug bool = true


type PP struct {
	X, Y, Z float64
	Material string
}

func (v PP) Describe() {
	fmt.Println("New Object Characterstics")
	fmt.Printf("width : %.2f\n height : %.2f\n depth : %.2f\n material : %s\n", v.X, v.Y, v.Z, v.Material)
}
func (p *PP) Scale(f float64) {
	if Debug {
		
		fmt.Printf("%p\n", p)

	}
	p.X *= f
	p.Y *= f
	p.Z *= f
}

func main() {
	v := PP {3, 4, 2, "wood"}
	if Debug {
		fmt.Printf("%p\n", &v)
	}
	p := &v
	v.Describe()
	p.Scale(10)
	v.Describe()

}

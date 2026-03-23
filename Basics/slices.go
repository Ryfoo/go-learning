package main


import (
	"fmt"
	"math/rand"
	"os"
)

func Pic(dx, dy int) [][]uint8 {
	var matrix [][]uint8
	for i := 0; i < dy; i++{
		var s []uint8
		for j := 0; j < dx; j++ {
			bluescale_value := uint8(rand.Intn(256))
			s = append(s, bluescale_value)
		}
		matrix = append(matrix, s)

	}
	return matrix
}


func main() {
	photo, err := os.Open("photo.png")
	if err != nil {
		fmt.Println("error happened, access to the file cannot be granted")
		return
	}
	defer photo.Close()


	photo.Show(Pic)
}

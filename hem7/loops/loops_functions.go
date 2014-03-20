package main

/*
Author: Daniel Cserhalmi
Version 1.0
*/
import (
	"fmt"
	"math"
)

const delta = 0.0001

func Sqrt(x float64) float64 {
	z := x
	for i := 0; i < 14; i++ {
		temp := z
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-temp) < delta {
			fmt.Println(i)
			return math.Abs(z)
		}
	}
	return math.Abs(z)
}

func main() {
	const x = 4
	my, gos := Sqrt(x), math.Sqrt(x)
	fmt.Println("my: ", my, "go's: ", gos, "difference: ", my-gos)

}

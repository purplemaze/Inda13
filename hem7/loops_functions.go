package main

/*
Author: Daniel Cserhalmi
Version 1.0
*/
import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	delta := 0.5
	z := 1.0
	for i := 0; i < 14; i++ {
		temp := z
		z = z - (math.Pow(z, 2)-x)/2*z
		if math.Abs(z-temp) < delta {
			fmt.Println(i)
			return math.Abs(z)
		}
	}
	return math.Abs(z)
}

func main() {
	fmt.Println(Sqrt(4))
}

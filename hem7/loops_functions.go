package hem7

/*
Author: Daniel Cserhalmi
Version 1.0
*/
import (
	"fmt"
	"math"
)

/*
Like for, the if statement can start with a short statement to execute before the condition.
Variables declared by the statement are only in scope until the end of the if.
*/
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		pow(3, 3, 30),
	)
}

/*
Summa i två delar

Gör klart programmet nedan. Det adderar alla tal i en vektor genom att dela vektorn på mitten och låta två gorutiner göra halva jobbet var.
Delresultaten skickas över en kanal. Lämna in koden för hela programmet. Glöm inte "go fmt".
*/

package main

import "fmt"

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan<- int) {
	sum := 0
	for _, s := range a {
		sum += s
	}
	res <- sum
	
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)
	sum := <-ch + <-ch
	close(ch)
	fmt.Println(sum)
}

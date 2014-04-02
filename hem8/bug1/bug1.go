// I want this program to print "Hello world!", but it doesn't work.

// Eftersom ingen annan gorutine tar emot det som skickas på kanalen ch, man kan inte skicka data till "sig själv" via en kanal,
//kommer det uppstå deadlock. Om man skapar en ny gorutine som tar emot det som skickas på kanalen fungerar det
package main

import "fmt"

func printS(s <-chan string) {
	fmt.Println(<-s)
}

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)
	go printS(ch)
	ch <- "Hello world!"
}

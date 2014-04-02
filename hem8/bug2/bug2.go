// This program should go to 11, but sometimes it only prints 1 to 10.

//Eftersom det inte går att veta vilken tråd som blir "klar" först kan det hända att kanlen stängs innan Print() 
//funktionen hunnit skriva ut alla heltal. 
//Om vi använder funktionen WaitGroup från paketet sync kan vi se till att main rutinen väntar på övriga rutiner, i detta fall bara en.
package main

import ( 
	"fmt"
	"sync"
	)

 var wg sync.WaitGroup

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
    for n := range ch { // reads from channel until it's closed
        fmt.Println(n)
    }
    wg.Done()
}


func main() {
    ch := make(chan int)
    go Print(ch)
    wg.Add(1)	//lägger till en gorutin i gruppen
    for i := 1; i <= 11; i++ {
        ch <- i
    }
    close(ch)
    wg.Wait() //väntar på att alla gorutiner ska hinna klart


}
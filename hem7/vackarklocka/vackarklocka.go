//Påminnelse
package main

import (
	"fmt"
	"time"
)

const layout = "15:04:05" // Format returns a textual representation of the time value formatted according to layout, which defines the format by showing how the reference time, Mon Jan 2 15:04:05 -0700 MST 2006, would be displayed if it were the value; it serves as an example of the desired output.

//Skriver ut en påminnelse om och om igen med en paus av given längd före varje påminnelse
func Remind(text string, paus time.Duration) {
	for { // infinite loop
		t := time.Now()
		fmt.Println("Klockan är", t.Format(layout), ": ", text)
		time.Sleep(paus)
	}
}

func main() {
	Remind("Gör klart veckans Inda", 5*time.Second)
}

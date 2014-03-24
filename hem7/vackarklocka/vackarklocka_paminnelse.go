//Äta-jobba-sovklocka 
package main

import (
	"fmt"
	"time"
)

//konstaten hour är satt till time.Second för att enklare testa programmet, det ska egentligen vara time.Hour. 
//Även layout är satt till "XX:XX:XX"
const hour, layout, eat, work, sleep = time.Second, "15:04:05", 3 * hour, 8 * hour, 24 * hour

//Skriver ut en påminnelse:
//var 3:e timme: "Klockan är XX.XX: Dags att äta",
//var 8:e timme: "Klockan är XX.XX: Dags att arbeta",
//var 24:e timme:"Klockan är XX.XX: Dags att sova".
//
func Remind() {
	go func() {
		for { // infinite loop
			time.Sleep(eat)
			t := time.Now()
			fmt.Println("Klockan är", t.Format(layout), ": ", "Dags att äta")
		}
	}()
	go func() {
		for { // infinite loop
			time.Sleep(work)
			t := time.Now()
			fmt.Println("Klockan är", t.Format(layout), ": ", "Dags att arbeta")
		}
	}()
	go func() {
		for { // infinite loop
			time.Sleep(sleep)
			t := time.Now()
			fmt.Println("Klockan är", t.Format(layout), ": ", "Dags att sova")
		}
	}()

}

func main() {
	Remind()
	select {}
}

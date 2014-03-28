// Stefan Nilsson 2013-03-13
// Daniel Cserhalmi 2014-03-28
// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
/*
 Filen oracle.go innehåller ett kodskelett till ett orakelprogram som besvarar frågor.

    Gör klart Oracle-metoden. Du får inte ändra i main-metoden och du får inte heller ändra metodsignaturerna. 
    Observera att svaren inte ska komma direkt, utan med fördröjning. 
    Glöm inte heller att oraklet ska skriva ut meddelanden även om det inte kommer några frågor. 
    Du får gärna dela upp din lösning på flera metoder.

Ditt program ska innehålla två stycken kanaler: en kanal för frågor samt en kanal för svar och förutsägelser.
 I Oracle-metoden ska du starta tre stycken permanenta gorutiner:

    En gorutin som tar emot alla frågor och för varje inkommande fråga skapar en separat gorutin som besvarar frågan.
    En gorutin som genererar förutsägelser.
    En gorutin som tar emot alla svar och förutsägelser och skriver ut dem på stdout.

Oracle-metoden är den viktigaste delen av uppgiften. Om du vill får du också förbättra svarsalgoritmen. 
Även här får gärna dela upp algoritmen på flera metoder. Här är några tips:

    Paketen strings och regexp kan vara användbara.
    Programmet kan verka mera mänskligt om oraklet skriver ut sina svar en bokstav i taget.
    Ta en titt på ELIZA, det första programmet av det här slaget.

*/
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.
	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}

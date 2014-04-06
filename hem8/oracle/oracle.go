// Stefan Nilsson 2013-03-13
// Daniel Cserhalmi 2014-04-01
// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
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

var answerMap = make(map[string]string)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.ToLower(strings.TrimSpace(line))
		if line == "" {
			continue
		}
		if line == "quit" { //exit from the program
			break
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
	answers := make(chan string)
	fillAnswers()
	go answerBuffer(questions, answers)
	go print(answers)
	go prophecy("", answers)
	return questions
}

//This is the answer buffer function
//It recives all the questions and creates a separate answer gorutine for all of them.
func answerBuffer(questions <-chan string, answers chan string) {
	for s := range questions {
		go answer(s, answers)
	}
}

//This is the oracle's answer algorithm.
//It waits for a while and then sends and answer on the answer channel
func answer(question string, answers chan<- string) {
	time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
	//todo fix the answering algorithm
	s := make([]string, 2)
	s[0] = "That's a silly question."
	s[1] = "That's not even a question.. you are wasting my time."

	answer := s[rand.Intn(len(s))] // default answer
	for i := range answerMap {
		if strings.Contains(question, i) {
			answer = answerMap[i]
		}
	}
	answers <- answer

}

//This is the print function.
//It prints the strings it recives on the answer channel one character at a
//time, "random time", to simulate a real person.
func print(ch <-chan string) {
	for s := range ch {
		for _, s := range strings.Split(s, "") { //splits the strings by ""
			time.Sleep(time.Duration(50+rand.Intn(150)) * time.Millisecond)
			fmt.Print(s)
		}
		fmt.Println("")
		fmt.Print(prompt)
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	for {
		time.Sleep(time.Duration(30+rand.Intn(20)) * time.Second)
		// Quotes from the Oracle in the Matrix
		nonsense := []string{
			"We are all here to do what we are all here to do...",
			"It seems that every time we meet, I have nothing but bad news. I'm sorry about that, I surely am. But for what it's worth, you've made a believer out of me. Good luck, kiddo.",
			"What do all men with power want? More power.",
			"Everything that has a beginning has an end. I see the end coming, I see the darkness spreading. I see death.",
			"I'll make the predictions around here.",
		}
		answer <- "... " + nonsense[rand.Intn(len(nonsense))]
	}
}

//Fills the map answerMap with answers
func fillAnswers() {
	answerMap["color"] = "The color you percive is not of importance.. "
	answerMap["how"] = "How do people do anything?"
	answerMap["god"] = "God is, even though the whole world deny him. Truth stands, even if there be no public support. It is self-sustained."
	answerMap["hello"] = "Don't waste my time.. what do you want?"
	answerMap["help"] = "I can't help you with that.."
	answerMap["art"] = "Art for art’s sake makes no more sense than gin for gin’s sake"
	answerMap["should"] = "The desires of our ego are often in conflict with the emotions of our heart.  You’ll always have what you want, if you want what you have"
	answerMap["death"] = "We so easily lose perspective on what takes up our energy and focus.  We’re all dying.  Sometimes we need to remind ourselves of this to enjoy living."
	answerMap["meaning"] = "42, the number 42 is the answer to The Ultimate Question of Life, the Universe, and Everything"
	answerMap["you"] = "Me?.. I thought we were talking about you"
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}

package main

/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

Author: Daniel Cserhalmi
Version 1.0
*/
import (
	"strings"
	"code.google.com/p/go-tour/wc"
)

var m map[string]int

func WordCount(s string) map[string]int {
	words := strings.Fields(s) //Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning an array of substrings of s or an empty list if s contains only white space.
	m = make(map[string]int)   //Maps must be created with make (not new) before use; the nil map is empty and cannot be assigned to.
	if words != nil {
		for _, i := range words {
			if m[i] != 0 {
				m[i]++
			} else {
				m[i] = 1
			}
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

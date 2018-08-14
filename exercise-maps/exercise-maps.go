package main

//Implement WordCount. It should return a map of the counts of each “word” 
//in the string s. The wc.Test function runs a test suite against the provided 
//function and prints success or failure.

//You might find strings.Fields helpful.

import (
	"golang.org/x/tour/wc"
	"strings"

)

func WordCount(s string) map[string]int {

	m := make(map[string]int)
	word_list := strings.Fields(s)

	//by default range just gives the index
	for i:= range(word_list){
		elem := m[word_list[i]]
		m[word_list[i]]= elem+1 //default elem is 0 as it is int type
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
package main

import (
	"fmt"
	"strings"
	//"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		ele, present := m[word]
		if present {
			m[word] = ele + 1
		} else {
			m[word] = 1
		}
		//fmt.Println(i, word)
	}
	return m
}

func main() {
	//wc.Test(WordCount)
	var input string
	//fmt.Println("Enter the input string")
	//fmt.Scanln(&input)
	
	//Enter input string below
	input="Part of the journey is the end"
	my_map := WordCount(input)

	for k := range my_map {
		fmt.Printf("%s : %d\n", k, my_map[k])
	}
}

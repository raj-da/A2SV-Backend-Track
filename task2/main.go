package main

import (
	"fmt"
	"strings"
)

func frequncyCount(input string) map[string]int {
	frequncy := map[string]int{}
	input = strings.ToLower(input)
	words := strings.Split(input, " ")

	for _, word := range words {
		frequncy[word]++
	} 

	return frequncy
}

func main() {
	input := "In Go characters are runes When you iterate over a string using range you get runes"
	fmt.Println("Word Frequncy: ", frequncyCount(input))
}
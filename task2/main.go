package main

import (
	"fmt"
	"strings"
	"unicode"
)

func FrequncyCount(input string) map[string]int {
	frequncy := map[string]int{}
	input = strings.ToLower(input)
	words := strings.Split(input, " ")

	// Filter out non-alphanumeric character
	filteredWord := []string{}
	for _, word := range words {
		lastIndex := len(word) - 1
		lastChar := word[lastIndex]
		if (!unicode.IsLetter(rune(lastChar)) && !unicode.IsDigit(rune(lastChar))) {
			filteredWord = append(filteredWord, word[:lastIndex])
		} else {
			filteredWord = append(filteredWord, word)
		}
	}

	for _, word := range filteredWord {
		frequncy[word]++
	} 

	return frequncy
}

func IsPalindrome(s string) bool {
	// Filter out non-alphanumeric characters
	var runes []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			runes = append(runes, unicode.ToLower(r))
		}
	}

	// Check if palindrom
	l := 0
	r := len(runes) - 1
	for l < r {
		if runes[l] != runes[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func main() {
	input := "In Go, characters are runes When you iterate over a string using range you get runes"
	fmt.Println("Word Frequncy: ", FrequncyCount(input))

	testCases := []string{
		"A man, a plan, a canal: Panama", // True
		"race a car",                     // False
		"Was it a car or a cat I saw?",   // True
		"No 'x' in Nixon",                // True
		"12321",                          // True
		"Hello World",                    // False
	}

	for _, tc := range testCases {
		fmt.Printf("Input: %-32q | Palindrome: %v\n", tc, IsPalindrome(tc))
	}
}
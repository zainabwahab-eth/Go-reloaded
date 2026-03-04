package operations

import (
	"fmt"
	"strings"
)

func CheckPunctuation(index int, inputSlice []string, mapCheck map[int]bool) ([]string, string) {
	punctuation := ".,!?;:"

	word := inputSlice[index]

	if index <= 0 || index >= len(inputSlice) {
		return inputSlice, word
	}

	in := 0

	for in < len(word) && strings.ContainsRune(punctuation, rune(word[in])) {
		in++
	}

	if in <= 0 {
		return inputSlice, word
	}

	previous := index - 1
	if mapCheck[index-1] == true {
		previous = index - 2
	}

	fmt.Println(word)
	//add punctuation to previous word
	inputSlice[previous] += word[:in]

	//remove punctuation from current word
	word = word[in:]

	inputSlice[index] = word
	fmt.Println("result", inputSlice, word)

	return inputSlice, word

}

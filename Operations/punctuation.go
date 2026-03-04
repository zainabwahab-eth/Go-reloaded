package operations

import (
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

	// loop back to find the first non-deleted word
	previous := index - 1
	for previous > 0 && mapCheck[previous] {
		previous--
	}

	inputSlice[previous] += word[:in]
	word = word[in:]
	inputSlice[index] = word

	return inputSlice, word
}

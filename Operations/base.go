package operations

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Base(sentence string) string {
	stringToDel := make(map[int]bool)
	sentenceSlice := strings.Fields(sentence)
	punctuation := ".,!?;:"

	for index, word := range sentenceSlice {
		switch {
		case strings.Contains(word, "(hex)"):
			word = ConvertBase(sentenceSlice, index, 16)
			sentenceSlice[index-1] = word
			stringToDel[index] = true

		case strings.Contains(word, "(bin)"):
			word = ConvertBase(sentenceSlice, index, 2)
			sentenceSlice[index-1] = word
			stringToDel[index] = true

		case strings.Contains(word, "(up)"):
			sentenceSlice[index-1] = strings.ToUpper(sentenceSlice[index-1])
			stringToDel[index] = true

		case strings.Contains(word, "(low)"):
			sentenceSlice[index-1] = strings.ToLower(sentenceSlice[index-1])
			stringToDel[index] = true

		case strings.Contains(word, "(cap)"):
			sentenceSlice[index-1] = strings.ToUpper(string(sentenceSlice[index-1][0])) + string(sentenceSlice[index-1][1:])
			stringToDel[index] = true

		case strings.Contains(word, ")") && (strings.Contains(sentenceSlice[index-1], "(up") || strings.Contains(sentenceSlice[index-1], "(low") || strings.Contains(sentenceSlice[index-1], "(cap")):
			re := regexp.MustCompile(`\d+`)
			matches := re.FindAllString(word, -1)

			//convert to number
			num, err := strconv.Atoi(matches[0])
			if err != nil {
				fmt.Println("Atoi Error:", err)
			}
			sentenceSlice = DoubleOperation(index, num, sentenceSlice)
			stringToDel[index] = true
			stringToDel[index-1] = true

		case (word == "a" || word == "A") && index != len(sentenceSlice)-1:
			if checkVowel(index, sentenceSlice) {
				sentenceSlice[index] = sentenceSlice[index] + "n"
			}

		case strings.Contains(punctuation, string(word[0])):
			sentenceSlice, word = CheckPunctuation(index, sentenceSlice, stringToDel)
			if word == "" {
				stringToDel[index] = true
			}
		}

	}

	parsed := []string{}
	for i, v := range sentenceSlice {
		if !stringToDel[i] {
			parsed = append(parsed, v)
		}

	}
	result := strings.Join(parsed, " ")
	result = FixQuotes(result)
	return result

}

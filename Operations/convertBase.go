package operations

import (
	"fmt"
	"strconv"
)

func ConvertBase(inputSlice []string, index int, base int) string {
	word := inputSlice[index-1]
	dec, err3 := strconv.ParseInt(word, base, 62)
	if err3 != nil {
		fmt.Println("Error:", err3)
	}
	inputSlice[index-1] = strconv.Itoa(int(dec))
	return inputSlice[index-1]
}

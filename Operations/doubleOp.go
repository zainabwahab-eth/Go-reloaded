package operations

import "strings"

func DoubleOperation(index int, num int, inputSlice []string) []string {
	switch inputSlice[index-1] {
	case "(up,":
		for i := index - 2; i > index-2-num; i-- {
			inputSlice[i] = strings.ToUpper(inputSlice[i])
		}
	case "(low,":
		for i := index - 2; i > index-2-num; i-- {
			inputSlice[i] = strings.ToLower(inputSlice[i])
		}
	case "(cap,":
		for i := index - 2; i > index-2-num; i-- {
			inputSlice[i] = strings.ToUpper(string(inputSlice[i][0])) + string(inputSlice[i][1:])
		}
	}
	return inputSlice
}

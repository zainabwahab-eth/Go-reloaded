package operations

import (
	"fmt"
	"os"
)

func ReadFile() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Please enter input and output file name")
	}

	inputFile := args[1]
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error", err)
	}

	res := Base(string(inputData))
	WriteFile(res, args[2])
}

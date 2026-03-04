package operations

import (
	"log"
	"os"
)

func WriteFile(data string, submFile string) {
	dataByte := []byte(data)
	err := os.WriteFile(submFile, dataByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

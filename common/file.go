package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FileContents(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	contents := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return contents
}

package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1Part1() {
	fmt.Println("Running Solution for Day 1!")

	file, err := os.Open("day1/input.txt")
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

	ints := []int{}

	for i := 0; i < len(contents); i++ {
		convertedInt, err := strconv.Atoi(contents[i])
		if err != nil {
			log.Fatalf("Expected file to have only integers")
		}
		ints = append(ints, convertedInt)
	}

	numIncreased := 0
	for i := 1; i < len(ints); i++ {
		if ints[i] > ints[i-1] {
			numIncreased++
		}
	}

	log.Printf("Number of increases %+v", numIncreased)
}

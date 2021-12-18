package main

import (
	"log"
	"os"
	"strconv"

	"github.com/dalecb13/advent-of-code-2021/day1"
)

func main() {
	numArgs := len(os.Args)
	if numArgs != 2 {
		log.Fatalf("Argument must be defined")
	}

	day := os.Args[1]

	dayInt, err := strconv.Atoi(day)
	if err != nil {
		log.Fatalf("Argument must be an integer")
	}

	switch dayInt {
	case 1:
		day1.Day1Part1()
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dalecb13/advent-of-code-2021/day1"
	"github.com/dalecb13/advent-of-code-2021/day2"
)

func main() {
	numArgs := len(os.Args)
	if numArgs != 3 {
		log.Fatalf("Argument must be defined")
	}

	var day, part int

	flag.IntVar(&day, "d", 1, "Day of the Advent Calendar")
	flag.IntVar(&part, "p", 1, "Part (1 or 2)")
	flag.Parse()

	if day < 1 || day > 32 {
		log.Fatalf("Day must be between 1 and 31, inclusive")
	}
	if part != 1 && part != 2 {
		log.Fatalf("Part must be 1 or 2")
	}

	switch day {
	case 1:
		if part == 1 {
			day1.Day1Part1()
		} else {
			day1.Day1Part2()
		}
	case 2:
		if part == 1 {
			day2.Day2Part1()
		} else {
			log.Fatal("Day 2 Part 2 solution not yet ready")
		}
	default:
		fmt.Println("No solution available for that day.")
	}
}

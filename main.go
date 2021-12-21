package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dalecb13/advent-of-code-2021/day1"
	"github.com/dalecb13/advent-of-code-2021/day2"
	"github.com/dalecb13/advent-of-code-2021/day3"
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
			day1.Part1()
		} else {
			day1.Part2()
		}
	case 2:
		if part == 1 {
			day2.Part1("day2/input.txt")
		} else {
			product := day2.Part2("day2/input.txt")
			log.Printf("Product is %+v", product)
		}
	case 3:
		if part == 1 {
			powerConsumption := day3.Part1("day3/input.txt")
			log.Printf("powerConsumption is %+v", powerConsumption)
		} else {
			lifeSupportRating := day3.Part2("day3/input.txt")
			log.Printf("lifeSupportRating is %+v", lifeSupportRating)
		}
	default:
		fmt.Println("No solution available for that day.")
	}
}

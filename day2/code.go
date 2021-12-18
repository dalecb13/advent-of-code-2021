package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2Part1() {
	fmt.Println("Running Solution for Day 2 Part 1!")

	file, err := os.Open("day2/input.txt")
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

	/*
		Initialize two variables to store values
	*/
	depth := 0
	horizontal := 0
	// iterate through `contents` and add/subtract distances
	for i := 0; i < len(contents); i++ {
		parts := strings.Fields(contents[i])
		if parts[0] == "forward" {
			horizontalDistance, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalf("Expected second rune to be an integer")
			}
			horizontal += horizontalDistance
		} else {
			if parts[0] == "down" {
				increaseDepth, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Fatalf("Expected second rune to be an integer")
				}
				depth += increaseDepth
			} else if parts[0] == "up" {
				decreaseDepth, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Fatalf("Expected second rune to be an integer")
				}
				depth -= decreaseDepth
			} else {
				log.Fatalf("Expected direction to be one of 'up', 'down', or 'horizontal'")
			}
		}
	}

	log.Printf("Product of distance %+v and depth %+v is %+v", horizontal, depth, horizontal*depth)
}

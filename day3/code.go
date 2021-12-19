package day3

import (
	"fmt"

	"github.com/dalecb13/advent-of-code-2021/common"
)

// RuneFrequencies counts the number of each character in a string
func RuneFrequencies(str string) map[rune]int {
	frequencies := map[rune]int{}
	// a string in Golang is a UTF-8 sequence of bytes
	// a for-range loop will correctly get characters
	// converting a string to a rune will treat characters in a string as a characgter
	runeRow := []rune(str)
	for runeIdx := 0; runeIdx < len(runeRow); runeIdx++ {
		char := runeRow[runeIdx]

		// check keys if any of them is the rune
		keys := make([]rune, 0, len(frequencies))
		for k := range frequencies {
			keys = append(keys, k)
		}

		hasRune := false
		for freqIdx := 0; freqIdx < len(keys); freqIdx++ {
			runeKey := rune(keys[freqIdx])
			if runeKey == char {
				hasRune = true
			}
		}

		if hasRune {
			oldFreq := frequencies[char]
			newFreq := oldFreq + 1
			frequencies[char] = newFreq
		} else {
			frequencies[char] = 1
		}
	}

	return frequencies
}

// DiagonalFlipStringSlices treats an array of strings as a matrix of runes and performs a transpose.
// It also assumes that each string is the same length
func TransposeStringSlices(slices []string) []string {
	runeified := [][]rune{}
	for _, s := range slices {
		rs := []rune(s)
		runeified = append(runeified, rs)
	}

	transposeRunes := TransposeRuneMatrix(runeified)

	// convert transposed slices of rune slices to slice of string
	converted := []string{}
	for _, transposed := range transposeRunes {
		convert := string(transposed)
		converted = append(converted, convert)
	}

	return converted
}

func TransposeRuneMatrix(matrix [][]rune) [][]rune {
	transposeds := [][]rune{}

	for rowIdx, row := range matrix {
		for runeIdx, r := range row {
			if rowIdx == 0 {
				// if it's the first index, initialize new rune slice
				rs := []rune{r}
				transposeds = append(transposeds, rs)
			} else {
				// else update existing
				existing := transposeds[runeIdx]
				existing = append(existing, r)
				transposeds[runeIdx] = existing
			}
		}
	}

	return transposeds
}

// Part1 reads a file of binary numbers. The binary numbers create two numbers,
// gammaRate and epsilonRate. powerConsumption = gammaRate * epsilonRate. Bits in the
// gammaRate is the most common bit in each digit, and then converted to decimal. The
// epsilonRate is the least common bit in each digit, converted to decimal.
func Part1(filePath string) int {
	fmt.Println("Running Solution for Day 3 Part 1!")

	fileContents := common.FileContents(filePath)
	transposedContents := TransposeStringSlices(fileContents)

	// holds the counts of zeroes and ones for each column
	indexRuneFrequencies := make(map[int]map[rune]int)
	for freqIdx, transposedContent := range transposedContents {
		frequencies := RuneFrequencies(transposedContent)
		indexRuneFrequencies[freqIdx] = frequencies
	}

	gammaRunes := []rune{}
	epsilonRunes := []rune{}
	zeroRune := []rune("0")[0]
	oneRune := []rune("1")[0]
	for i := 0; i < len(indexRuneFrequencies); i++ {
		frequencies := indexRuneFrequencies[i]

		// each `frequencies` is histogram of the number of 0's and 1's in the given column
		numZeroes := frequencies[zeroRune]
		numOnes := frequencies[oneRune]

		if numZeroes > numOnes {
			// zeroes are more common, so assign 0 to gamma
			gammaRunes = append(gammaRunes, zeroRune)
			epsilonRunes = append(epsilonRunes, oneRune)
		} else {
			// ones are more common, so assign 1 to gamma
			gammaRunes = append(gammaRunes, oneRune)
			epsilonRunes = append(epsilonRunes, zeroRune)
		}
	}

	gammaString := string(gammaRunes)
	epsilonString := string(epsilonRunes)

	gammaRate := common.BinaryStringToDecimal(gammaString)
	epsilonRate := common.BinaryStringToDecimal(epsilonString)

	powerConsumption := int(gammaRate) * int(epsilonRate)

	return powerConsumption
}

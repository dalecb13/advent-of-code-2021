package day3

import (
	"testing"

	"github.com/dalecb13/advent-of-code-2021/common"
)

func TestRuneFrequencies(t *testing.T) {
	input := "test1"
	expectedOutput := make(map[rune]int)
	expectedOutput[[]rune("t")[0]] = 2
	expectedOutput[[]rune("e")[0]] = 1
	expectedOutput[[]rune("s")[0]] = 1
	expectedOutput[[]rune("1")[0]] = 1

	frequencies := RuneFrequencies(input)

	for freqKey := range frequencies {
		hasFreqKey := false
		for expectedKey := range expectedOutput {
			if freqKey == expectedKey {
				hasFreqKey = true
			}
		}

		if !hasFreqKey {
			t.Log("Expected frequency map to have rune")
			t.Fail()
		} else {
			if frequencies[freqKey] != expectedOutput[freqKey] {
				t.Log("Expected frequency value to match")
				t.Fail()
			}
		}
	} // end for-range frequencies
}

func TestTransposeRuneMatrix(t *testing.T) {
	rune2d := [][]rune{
		[]rune("test1"),
		[]rune("test2"),
		[]rune("test3"),
	}
	expectedResult := [][]rune{
		[]rune("ttt"),
		[]rune("eee"),
		[]rune("sss"),
		[]rune("ttt"),
		[]rune("123"),
	}

	actualResult := TransposeRuneMatrix(rune2d)

	if len(expectedResult) != len(actualResult) {
		t.Logf("Expected number of rows to match; len(expectedResult) %+v, len(actualResult) %+v", len(expectedResult), len(actualResult))
		t.Fail()
	}

	for rowIdx, expectedRow := range expectedResult {
		if len(expectedRow) != len(actualResult[rowIdx]) {
			t.Logf("Expected lengths of rune slices to match; len(expectedRow) %+v != len(actualResult[rowIdx] %+v", len(expectedRow), len(actualResult[rowIdx]))
			t.Fail()
		}
		for runeIdx, expectedRune := range expectedRow {
			if expectedRune != actualResult[rowIdx][runeIdx] {
				t.Logf("Expected runes to match")
				t.Fail()
			}
		}
	}
}

func TestTransposedStringSlices(t *testing.T) {
	slices := []string{
		"test1",
		"test2",
		"test3",
	}

	expectedResult := []string{
		"ttt",
		"eee",
		"sss",
		"ttt",
		"123",
	}

	actualResult := TransposeStringSlices(slices)

	if len(expectedResult) != len(actualResult) {
		t.Logf("Expected number of rows to match; len(expectedResult) %+v, len(actualResult) %+v", len(expectedResult), len(actualResult))
		t.Fail()
	}

	for rowIdx, expectedRow := range expectedResult {
		if len(expectedRow) != len(actualResult[rowIdx]) {
			t.Logf("Expected lengths of rune slices to match; len(expectedRow) %+v != len(actualResult[rowIdx] %+v", len(expectedRow), len(actualResult[rowIdx]))
			t.Fail()
		}
		if expectedRow != actualResult[rowIdx] {
			t.Logf("Expected strings to match")
			t.Fail()
		}
	}
}

func TestPart1(t *testing.T) {
	result := Part1("./test_input.txt")
	if result != 198 {
		t.Fatalf("Expected result to be 198")
		t.Fail()
	}
}

func TestFindOxygenGeneratingRate(t *testing.T) {
	input := [][]rune{
		[]rune("00100"),
		[]rune("11110"),
		[]rune("10110"),
		[]rune("10111"),
		[]rune("10101"),
		[]rune("01111"),
		[]rune("00111"),
		[]rune("11100"),
		[]rune("10000"),
		[]rune("11001"),
		[]rune("00010"),
		[]rune("01010"),
	}
	expectedResult := int64(23)

	rate := FindOxygenGeneratingRate(input, 0)
	str := string(rate)
	actualResult := common.BinaryStringToDecimal(str)

	if actualResult != expectedResult {
		t.Fatalf("Expected result to be 23")
		t.Fail()
	}
}

func TestFindCo2ScrubbingRate(t *testing.T) {
	input := [][]rune{
		[]rune("00100"),
		[]rune("11110"),
		[]rune("10110"),
		[]rune("10111"),
		[]rune("10101"),
		[]rune("01111"),
		[]rune("00111"),
		[]rune("11100"),
		[]rune("10000"),
		[]rune("11001"),
		[]rune("00010"),
		[]rune("01010"),
	}
	expectedResult := int64(10)

	rate := FindCo2ScrubbingRate(input, 0)
	str := string(rate)
	actualResult := common.BinaryStringToDecimal(str)

	if actualResult != expectedResult {
		t.Fatalf("Expected result to be 23")
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	result := Part2("./test_input.txt")
	if result != 230 {
		t.Fatalf("Expected result to be 230")
		t.Fail()
	}
}

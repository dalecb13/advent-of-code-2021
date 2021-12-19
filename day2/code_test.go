package day2

import (
	"testing"
)

func TestDay2Part2(t *testing.T) {
	result := Part2("test_input.txt")
	if result != 900 {
		t.Log("Expected result to be 900")
		t.Fail()
	}
}

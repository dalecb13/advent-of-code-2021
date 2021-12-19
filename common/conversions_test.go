package common

import "testing"

func TestBinaryStringToDecimal(t *testing.T) {
	inputs := []string{
		"0",
		"1",
		"10",
		"11",
		"100",
		"101",
		"110",
		"111",
	}

	expectedOutputs := []int64{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
	}

	for inputIdx, input := range inputs {
		actualOutput := BinaryStringToDecimal(input)
		if actualOutput != expectedOutputs[inputIdx] {
			t.Logf("Expected actualResult %+v to be equal to expectedResult %+v", actualOutput, expectedOutputs[inputIdx])
			t.Fail()
		}
	}
}

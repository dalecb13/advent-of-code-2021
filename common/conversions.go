package common

import (
	"fmt"
	"strconv"
)

func StringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

func BinaryStringToDecimal(inputBinary string) int64 {
	i, _ := strconv.ParseInt(inputBinary, 2, 64)
	return i
}

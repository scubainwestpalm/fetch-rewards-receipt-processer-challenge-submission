package utils

import (
	"fmt"
	"math"
	"regexp"
)

func CountAlphanumericCharsInString(str string) int {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println(err)
	}
	strippedString := reg.ReplaceAllString(str, "")
	return len(strippedString)
}

func Float64IsWholeNumber(n float64) bool {
	return n == math.Trunc(n)
}

func Float64IsMultipleOf(n, divisor float64) bool {
	return n > divisor && math.Mod(n, divisor) == 0
}

func IntIsMultipleOf(n, divisor int) bool {
	return n > divisor && n%divisor == 0
}

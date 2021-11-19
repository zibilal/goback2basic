package presentation

import (
	"fmt"
	"testing"
)

func div2(numerator int, denominator int, result *float64)  {
	if denominator == 0 {
		*result = 0
	}
	*result = float64(numerator) / float64(denominator)
}

func TestUsingFunc(t *testing.T) {
	result := 0.0
	div2(5, 2 , &result)
	fmt.Println(result)
}

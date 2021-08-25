package presentation

import (
	"testing"
)

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func TestFunctionReturnsFunction(t *testing.T) {
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		t.Log(twoBase(i), threeBase(i))
	}
}

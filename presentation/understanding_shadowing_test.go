package presentation

import (
	"testing"
)

func TestShadowVariables(t *testing.T) {
	x := 10
	if x > 5 {
		t.Log(x)
		x := 5
		t.Log(x)
	}
	t.Log(x)
}

func TestShadowingWithMultipleAssignment(t *testing.T) {
	var (
		x, y int
	)
	x, y = 10, 11
	if x > 5 {
		x1, y1 := 5, 20
		t.Log(x1, y1)
	}
	t.Log(x, y)
}

func TestShadowingPackageName(t *testing.T) {
	x := 10
	t.Log(x)
	fmt := "oops"
	t.Log(fmt)
	// fmt.Println("Test")
}

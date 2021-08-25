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
	x := 10
	if x > 5 {
		x, y := 5, 20
		t.Log(x, y)
	}
	t.Log(x)
}

func TestShadowingPackageName(t *testing.T) {
	x := 10
	t.Log(x)
	fmt := "oops"
	t.Log(fmt)
	//fmt.Println("Test")
}

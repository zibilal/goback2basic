package presentation

import (
	"fmt"
	"testing"
)

type Adder struct {
	start int
}

func (a Adder)  AddTo(val int) int {
	return a.start + val
}

func TestAdder(t *testing.T) {
	myAdder := Adder {10}
	t.Log(myAdder.AddTo(5))

	// assign the method to func(int)int, called method value
	f1 := myAdder.AddTo
	t.Log(f1(10))

	// Create a function from the type itself. Called method expression
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 10))
}

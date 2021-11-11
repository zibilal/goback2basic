package presentation

import (
	"fmt"
	"testing"
)

func TestSliceFirst(t *testing.T) {
	var x = []int {10, 20, 30}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 40, 50, 60, 70)
	fmt.Println("x", len(x))
	fmt.Println("x", cap(x))

	x = append(x, 80)
	fmt.Println("2-x", len(x))
	fmt.Println("2-x", cap(x))


	fmt.Println("The list", x)
	fmt.Println("The list[2:5]", x[2:5])
	fmt.Println("The list[2:5]", x[2:5:6])

	xslice := x[2:5:5]

	fmt.Println(cap(x), len(x))
	fmt.Println(cap(xslice), len(xslice))
	fmt.Println(xslice)
}

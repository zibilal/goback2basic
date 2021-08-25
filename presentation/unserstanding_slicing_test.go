package presentation

import (
	"reflect"
	"testing"
)

func TestSlicingSlice(t *testing.T) {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	t.Log("x:", x)
	t.Log("y:", y)
	t.Log("z:", z)
	t.Log("d:", d)
	t.Log("e:", e)
}

func TestSliceShareStorage(t *testing.T) {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	t.Log("Before:")
	t.Log("x:", x)
	t.Log("y:", y)
	t.Log("z:", z)
	x[1] = 20
	y[0] = 10
	z[1] = 30
	t.Log("After")
	t.Log("x:", x)
	t.Log("y:", y)
	t.Log("z:", z)
}

func TestSliceShareWithAppend(t *testing.T) {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	t.Log("Before:", y)
	t.Log(cap(x), cap(y))
	y = append(y, 30)
	t.Log("x:", x)
	t.Log("y:", y)

	// y = append(y, 40)
	// t.Log("x:", x)
	// t.Log("y:", y)
}

func TestSliceWithMoreAppends(t *testing.T) {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	t.Log(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	t.Log("x:", x)
	t.Log("y:", y)
	t.Log("z:", z)
}

func TestSliceSlicingComplete(t *testing.T) {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	// y x[low:end:max]
	y := x[:2:2]
	z := x[2:4:4] // capacity = max - low
	t.Log(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	t.Log("x:", x)
	t.Log("y:", y)
	t.Log("z:", z)
}

func TestConvertingArraysToSlices(t *testing.T){
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	x[0] = 10
	t.Log("x:", x, reflect.ValueOf(x).Type())
	t.Log("y:", y, reflect.ValueOf(x).Type())
	t.Log("z:", z, reflect.ValueOf(x).Type())
}

func TestCopySlice(t *testing.T) {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	t.Log(y, num)
}

func TestCopyPartially(t *testing.T) {
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)
	num := copy(y, x)
	t.Log(y, num)
}

func TestCopyMiddle(t *testing.T) {
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)
	copy(y, x[2:])
	t.Log(y)
}
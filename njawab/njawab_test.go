package njawab

import (
	"fmt"
	"testing"
)

func TestUpdateIdxTerakhir(t *testing.T) {
	aslice := []int {1, 2, 3, 4}
	fmt.Println("Before", aslice)
	_ = UpdateIdxTerakhir(aslice, 15)

	fmt.Println("After", aslice)
}

func TestUpdateAppend(t *testing.T) {
	aslice := make([]int, 0, 5)
	fmt.Println("Before", aslice, cap(aslice))
	aslice = append(aslice, 1)
	aslice = append(aslice, 2)
	aslice = append(aslice, 3)
	aslice = append(aslice, 4)
	aslice = UpdateAppend(aslice, 15)

	fmt.Println("After", aslice)
}

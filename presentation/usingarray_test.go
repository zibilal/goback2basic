package presentation

import (
	"fmt"
	"testing"
)

func TestFirstArray(t *testing.T) {
	arr1 := [12]string {
		"surya", "sami", "bilal", "surya1", "sami1", "bilal1", "surya2", "sami2", "bilal2", "surya3", "sami3", "bilal3",
	}

	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))

	for i, v := range arr1 {
		fmt.Println(i, "[", v, "]")
	}
	fmt.Println("T1", arr1[0:5])
	fmt.Println("T2", arr1[1:5])
}

func TestMultipleDimArray(t *testing.T) {
	mat1 := [2][3]int {
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(mat1[0][2])
	fmt.Println(mat1[1])
	fmt.Println(len(mat1[0]))
}

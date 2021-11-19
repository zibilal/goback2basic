package presentation

import (
	"fmt"
	"testing"
)

func TestUsingForLoop(t *testing.T) {
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++{
		fmt.Println(i)
	}
}

func TestFibonacci(t *testing.T) {
	fibEx := []int{ 0, 1, 1, 2, 3, 5, 8, 13}

	for _, v := range fibEx {
		fmt.Println(v)
	}
	fmt.Println(fibEx)
	var fibEx2 = []int{0}
	for i, j, index := 0, 1, 0; index < 7; index++ {
		fibEx2 = append(fibEx2, j)
		i, j = j, i+j
	}
	fmt.Println(fibEx2)
}

func TestMax(t *testing.T) {
	pivot := 6

	left, right := make([]int, 0), make([]int, 0)
	v := 0
	for v <= pivot {
		left = append(left, v)
		v++
	}
	fmt.Println("Test", left, v)
	for ;; v++ {
		if v > pivot && v < 13 {
			right = append(right, v)
		} else {
			break
		}
	}
	fmt.Println("Test", right)
}

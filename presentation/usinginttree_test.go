package presentation

import (
	"fmt"
	"testing"
)

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func TestIntTree(t *testing.T) {
	aslice := []int {
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	pivot := aslice[len(aslice) / 2]
	fmt.Println(pivot)

	atree := &IntTree{}

	for _, v := range aslice {
		atree.Insert(v)
	}

	fmt.Println(atree.Contains(9))
}

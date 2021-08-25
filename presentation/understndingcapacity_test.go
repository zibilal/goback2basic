package presentation

import (
	"testing"
)

func TestCapacitySlice(t *testing.T) {

	var x []int
	t.Log(x, len(x), cap(x))
	x = append(x, 10)
	t.Log(x, len(x), cap(x))
	x = append(x, 20)
	t.Log(x, len(x), cap(x))
	x = append(x, 30)
	t.Log(x, len(x), cap(x))
	x = append(x, 40)
	t.Log(x, len(x), cap(x))
	x = append(x, 50)
	t.Log(x, len(x), cap(x))

}

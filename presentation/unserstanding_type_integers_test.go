package presentation

import (
	"fmt"
	"testing"
)

func TestAssigning(t *testing.T) {
	var aBool bool

	t.Log("Abool", aBool)

	aBool2 := true

	t.Log("Abool2", aBool2)

	aBool2 = true

	aBool4 := true

	t.Log(aBool4)

}

func TestInteger(t *testing.T) {
	var var1 int8 = 12
	var var2 int32 = 12

	var var3 int

	var3 = int(var1) + int(var2)

	fmt.Println("Var3", var3)
}

func TestInteger2(t *testing.T) {
	x := [...]int{1, 5: 4, 6, 10: 100, 15}
	t.Log(x)

	var y = [5]byte{
		1, 3, 3, 4, 5,
	}
	t.Log(y)

	var z [5]int
	t.Log(z)
	t.Log(z[3])
}

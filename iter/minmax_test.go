package iter

import (
	"goback2basic/assert"
	"testing"
)

func TestMax(t *testing.T) {
	t.Log("Find max on a given array")
	{
		arr := []int {45, 1, 3, -1, 18, 200, 5, 100}

		expected := 200
		actual := Max(arr)

		assert.EqualWithMessage(t, "Array maximum value is 200", expected, actual)
	}
}

func TestMin(t *testing.T) {
	t.Log("Find min on a given array")
	{
		arr := []int {45, 1, 3, -1, 18, 200, 5, 100}

		expected := -1
		actual := Min(arr)

		assert.EqualWithMessage(t, "Array minimum value is -1", expected, actual)
	}
}
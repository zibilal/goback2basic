package quicksort

import (
	"goback2basic/assert"
	"testing"
)

func TestSortOne(t *testing.T) {
	t.Log("Test Sorting of [57,16,207,94,17,2,138,12,73,103,77] array")
	t.Log("----------------------------------------------------------")
	{
		arr := []int {57,16,207,94,17,2,138,12,73,103,77}
		sorted := SortOne(arr)

		assert.EqualWithMessage(t, "Array should be sorted", []int{2, 12, 16, 17, 57, 73, 77, 94, 103, 138, 207}, sorted)
	}
}

func TestSortTwo(t *testing.T) {
	t.Log("Test Sorting of [57,16,207,94,17,2,138,12,73,103,77] array")
	t.Log("----------------------------------------------------------")
	{
		arr := []int {57,16,207,94,17,2,138,12,73,103,77}
		SortTwo(arr, 0, len(arr)-1)

		assert.EqualWithMessage(t, "Array should be sorted", []int{2, 12, 16, 17, 57, 73, 77, 94, 103, 138, 207}, arr)
	}
}
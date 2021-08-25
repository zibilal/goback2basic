package binary_search

import (
	"goback2basic/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Log("Testing Binary Search find 15 in a given array")
	t.Log("----------------------------------------------")
	{
		arr := []int {-5, 11, 15, 21, 151, 201, 2000}
		idx := Search(15, arr, 0, len(arr))
		assert.EqualWithMessage(t, "15 is in second index", 2, idx)
	}
}

package growth2dim

import "testing"

func Test2Dim(t *testing.T) {
	t.Log("Starting...")
	row, col := 4, 4
	arr := make([][]int, row)

	for i:=0; i < row; i++ {
		arr[i] = make([]int, col)
	}
	t.Log("Finished", arr)
}

func TestAgain(t *testing.T) {
	var result []string
	t.Log("Again:", len(result))
}
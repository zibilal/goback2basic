package growth2dim

import (
	"strconv"
	"strings"
)

func countMax(upRight []string) int64 {

	row, col := 0, 0
	for _, v := range upRight {
		splits := strings.Split(v, " ")
		tmpRow, _ := strconv.Atoi(splits[0])
		tmpCol, _ := strconv.Atoi(splits[1])

		if tmpRow > row {
			row = tmpRow
		}
		if tmpCol > col {
			col = tmpCol
		}
	}

	arr := make([][]int8, row)
	for i:=0; i < row; i++ {
		arr[i] = make([]int8, col)
	}

	var max int8
	for _, v := range upRight {
		splits := strings.Split(v, " ")
		tmpRow, _ := strconv.Atoi(splits[0])
		tmpCol, _ := strconv.Atoi(splits[1])

		for i:=0; i < tmpRow; i++ {
			for j:=0; j < tmpCol; i++ {
				arr[i][j] += 1

				if arr[i][j] > max {
					max = arr[i][j]
				}
			}
		}
	}

	var count int64
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if arr[i][j] == max {
				count++
			}
		}
	}

	return count
}

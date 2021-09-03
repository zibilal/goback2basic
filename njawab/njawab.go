package njawab

import "fmt"

func UpdateIdxTerakhir(data []int, v int) bool {
	data[len(data)-1] = v
	return v > 5
}

func UpdateAppend(data []int, v int) []int {
	fmt.Println("Cap1", cap(data))
	data[len(data)-1] = v
	fmt.Println("Cap2", cap(data))
	fmt.Println("Inside", data)

	return data
}

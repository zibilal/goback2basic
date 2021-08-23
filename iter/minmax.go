package iter

func Max(arr []int) int{
	max := -9999
	for _, v :=range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(arr []int) int {
	min := 9999
	for _, v :=range arr {
		if v < min{
			min = v
		}
	}
	return min
}
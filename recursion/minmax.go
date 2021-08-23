package recursion

func Max(arr []int, index, l int)int {
	if  index >= l-2 {
		if arr[index] > arr[index+1] {
			return arr[index]
		} else {
			return arr[index+1]
		}
	}
	max := Max(arr, index + 1, l)
	if arr[index] > max {
		return arr[index]
	} else {
		return max
	}
}

func Min(arr []int, index,l int) int {
	if index >= l-2 {
		if arr[index] < arr[index + 1] {
			return arr[index]
		} else {
			return arr[index+1]
		}
	}
	min := Min(arr, index + 1, l)
	if arr[index] < min {
		return arr[index]
	} else {
		return min
	}
}

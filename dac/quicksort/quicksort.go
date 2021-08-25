package quicksort
var a = 10
func SortOne(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	var less []int
	var more []int

	// partition process
	for _, v := range arr {
		if v < pivot {
			less = append(less, v)
		} else if v > pivot {
			more = append(more, v)
		}
	}

	less = SortOne(less)
	more = SortOne(more)

	// combining
	less = append(less, pivot)
	more = append(less, more...)

	return more
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i, j := low-1, low

	for ;j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i+1
}

func SortTwo(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)

		SortTwo(arr, low, pi-1)
		SortTwo(arr, pi+1, high)
	}
}
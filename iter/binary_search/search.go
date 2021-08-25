package binary_search

func Search(n int, arr []int) int {
	low := 0
	high := len(arr) - 1

	for high >= low  {
		mid := (low + high) / 2

		if arr[mid] == n {
			return mid
		}

		if arr[mid] < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

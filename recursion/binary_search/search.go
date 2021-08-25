package binary_search

func Search(n int, arr [] int, low, high int) int {
	if high < low {
		return -1
	}
	mid := (low + high) / 2
	if arr[mid] == n {
		return mid
	}

	if arr[mid] > n {
		return Search(n, arr, low, mid-1)
	} else {
		return Search(n, arr, mid+1, high)
	}
}
package iter

func Fib(n int) int {
	if n < 2 {
		return n
	}
	first, second  := 0, 1
	for i:=2; i <= n; i++ {
		first, second = second, first+second
	}

	return second
}

func Fib2(n int) int {
	var result int
	for i, first, second := 0, 0, 1; i <= n; i, first, second = i+1, second, first+second {
		if i == n {
			result = first
		}
	}

	return result
}

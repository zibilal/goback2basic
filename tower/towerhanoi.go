package tower

import "fmt"

func Tohmove(n int , a, b, c string) {

	if n == 1 {
		fmt.Printf("Move disk %d from source %s to destination %s\n", n, a, b)
		return
	}

	Tohmove(n-1, a, c, b)
	fmt.Printf("Move disk %d from source %s to destination %s\n", n, a, b)
	Tohmove(n-1, c, b, a)
}

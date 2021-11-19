package presentation

import (
	"fmt"
	"testing"
)

func TestUsingSwitch(t *testing.T) {
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}

func TestAnotherUsingSwitch(t *testing.T) {

	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 5:
			fmt.Println(word, "is a long word!")
		}

	}

	// worldlen is garbace collected

	a := 5

	switch {
	case a == 2:
		fmt.Println("a is 2")
	case a == 3:
		fmt.Println("a is 3")
	case a == 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("a is ", a)
	}

	// a still alive
}

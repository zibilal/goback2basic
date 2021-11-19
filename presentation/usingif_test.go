package presentation

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestUsingIf(t *testing.T) {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}

func TestUsingIf_Two(t *testing.T) {

	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
	//
	//
	//
	/// lama
	// n is out of scope
	// fmt.Println(n)
}

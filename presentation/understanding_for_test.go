package presentation

import (
	"testing"
)

func TestConfusingCode(t *testing.T) {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				t.Log("FizzBuzz")
			} else {
				t.Log("Fizz")
			}
		} else if i%5 == 0 {
			t.Log("Buzz")
		} else {
			t.Log(i)
		}
	}
}

func TestForWithContinue(t *testing.T) {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			t.Log("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			t.Log("Fizz")
			continue
		}
		if i%5 == 0 {
			t.Log("Buzz")
			continue
		}
		t.Log(i)
	}
}

func TestForModifyingForRange(t *testing.T) {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	t.Log(evenVals)
}

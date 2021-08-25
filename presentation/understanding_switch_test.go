package presentation

import (
	"testing"
)

func TestForSwitchBlank(t *testing.T) {
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			t.Log(word, "is a short word!")
		case wordLen > 10:
			t.Log(word, "is a long word!")
		default:
			t.Log(word, "is exactly the right length.")
		}
	}
}

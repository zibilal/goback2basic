package hashment

import (
	"fmt"
	"testing"
)

func TestBitwise(t *testing.T) {
	var x byte = 1<<1 | 1<<5
	var y byte = 1<<1 | 1<<2

	t.Logf("%08b - %d\n", x, x)
	t.Logf("%08b - %d\n", y, y)

	t.Log("-----------------")

	t.Logf("%08b - %d\n", x&y, x&y)
	t.Logf("%08b - %d\n", x|y, x|y)
	t.Logf("%08b - %d\n", x^y, x^y)	// The symmetric difference {2, 5}
	t.Logf("%08b - %d\n", x&^y, x&^y)	// The difference {5}

	for i := uint(0); i < 8; i++ {
		if x & (1<<i) != 0 {
			fmt.Println(i)
		}
	}

	t.Logf("%08b\n", x<<1)	// 01000100, the set {2, 6}
	t.Logf("%08b\n", x>>1)	// 00010001, the set {0, 4}

	t.Log("-----------------")

	t.Logf("1: %08b - 1<<1: %08b - 1>>1: %08b\n", 1, 1<<1, 2>>1)
	t.Logf("1: %d - 1<<5: %d - \n", 1, 1<<5)
	t.Logf("1: %08b - 1<<5: %08b\n", 1, 1 <<5)
}

func TestHashCodeSimple(t *testing.T) {
	t.Log(HashCodeSimple(12))
	t.Log(HashCodeSimple(13))
}

func TestHashCodeOther(t *testing.T) {
	t.Log(HashCodeOther(12))
	t.Log(HashCodeOther(13))
}

func TestHashCodePolynomial(t *testing.T) {
	l := 13
	h1 := HashCodePolynomial("HELLO")
	h2 := HashCodePolynomial("A Test")
	h3 := HashCodePolynomial("On Data")
	h4 := HashCodePolynomial("On Data")
	t.Log(h1, CompressHashCodeDivision(h1, l))
	t.Log(h2, CompressHashCodeDivision(h2, l))
	t.Log(h3, CompressHashCodeDivision(h3, l))
	t.Log(h4, CompressHashCodeDivision(h4, l))
}

func TestCharOfString(t *testing.T) {
	s := "HELLO"
	for i:=0; i < len(s); i++ {
		t.Log(s[i], int(s[i]))
	}
}

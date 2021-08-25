package presentation

import (
	"testing"
)

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func TestConst(t *testing.T) {
	const y = "hello"

	t.Log(x)
	t.Log(y)

	// cannot assign to x
	// x = x + 1

	// cannot assign to y
	// y = "bye"

	t.Log(x)
	t.Log(y)
}

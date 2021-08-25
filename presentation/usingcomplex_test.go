package presentation

import (
	"math/cmplx"
	"reflect"
	"testing"
)

func TestComplex(t *testing.T) {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	t.Log(x + y)
	t.Log(x - y)
	t.Log(x * y)
	t.Log(x / y)
	t.Log(real(x))
	t.Log(imag(x))
	t.Log(cmplx.Abs(x))

	t.Log(reflect.ValueOf(x).Type())
}

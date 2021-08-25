package recursion
import (
	"goback2basic/assert"
	"testing"
)

func TestFib(t *testing.T) {
	t.Log("Test Fib function, 8-th element equal 21")
	t.Log("----------------------------------------")
	{
		actual := Fib(8)
		assert.EqualWithMessage(t, "8-th element should equal 21", 21, actual)
	}
}

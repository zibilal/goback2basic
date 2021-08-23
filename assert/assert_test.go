package assert

import (
	"errors"
	"path"
	"runtime"
	"testing"
)

func TestCaller(t *testing.T) {
	t.Log("Testing runtime.Caller")
	t.Log("----------------------")
	{
		afunc, file, line, ok := runtime.Caller(1)

		t.Log("Afunc:", afunc, ", File:", file, ", Line:", line, ", Ok:", ok)
		t.Log("Afunc:", afunc, ", File:", path.Base(file), ", Line:", line, ", Ok:", ok)
	}
}

func TestEqualPatternWithMessage(t *testing.T) {
	t.Log("Testing EqualPatternWithMessage")
	t.Log("Test starts with \"failed\"")
	t.Log("-------------------------------")
	{
		pattern := "^failed"
		sampleError := errors.New("failed: Something not cool is happened")

		EqualPatternWithMessage(t, "Error starts with failed", pattern, sampleError.Error())
	}
}

func TestEqualWithMessage(t *testing.T) {
	t.Log("TestingEqualWithMessage")
	t.Log("input is equal to output")
	t.Log("------------------------")
	{
		input := struct {
			Email string
		}{
			"test@example.com",
		}

		output := struct {
			Email string
		}{
			"test@example.com",
		}

		EqualWithMessage(t, "Input is equl Output", input, output)
	}
}
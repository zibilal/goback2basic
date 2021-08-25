package assert

import (
	"fmt"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
	skip = 1
)

func EqualWithMessage(t *testing.T, msg string, expected, actual interface{}) {
	_, filePath, line, _ := runtime.Caller(skip)

	if reflect.DeepEqual(expected, actual) {
		fmt.Printf("%s:%d - %s - %s:[Actual %v ] equal [Expected %v]\n", path.Base(filePath), line, success, msg, actual, expected)
	} else {
		fmt.Printf("%s:%d - %s - %s:[Actual %v ] NOT equal [Expected %v]\n", path.Base(filePath), line, failed, msg, actual, expected)
		t.FailNow()
	}
}

func EqualPatternWithMessage(t *testing.T, msg string, pattern, value string) {
	matched, _ := regexp.MatchString(pattern, value)
	_, filePath, line, _ := runtime.Caller(skip)
	if matched {
		fmt.Printf("%s:%d - %s - %s:[Value %s ] matched [Pattern %s]\n", path.Base(filePath), line, success, msg, value, pattern)
	} else {
		fmt.Printf("%s:%d - %s - %s:[Value %s ] NOT matched [Pattern %s]\n", path.Base(filePath), line, failed, msg, value, pattern)
		t.FailNow()
	}
}

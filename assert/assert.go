package assert

import (
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
		t.Logf("%s:%d - %s - %s:[Actual %v ] equal [Expected %v]", path.Base(filePath), line, success, msg, actual, expected)
	} else {
		t.Fatalf("%s:%d - %s - %s:[Actual %v ] NOT equal [Expected %v]", path.Base(filePath), line, failed, msg, actual, expected)
	}
}

func EqualPatternWithMessage(t *testing.T, msg string, pattern, value string) {
	matched, _ := regexp.MatchString(pattern, value)
	_, filePath, line, _ := runtime.Caller(skip)
	if matched {
		t.Logf("%s:%d - %s - %s:[Value %s ] matched [Pattern %s]", path.Base(filePath), line, success, msg, value, pattern)
	} else {
		t.Fatalf("%s:%d - %s - %s:[Value %s ] NOT matched [Pattern %s]", path.Base(filePath), line, failed, msg, value, pattern)
	}
}

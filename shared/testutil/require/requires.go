package require

import (
	"github.com/prysmaticlabs/prysm/shared/testutil/assertions"
)

// Equal compares values using comparison operator.
func Equal(tb assertions.AssertionTestingTB, expected, actual interface{}, msg ...interface{}) {
	assertions.Equal(tb.Fatalf, expected, actual, msg...)
}

// NotEqual compares values using comparison operator.
func NotEqual(tb assertions.AssertionTestingTB, expected, actual interface{}, msg ...interface{}) {
	assertions.NotEqual(tb.Fatalf, expected, actual, msg...)
}

// DeepEqual compares values using DeepEqual.
func DeepEqual(tb assertions.AssertionTestingTB, expected, actual interface{}, msg ...interface{}) {
	assertions.DeepEqual(tb.Fatalf, expected, actual, msg...)
}

// NoError asserts that error is nil.
func NoError(tb assertions.AssertionTestingTB, err error, msg ...interface{}) {
	assertions.NoError(tb.Fatalf, err, msg...)
}

// ErrorContains asserts that actual error contains wanted message.
func ErrorContains(tb assertions.AssertionTestingTB, want string, err error, msg ...interface{}) {
	assertions.ErrorContains(tb.Fatalf, want, err, msg...)
}

// NotNil asserts that passed value is not nil.
func NotNil(tb assertions.AssertionTestingTB, obj interface{}, msg ...interface{}) {
	assertions.NotNil(tb.Fatalf, obj, msg...)
}

package assert

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func ExpectEq(actual interface{}, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		debug.PrintStack()
		t.Errorf("expected %T: %#v, %T: got %#v", expected, expected, actual, actual)
		t.FailNow()
	}
}

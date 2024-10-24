package assertions

import (
	"testing"
)

func True(t *testing.T, args ...bool) {
	for i := range args {
		if !args[i] {
			t.Logf("expected %v (%T) to be true", args[i], args[i])
			t.Fail()
		}
	}
}

func False(t *testing.T, args ...bool) {
	for i := range args {
		if args[i] {
			t.Logf("expected %v (%T) to be false", args[i], args[i])
			t.Fail()
		}
	}
}

func Equal[T comparable](t *testing.T, args ...T) {
	if len(args) <= 1 {
		return
	}
	for i := range args[:len(args)-1] {
		if args[i] != args[i+1] {
			t.Logf("expected %v (%T) to equal %v (%T)", args[i], args[i], args[i+1], args[i+1])
			t.Fail()
		}
	}
}

func Unequal[T comparable](t *testing.T, args ...T) {
	if len(args) <= 1 {
		return
	}
	for i := range args[:len(args)-1] {
		if args[i] == args[i+1] {
			t.Log("expected", args[i], "to unequal", args[i+1])
			t.Fail()
		}
	}
}

func NilPtr[T any](t *testing.T, args ...*T) {
	for _, arg := range args {
		if arg != nil {
			t.Log("expected", arg, "to equal nil")
			t.Fail()
		}
	}
}

func NotNilPtr[T any](t *testing.T, args ...*T) {
	for _, arg := range args {
		if arg == nil {
			t.Log("expected", arg, "to not equal nil")
			t.Fail()
		}
	}
}

func PanicError(t *testing.T, f func(), e string) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				if err.Error() != e {
					t.Log("function panicked with different error message")
					t.Fail()
				}
			} else {
				t.Log("function panicked without error")
				t.Fail()
			}
		} else {
			t.Log("function did not panic or no error was passed")
			t.Fail()
		}
	}()
	f()
}

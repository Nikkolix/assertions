package assertions

import (
	"runtime"
	"testing"
)

var instance *testing.T

func InitInstance(t *testing.T) {
	instance = t
}

func True(args ...bool) {
	for i := range args {
		if !args[i] {
			pc, file, line, ok := runtime.Caller(1)
			if ok {
				f := runtime.FuncForPC(pc)
				instance.Logf("called by %s at %s: %d", f.Name(),file, line)
			}
			instance.Logf("expected %v (%T) to be true", args[i], args[i])
			instance.Fail()
		}
	}
}

func False(args ...bool) {
	for i := range args {
		if args[i] {
			pc, file, line, ok := runtime.Caller(1)
			if ok {
				f := runtime.FuncForPC(pc)
				instance.Logf("called by %s at %s: %d", f.Name(), file, line)
			}
			instance.Logf("expected %v (%T) to be false", args[i], args[i])
			instance.Fail()
		}
	}
}

func Equal[T comparable](args ...T) {
	if len(args) <= 1 {
		return
	}
	for i := range args[:len(args)-1] {
		if args[i] != args[i+1] {
			instance.Logf("expected %v (%T) to equal %v (%T)", args[i], args[i], args[i+1], args[i+1])
			instance.Fail()
		}
	}
}

func Unequal[T comparable](args ...T) {
	if len(args) <= 1 {
		return
	}
	for i := range args[:len(args)-1] {
		if args[i] == args[i+1] {
			instance.Log("expected", args[i], "to unequal", args[i+1])
			instance.Fail()
		}
	}
}

func NilPtr[T any](args ...*T) {
	for _, arg := range args {
		if arg != nil {
			instance.Log("expected", arg, "to equal nil")
			instance.Fail()
		}
	}
}

func NotNilPtr[T any](args ...*T) {
	for _, arg := range args {
		if arg == nil {
			instance.Log("expected", arg, "to not equal nil")
			instance.Fail()
		}
	}
}

func PanicError(f func(), e string) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				if err.Error() != e {
					instance.Log("function panicked with different error message")
					instance.Fail()
				}
			} else {
				instance.Log("function panicked without error")
				instance.Fail()
			}
		} else {
			instance.Log("function did not panic or no error was passed")
			instance.Fail()
		}
	}()
	f()
}

package assertions

import (
	"errors"
	"math/rand"
	"testing"
)

const NFuzzInput = 1000

//Equal

func TestEqual(t *testing.T) {
	Equal(t, 1, 1, 1)
	Equal(t, 1, 1)
	Equal(t, 1)
	Equal[int](t)
}

func TestEqual2(t *testing.T) {
	t2 := new(testing.T)
	Equal(t2, 1, 2)
	if !t2.Failed() {
		t.Fatal("test should fail")
	}
}

func FuzzEqual(f *testing.F) {
	for x := 0; x < NFuzzInput; x++ {
		f.Add(rand.Int(), rand.Int(), rand.Int())
	}
	f.Fuzz(func(t *testing.T, x, y, z int) {
		t2 := new(testing.T)
		Equal(t2, x, y, z)
		if t2.Failed() && x == y && y == z && x == z {
			t.Fatal("test should fail", x, "==", y, "==", z)
		}
	})
}

//Unequal

func TestUnequal(t *testing.T) {
	Unequal(t, 0, 1)
	Unequal(t, 1)
	Unequal[int](t)
}

func TestUnequal2(t *testing.T) {
	t2 := new(testing.T)
	Unequal(t2, 1, 1)
	if !t2.Failed() {
		t.Fatal()
	}
}

func FuzzUnequal(f *testing.F) {
	for x := 0; x < NFuzzInput; x++ {
		f.Add(rand.Int(), rand.Int(), rand.Int())
	}
	f.Fuzz(func(t *testing.T, x, y, z int) {
		t2 := new(testing.T)
		Unequal(t2, x, y)
		if t2.Failed() && x != y && y != z && x != z {
			t.Fatal("test should fail", x, "!=", y, "!=", z)
		}
	})
}

//NilPtr

func TestNilPtr(t *testing.T) {
	var p *int = nil
	var p2 *int = nil
	NilPtr(t, p, p2)
	NilPtr[*any](t, nil, nil)
}

func TestNilPtr2(t *testing.T) {
	t2 := new(testing.T)
	NilPtr(t2, new(int))
	if !t2.Failed() {
		t.Fatal()
	}
}

//NotNilPtr

func TestNotNilPtr(t *testing.T) {
	NotNilPtr(t, new(int), new(int))
}

func TestNotNilPtr2(t *testing.T) {
	t2 := new(testing.T)
	NotNilPtr[*any](t2, nil)
	if !t2.Failed() {
		t.Fatal()
	}
}

//PanicError

func TestPanicError(t *testing.T) {
	PanicError(t, func() { panic(errors.New("test")) }, "test")
	t2 := new(testing.T)
	PanicError(t2, func() { panic(errors.New("true")) }, "false")
	if !t2.Failed() {
		t.Fatal()
	}
	t3 := new(testing.T)
	PanicError(t3, func() { panic("no error") }, "wrong")
	if !t3.Failed() {
		t.Fatal()
	}
	t4 := new(testing.T)
	PanicError(t4, func() {}, "nothing")
	if !t4.Failed() {
		t.Fatal()
	}
}

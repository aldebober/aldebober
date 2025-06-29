package main

import "testing"

func TestFoo(t *testing.T) {
	total := foo(23, 56, 78)
	if total != 160 {
		t.Errorf("Func Foo: %d mismatch result %d", total, 160)
	}
}

func TestSubtract(t *testing.T) {
	r := subtract(10, 3)
	if r != 7 {
		t.Errorf("subtract func result %d mismatch test %d", r, 7)
	}
}

func TestAdd(t *testing.T) {
	r := add(13, 23)
	if r != 36 {
		t.Errorf("Add func rusalt %d mismatch %d", r, 36)
	}
}

func TestDoMath(t *testing.T) {
	r := doMath(12, 54, add)
	if r != 66 {
		t.Errorf("doMath result %d mismatch %d", r, 66)
	}

}

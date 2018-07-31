package main_test

import "testing"

func TestEqual(t *testing.T) {
	if 1 != 1 {
		t.Errorf("1 != 1")
	}
}

func TestBoolean(t *testing.T) {
	if true == false {
		t.Errorf("true != false")
	}
}

func TestSomething(t *testing.T) {
	if 2 == 3 {
		t.Errorf("2 == 3")
	}
}

func TestSilly(t *testing.T) {
	if 4 == 2 {
		t.Errorf("4 == 2")
	}
}

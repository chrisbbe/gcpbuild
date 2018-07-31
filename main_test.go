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

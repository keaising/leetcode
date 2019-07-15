package main

import "testing"

func TestCal(t *testing.T) {
	ans := calculate("3+2*2")
	if ans != 7 {
		t.Error("!!!")
	}
}

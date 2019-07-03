package main

import "testing"

func TestDuplicate(t *testing.T) {
	ret := containsNearbyDuplicate([]int{1,2,3,1}, 3)
	if ret {
		t.Error("!!!!")
	}
}

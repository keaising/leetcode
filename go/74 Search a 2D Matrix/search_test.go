package main

import "testing"

func TestSearch(t *testing.T) {
	arr := [][]int{
		{1},
	}
	ret := searchMatrix(arr, 1)
	if ret {
		t.Error("!!!")
	}
}

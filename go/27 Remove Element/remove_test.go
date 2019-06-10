package remove

import "testing"

func TestRemove(t *testing.T) {
	arr := []int{0,1,2,2,3,0,4,2}
	var ret = removeElement(arr, 2)
	t.Log(ret)
	t.Log(arr)
	t.Log(arr[0:ret])
}

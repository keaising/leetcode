package add

import (
	"fmt"
	"testing"
)

var arr = make([]string, 0)

func TestAdd(t *testing.T) {
	l1 := links([]int{2, 4, 3})
	l2 := links([]int{5, 6, 4})
	ret := addTwoNumbers(l1, l2)
	//printf(l1)
	//printf(l2)
	printf(ret)
	t.Fatal(arr)
	for _, s := range arr {
		t.Logf(s)
	}
}

func links(arr []int) *ListNode {
	head := new(ListNode)
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head.Val = arr[0]
	tail := arr[1:]
	head.Next = links(tail)
	return head
}

func printf(l *ListNode) {
	if l == nil {
		return
	}
	arr = append(arr, fmt.Sprintf("value: %d", l.Val))
	printf(l.Next)
}

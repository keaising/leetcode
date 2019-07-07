package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	n := 0
	for cur := root; cur != nil; {
		cur = cur.Next
		n++
	}
	r := n / k
	/*
	   根据数据关系可以得到下列公式
	   x+y=k
	   n=x*r+y*(r+1)
	   x ：其中个数少1的元素数量，x中元素个数：r
	   y ：其中个数多1的元素数量，y中元素个数：r+1
	*/
	y := n - k*r
	x := k - y
	cur := root
	result := []*ListNode{}
	for i := 0; i < y; i++ {
		ansHead := cur
		ans := ansHead
		cur = cur.Next
		for j := 1; j < r+1; j++ {
			ans.Next = cur
			ans = ans.Next
			cur = cur.Next
		}
		if ans != nil {
			ans.Next = nil
		}
		result = append(result, ansHead)
	}
	for i := 0; i < x; i++ {
		ansHead := cur
		ans := ansHead
		if cur != nil {
			cur = cur.Next
			for j := 1; j < r; j++ {
				ans.Next = cur
				ans = ans.Next
				cur = cur.Next
			}
		}
		if ans != nil {
			ans.Next = nil
		}
		result = append(result, ansHead)
	}
	return result
}

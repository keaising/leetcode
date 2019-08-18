package structure

func buildListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	node := head
	for i := 1; i < len(nums); i++ {
		node.Next = &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		node = node.Next
	}
	return head
}

func buildNums(head *ListNode) []int {
	if head == nil {
		return nil
	}
	ans := []int{}
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

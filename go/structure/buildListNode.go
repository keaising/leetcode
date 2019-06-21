package structure

func buildListNode(nums []int) *ListNode {
	nums = []int{1, 2, 3, 3, 4, 4, 5}
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

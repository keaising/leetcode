# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def detectCycle(self, head: ListNode) -> ListNode:
        l1 = l2 = head
        while l2 and l2.next:
            l1 = l1.next
            l2 = l2.next.next
            if l1 == l2:
                l1 = head
                while l1 != l2:
                    l1 = l1.next
                    l2 = l2.next
                return l2
        return None
        
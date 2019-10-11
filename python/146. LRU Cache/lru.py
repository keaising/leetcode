class LRUCache:
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.dic = {}
        self.dummy_head = Node(0, 0)
        self.dummy_tail = Node(0, 0)
        self.dummy_head.next = self.dummy_tail
        self.dummy_tail.prev = self.dummy_head

    def get(self, key: int) -> int:
        if key not in self.dic:
            return -1
        node = self.dic[key]
        self.remove(node)
        self.append(node)
        return node.val

    def put(self, key: int, value: int) -> None:
        if key in self.dic:
            self.remove(self.dic[key])
        node = Node(key, value)
        self.append(node)
        self.dic[key] = node

        if len(self.dic) > self.capacity:
            head = self.dummy_head.next
            self.remove(head)
            del self.dic[head.key]

    def remove(self, node):
        p = node.prev
        n = node.next
        p.next = n
        n.prev = p

    def append(self, node):
        tail = self.dummy_tail.prev
        tail.next = node
        self.dummy_tail.prev = node
        node.prev = tail
        node.next = self.dummy_tail


class Node:
    def __init__(self, k, v):
        self.key = k
        self.val = v
        self.prev = None
        self.next = None


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        level = [root]  # 遍历整棵树
        dic = {root:None}  # 构建了一个每个节点的父节点的 map，空间换时间
        while p not in dic or q not in dic:
            node = level.pop()
            if node.left:
                level.append(node.left)
                dic[node.left] = node
            if node.right:
                level.append(node.right)
                dic[node.right] = node
        
        ans = set()  # 将p的所有父节点都放到 ans
        while p:
            ans.add(p)
            p = dic[p]
        while q not in ans: # 从q开始，找到的第一个跟p相同的父节点就是LCA
            q = dic[q]
        return q
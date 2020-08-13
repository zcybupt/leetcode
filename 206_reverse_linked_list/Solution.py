class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        return self.reverse(None, head)

    def reverse(self, pre: ListNode, cur: ListNode):
        if not cur: return pre

        next = cur.next
        cur.next = pre

        return self.reverse(cur, next)

    def reverseList2(self, head: ListNode) -> ListNode:
        cur, pre = head, None
        while cur:
            cur.next, pre, cur = pre, cur, cur.next

        return pre


if __name__ == '__main__':
    head = ListNode(0)
    cur = head
    for i in range(1, 10):
        cur.next = ListNode(i)
        cur = cur.next

    solution = Solution()
    cur = solution.reverseList2(head)

    while cur:
        print(cur.val, end = " ")
        cur = cur.next

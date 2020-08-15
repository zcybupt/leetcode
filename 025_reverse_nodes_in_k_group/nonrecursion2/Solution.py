class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        dummy_head = ListNode(-1)
        dummy_head.next = head
        pre, end = dummy_head, dummy_head

        while end.next:
            for i in range(k):
                end = end.next
                if not end: return dummy_head.next
            start = pre.next
            next_node = end.next
            pre.next = self.reverse(start, next_node)
            start.next = next_node
            pre = start
            end = pre

        return dummy_head.next

    def reverse(self, head: ListNode, tail: ListNode) -> ListNode:
        pre = None
        while head != tail:
            next_node = head.next
            head.next, pre, head = pre, head, next_node

        return pre


if __name__ == '__main__':
    head = ListNode(1)
    cur = head
    for i in range(2, 10):
        cur.next = ListNode(i)
        cur = cur.next

    solution = Solution()
    cur = solution.reverseKGroup(head, 2)
    while cur:
        print(cur.val, end = ' ')
        cur = cur.next

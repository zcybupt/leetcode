class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        if not head or not head.next: return head

        tail = head
        for i in range(k):
            if not tail: return head
            tail = tail.next

        new_head = self.reverse(head, tail)
        head.next = self.reverseKGroup(tail, k)

        return new_head

    def reverse(self, head: ListNode, tail: ListNode) -> ListNode:
        pre = None
        while head != tail:
            next_node = head.next
            head.next, pre, head = pre, head, next_node

        return pre


if __name__ == '__main__':
    head = ListNode(0)
    cur = head
    for i in range(1, 10):
        cur.next = ListNode(i)
        cur = cur.next

    solution = Solution()
    cur = solution.reverseKGroup(head, 2)
    while cur:
        print(cur.val, end = ' ')
        cur = cur.next

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        dummy_head = ListNode(0)
        p, q = l1, l2
        carry = 0
        curr = dummy_head

        while p or q:
            x = p.val if p else 0
            y = q.val if q else 0
            tmp_sum = carry + x + y
            carry = tmp_sum // 10
            curr.next = ListNode(tmp_sum % 10)
            curr = curr.next

            p = p.next if p else None
            q = q.next if q else None

        if carry > 0:
            curr.next = ListNode(carry)

        return dummy_head.next


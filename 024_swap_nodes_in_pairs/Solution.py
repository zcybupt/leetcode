class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        if not head or not head.next: return head

        cur = head
        head = head.next
        latter = None

        while cur and cur.next:
            tmp = cur.next
            cur.next = tmp.next
            tmp.next = cur
            if latter: latter.next = tmp
            latter = cur
            cur = cur.next

        return head

    def swapPairs2(self, head: ListNode) -> ListNode:
        if not head or not head.next: return head

        first, second = head, head.next

        first.next = self.swapPairs2(second.next)
        second.next = first

        return second


if __name__ == '__main__':
    head = ListNode(1)
    cur = head
    for i in range(2, 10):
        cur.next = ListNode(i)
        cur = cur.next

    solution = Solution()
    cur = solution.swapPairs2(head)
    while cur:
        print(cur.val, end=" ")
        cur = cur.next

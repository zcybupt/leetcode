class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if not head or not head.next: return head

        dummy = ListNode(-1)
        tail = dummy
        l, r = head, head
        while l:
            while r and l.val == r.val: r = r.next
            if l.next == r:
                tail.next = l
                tail = l
                tail.next = None
            l = r

        return dummy.next


if __name__ == '__main__':
    head = ListNode(1)
    head.next = ListNode(1)
    head.next.next = ListNode(2)
    head.next.next.next = ListNode(3)
    head.next.next.next.next = ListNode(3)
    result = Solution().deleteDuplicates(head)

    while result:
        print(result.val)
        result = result.next

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if not head: return head

        tmp = head
        while tmp and tmp.next:
            if tmp.next.val == tmp.val: tmp.next = tmp.next.next
            else: tmp = tmp.next

        return head


if __name__ == '__main__':
    head = ListNode(1)
    head.next = ListNode(1)
    head.next.next = ListNode(2)
    head.next.next.next = ListNode(3)
    head.next.next.next.next = ListNode(3)
    Solution().deleteDuplicates(head)

    while head:
        print(head.val)
        head = head.next


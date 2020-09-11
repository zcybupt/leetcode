class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        if not head: return None
        tail = head
        length = 1
        while tail.next:
            length += 1
            tail = tail.next
        tail.next = head
        k = length - k % length
        for _ in range(k, 0, -1): tail = tail.next
        head = tail.next
        tail.next = None

        return head


def print_list(head: ListNode):
    while head:
        print(head.val, end = ' ')
        head = head.next


if __name__ == '__main__':
    head = ListNode(1)
    tmp = head
    for i in range(2, 6):
        tmp.next = ListNode(i)
        tmp = tmp.next
    head = Solution().rotateRight(head, 12)
    print_list(head)

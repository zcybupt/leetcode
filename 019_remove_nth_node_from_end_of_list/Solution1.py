class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        return head.next if self.remove_node(head, n) == n else head

    def remove_node(self, head: ListNode, n: int) -> int:
        if not head.next: return 1
        m = self.remove_node(head.next, n)
        if m == n:
            head.next = head.next.next

        return m + 1


if __name__ == '__main__':
    head = ListNode(1)
    head.next = ListNode(2)
    head.next.next = ListNode(3)
    head.next.next.next = ListNode(4)
    head.next.next.next.next = ListNode(5)

    solution = Solution()
    result = solution.removeNthFromEnd(head, 2)

    while result:
        print(result.val, end = " ")
        result = result.next


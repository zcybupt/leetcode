class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

def create_list(nums: list) -> ListNode:
    head = ListNode(nums[0])
    tmp = head
    for num in nums[1:]:
        tmp.next = ListNode(num)
        tmp = tmp.next

    return head

def print_list(head: ListNode) -> None:
    while head:
        print(head.val, end = ' ')
        head = head.next
    print()


class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        if not head or not head.next or m >= n: return head

        dummy = ListNode(-1)
        dummy.next = head
        pre = dummy

        for i in range(1, m): pre = pre.next
        if not pre.next: return head

        node_before_m, node_m, cur = pre, pre.next, pre.next
        for i in range(m, n + 1):
            next_node = cur.next
            cur.next = pre
            pre = cur
            cur = next_node

        node_before_m.next = pre
        node_m.next = next_node

        return dummy.next


if __name__ == '__main__':
    head = create_list([1, 2, 3, 4, 5])
    result = Solution().reverseBetween(head, 2, 4)
    print_list(result)


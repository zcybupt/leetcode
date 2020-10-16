class ListNode:
    def __init__(self, x, next = None):
        self.val = x
        self.next = next

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
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        if not l1: return l2
        if not l2: return l1

        dummy = ListNode(-1)
        tmp = dummy
        while l1 and l2:
            if l1.val < l2.val:
                tmp.next = l1
                l1 = l1.next
            else:
                tmp.next = l2
                l2 = l2.next
            tmp = tmp.next

        if l1: tmp.next = l1
        if l2: tmp.next = l2

        return dummy.next


if __name__ == '__main__':
    head1 = create_list([1, 2, 4])
    head2 = create_list([1, 3, 5])
    head = Solution().mergeTwoLists(head1, head2)
    print_list(head)


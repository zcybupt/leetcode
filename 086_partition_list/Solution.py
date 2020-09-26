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
    def partition(self, head: ListNode, x: int) -> ListNode:
        left, right = ListNode(-1), ListNode(-1)

        tmp1, tmp2 = left, right
        while head:
            if head.val < x:
                tmp1.next = head
                head = head.next
                tmp1 = tmp1.next
                tmp1.next = None
            else:
                tmp2.next = head
                head = head.next
                tmp2 = tmp2.next
                tmp2.next = None
        tmp1.next = right.next
        return left.next


if __name__ == '__main__':
    nums = [1, 4, 3, 2, 5, 2]
    head = create_list(nums)
    print_list(head)

    result = Solution().partition(head, 3)
    print_list(result)


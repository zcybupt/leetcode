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


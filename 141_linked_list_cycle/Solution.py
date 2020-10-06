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
    tmp.next = head

    return head

class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        if not head: return False

        slow, fast = head, head
        while fast.next and fast.next.next:
            slow = slow.next
            fast = fast.next.next
            if slow == fast: return True

        return False


if __name__ == '__main__':
    head = create_list([1, 2, 3, 4, 5])
    print(Solution().hasCycle(head))

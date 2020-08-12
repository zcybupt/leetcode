import heapq

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def mergeKLists(self, lists: list) -> ListNode:
        if not lists or len(lists) == 0: return None

        def __lt__(self, other):
            return self.val < other.val

        ListNode.__lt__ = __lt__

        dummy_head = ListNode(0)
        cur_node = dummy_head
        heap = []

        for tmp_list in lists:
            tmp_node = tmp_list
            while tmp_node:
                heapq.heappush(heap, tmp_node)
                tmp_node = tmp_node.next

        while len(heap) != 0:
            cur_node.next = heapq.heappop(heap)
            cur_node = cur_node.next
        cur_node.next = None

        return dummy_head.next


if __name__ == '__main__':
    l1 = ListNode(1)
    l1.next = ListNode(4)
    l1.next.next = ListNode(5)

    l2 = ListNode(1)
    l2.next = ListNode(3)
    l2.next.next = ListNode(4)

    l3 = ListNode(2)
    l3.next = ListNode(6)

    solution = Solution()
    result = solution.mergeKLists([l1, l2, l3])
    while result:
        print(result.val, end=" ")
        result = result.next


class ListNode {
    int val;
    ListNode next;
    ListNode(int val) {
        this.val = val;
    }
}

class Solution {
    public ListNode deleteDuplicates(ListNode head) {
        if (head == null || head.next == null) return head;

        ListNode dummyHead = new ListNode(0);
        ListNode tail = dummyHead;
        for (ListNode l = head, r = head; l != null; l = r) {
            while (r != null && l.val == r.val) r = r.next;
            if (l.next == r) {
                tail.next = l;
                tail = l;
                tail.next = null;
            }
        }

        return dummyHead.next;
    }

    public static void main(String[] args) {
        ListNode head = new ListNode(1);
        head.next = new ListNode(1);
        head.next.next = new ListNode(2);
        head.next.next.next = new ListNode(3);
        head.next.next.next.next = new ListNode(3);

        ListNode result = new Solution().deleteDuplicates(head);

        while (result != null) {
            System.out.println(result.val);
            result = result.next;
        }
    }
}

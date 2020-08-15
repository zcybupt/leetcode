class ListNode {
    int val;
    ListNode next;

    ListNode(int val) {
        this.val = val;
    }
}

class Solution {
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode dummyHead = new ListNode(-1);
        dummyHead.next = head;
        ListNode pre = dummyHead, end = dummyHead;
        ListNode start, next;

        while (end.next != null) {
            for (int i = 0; i < k; i++) {
                end = end.next;
                if (end == null) return dummyHead.next;
            }
            start = pre.next;
            next = end.next; 
            pre.next = reverse(start, next); 
            start.next = next;
            pre = start;
            end = pre;
        }

        return dummyHead.next;
    }

    public ListNode reverse(ListNode head, ListNode tail) {
        ListNode pre = null, next;

        while (head != tail) {
            next = head.next;
            head.next = pre;
            pre = head;
            head = next;
        }

        return pre;
    }

    public static void main(String[] args) {
        ListNode head = new ListNode(1);
        ListNode cur = head;
        for (int i = 2; i < 10; i++) {
            cur.next = new ListNode(i);
            cur = cur.next;
        }

        Solution solution = new Solution();
        cur = solution.reverseKGroup(head, 2);
        while (cur != null) {
            System.out.print(cur.val + " ");
            cur = cur.next;
        }
    }
}

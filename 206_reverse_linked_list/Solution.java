class ListNode {
    int val;
    ListNode next;
    ListNode(int val) {
        this.val = val;
    }
}

class Solution {
    public ListNode reverseList(ListNode head) {
        return reverse(null, head);
    }

    public ListNode reverse(ListNode pre, ListNode cur) {
        if (cur == null) return pre;
        ListNode next = cur.next;
        cur.next = pre;

        return reverse(cur, next);
    }

    public ListNode reverseList2(ListNode head) {
        if (head == null || head.next == null) return head;

        ListNode pre = null;
        ListNode cur = head;
        ListNode next = null;

        while (cur != null) {
            next = cur.next;
            cur.next = pre;
            pre = cur;
            cur = next;
        }

        return pre;
    }

    public static void main(String[] args) {
        ListNode head = new ListNode(0);
        ListNode cur = head;

        for (int i = 1; i < 10; i++) {
            cur.next = new ListNode(i);
            cur = cur.next;
        }

        Solution solution = new Solution();
        cur = solution.reverseList2(head);

        while (cur != null) {
            System.out.print(cur.val + " ");
            cur = cur.next;
        }
    }
}

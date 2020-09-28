class ListNode {
    int val;
    ListNode next;
    ListNode(int val) {
        this.val = val;
    }
}

class Solution {
    public ListNode reverseBetween(ListNode head, int m, int n) {
        if (head == null || m >= n) return head;

        ListNode dummyHead = new ListNode(-1);
        dummyHead.next = head;
        ListNode pre = dummyHead;

        for (int i = 1; i < m; i++) pre = pre.next;
        if (pre.next == null) return head;

        ListNode nodeBeforeM = pre, nodeM = pre.next, cur = pre.next, next = null;
        for (int i = m; i <= n; i++) {
            next = cur.next;
            cur.next = pre;
            pre = cur;
            cur = next;
        }
        nodeBeforeM.next = pre;
        nodeM.next = next;

        return dummyHead.next;
    }

    public static ListNode creatList(int[] nums) {
        ListNode head = new ListNode(nums[0]);
        ListNode tmp = head;
        for (int i = 1; i < nums.length; i++) {
            tmp.next = new ListNode(nums[i]);
            tmp = tmp.next;
        }

        return head;
    }

    public static void printList(ListNode head) {
        while (head != null) {
            System.out.print(head.val + " ");
            head = head.next;
        }
        System.out.println();
    }

    public static void main(String[] args) {
        int[] nums = {1, 2, 3, 4, 5};
        ListNode head = creatList(nums);
        ListNode result = new Solution().reverseBetween(head, 2, 4);
        printList(result);
    }
}

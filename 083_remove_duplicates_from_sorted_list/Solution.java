class ListNode {
    int val;
    ListNode next;
    ListNode(int val) {
        this.val = val;
    }
}

class Solution {
    public ListNode deleteDuplicates(ListNode head) {
        if (head == null) return head;

        ListNode tmp = head;
        while(tmp != null && tmp.next != null) {
            if (tmp.next.val == tmp.val) tmp.next = tmp.next.next;
            else tmp = tmp.next;
        }

        return head;
    }

    public static void main(String[] args) {
        ListNode head = new ListNode(1);
        head.next = new ListNode(1);
        head.next.next = new ListNode(2);
        head.next.next.next = new ListNode(3);
        head.next.next.next.next = new ListNode(3);

        new Solution().deleteDuplicates(head);

        while (head != null) {
            System.out.println(head.val);
            head = head.next;
        }
    }
}

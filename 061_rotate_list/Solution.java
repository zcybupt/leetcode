class ListNode {
    int val;
    ListNode next;

    ListNode(int val) {
        this.val = val;
    }
}

class Solution {
    public ListNode rotateRight(ListNode head, int k) {
        if (head == null) return null;
        ListNode tail = head;
        int length = 1;
        while (tail.next != null) {
            length++;
            tail = tail.next;
        }
        tail.next = head;
        k = length - k % length;
        while (k-- > 0) tail = tail.next;
        head = tail.next;
        tail.next = null;
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
        ListNode head = new ListNode(1);
        ListNode tmp = head;
        for (int i = 2; i <= 5; i++) {
            tmp.next = new ListNode(i);
            tmp = tmp.next;
        }
        Solution solution = new Solution();
        head = solution.rotateRight(head, 11);
        printList(head);
    }
}

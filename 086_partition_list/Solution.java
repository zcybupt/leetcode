class ListNode {
    int val;
    ListNode next;
    ListNode(int val) {
        this.val = val;
    }
}

class Solution {
    public ListNode partition(ListNode head, int x) {
        ListNode left = new ListNode(-1);
        ListNode right = new ListNode(-1);

        ListNode tmp1 = left, tmp2 = right;
        while (head != null) {
            if (head.val < x) {
                tmp1.next = head;
                head = head.next;
                tmp1 = tmp1.next;
                tmp1.next = null;
            } else {
                tmp2.next = head;
                head = head.next;
                tmp2 = tmp2.next;
                tmp2.next = null;
            }
        }
        tmp1.next = right.next;

        return left.next;
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
        int[] nums = {1, 4, 3, 2, 5, 2};
        ListNode head = creatList(nums);
        printList(head);

        ListNode result = new Solution().partition(head, 3);
        printList(result);
    }
}

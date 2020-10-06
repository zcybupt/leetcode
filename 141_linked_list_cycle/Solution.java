class ListNode {
    int val;
    ListNode next;
    ListNode(int val) {
        this.val = val;
    }
}

public class Solution {
    public static ListNode creatList(int[] nums) {
        ListNode head = new ListNode(nums[0]);
        ListNode tmp = head;
        for (int i = 1; i < nums.length; i++) {
            tmp.next = new ListNode(nums[i]);
            tmp = tmp.next;
        }
        tmp.next = head;

        return head;
    }

    public boolean hasCycle(ListNode head) {
        if (head == null) return false;

        ListNode slow = head, fast = head;
        while (fast.next != null && fast.next.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            if (fast == slow) return true;
        }

        return false;
    }

    public static void main(String[] args) {
        ListNode head = creatList(new int[]{1, 2, 3, 4, 5});
        System.out.println(new Solution().hasCycle(head));
    }
}

class ListNode {
    int val;
    ListNode next;
    ListNode (int val) {
        this.val = val;
    }
}

public class Solution {
    public ListNode swapPairs1(ListNode head) {
        if (head == null || head.next == null) return head;

        ListNode tmp, latter = null, cur = head;
        head = head.next;
        
        while (cur != null && cur.next != null) {
            tmp = cur.next;
            cur.next = tmp.next;
            tmp.next = cur;
            if (latter != null) latter.next = tmp;
            latter = cur;
            cur = cur.next;
        }

        return head;
    }

    public ListNode swapPairs2(ListNode head) {
        if (head == null || head.next == null)
            return head;

        ListNode first = head, second = head.next;
        first.next = swapPairs2(second.next);

        second.next = first;

        return second;
    }
    
    public static void main(String[] args) {
        ListNode head = new ListNode(1);
        ListNode cur = head;

        for (int i = 2; i < 10; i++) {
            cur.next = new ListNode(i);
            cur = cur.next;
        }

        Solution solution = new Solution();
        cur = solution.swapPairs2(head);
        
        while (cur != null) {
            System.out.print(cur.val + " ");
            cur = cur.next;
        }
    }
}


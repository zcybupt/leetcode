import java.util.PriorityQueue;

class ListNode {
    int val;
    ListNode next;
    ListNode(int val) {
        this.val = val;
    }
}

public class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0) return null;

        ListNode dummyHead = new ListNode(0);
        ListNode curNode = dummyHead;
        PriorityQueue<ListNode> pq = new PriorityQueue<>((a, b) -> a.val - b.val);

        for (ListNode list : lists) {
            ListNode tmp = list;
            while (tmp != null) {
                pq.add(tmp);
                tmp = tmp.next;
            }
        }

        while (!pq.isEmpty()) {
            curNode.next = pq.poll();
            curNode = curNode.next;
        }
        curNode.next = null; // 防止出现环
        
        return dummyHead.next;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        //ListNode l1 = new ListNode(1);
        //l1.next = new ListNode(4);
        //l1.next.next = new ListNode(5);
        //
        //ListNode l2 = new ListNode(1);
        //l2.next = new ListNode(3);
        //l2.next.next = new ListNode(4);
        //
        //ListNode l3 = new ListNode(2);
        //l3.next = new ListNode(6);
        //ListNode[] lists = new ListNode[]{l1, l2, l3};

        ListNode l1 = new ListNode(-2);
        l1.next = new ListNode(-1);
        l1.next.next = new ListNode(-1);
        l1.next.next.next = new ListNode(-1);
        ListNode[] lists = new ListNode[]{l1, null};

        ListNode head = solution.mergeKLists(lists);
        while (head != null) {
            System.out.print(head.val + " ");
            head = head.next;
        }
    }
}

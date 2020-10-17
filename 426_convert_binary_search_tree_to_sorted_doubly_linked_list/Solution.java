class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val,Node _left,Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
}

class Solution {
    public static Node createTreeUsingBFS(int[] nums) {
        Node[] nodeArr = new Node[nums.length];
        for (int i = 0; i < nums.length; i++) {
            nodeArr[i] = new Node(nums[i]);
            if (i == 0) continue;
            if (i % 2 == 1) nodeArr[(i - 1) / 2].left = nodeArr[i];
            else nodeArr[(i - 1) / 2].right = nodeArr[i];
        }

        return nodeArr[0];
    }

    public static void printTree(Node root) {
        if (root == null) return;
        printTree(root.left);
        System.out.print(root.val + " ");
        printTree(root.right);
    }

    public static void printList(Node head) {
        Node tmp = head;
        do {
            System.out.print(tmp.val + " ");
            tmp = tmp.right;
        } while (tmp != head);
        System.out.println();
    }

    Node pre, head;
    public Node treeToDoublyList(Node root) {
        if (root == null) return null;
        dfs(root);
        head.left = pre;
        pre.right = head;

        return head;
    }

    private void dfs(Node cur) {
        if (cur == null) return;
        dfs(cur.left);
        if (pre != null) pre.right = cur;
        else head = cur;
        cur.left = pre;
        pre = cur;
        dfs(cur.right);
    }

    public static void main(String[] args) {
        Node root = createTreeUsingBFS(new int[]{10, 5, 15, 2, 6, 11, 16, 1, 3});
        printTree(root);
        System.out.println();
        Node head = new Solution().treeToDoublyList(root);
        printList(head);
    }
}

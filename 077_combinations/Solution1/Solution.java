import java.util.*;

class Solution {
    List<List<Integer>> results = new ArrayList<>();

    public List<List<Integer>> combine(int n, int k) {
        if (n == 0 || k == 0 || n < k) return results;

        Deque<Integer> path = new ArrayDeque<>();
        dfs(1, n, k, path);

        return results;
    }

    public void dfs(int start, int n, int k, Deque<Integer> path) {
        if (k == 0) {
            results.add(new ArrayList<>(path));
            return;
        }

        for (int i = start; i <= n; i++) {
            path.addLast(i);
            dfs(i + 1, n, k - 1, path);
            path.removeLast();
        }
    }

    public static void main(String[] args) {
        System.out.println(new Solution().combine(4, 2));
    }
}

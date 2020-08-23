import java.util.*;

class Solution {
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        int len = candidates.length;
        List<List<Integer>> result = new ArrayList<>();
        if (len == 0) return result;

        Deque<Integer> path = new ArrayDeque<>();
        dfs(candidates, 0, len, target, path, result);
        return result;
    }

    public void dfs(int[] candidates, int begin, int len, int target, Deque<Integer> path, List<List<Integer>> result) {
        if (target < 0) return;
        if (target == 0) {
            result.add(new ArrayList<>(path));
            return;
        }

        for (int i = begin; i < len; i++) {
            path.addLast(candidates[i]);
            dfs(candidates, i, len, target - candidates[i], path, result);
            path.removeLast();
        }
    }

    public static void main(String[] args) {
        int[] candidates = new int[]{2, 3, 6, 7};
        int target = 7;
        Solution solution = new Solution();
        List<List<Integer>> results = solution.combinationSum(candidates, target);
        for (List<Integer> result : results) {
            for (int ele : result) {
                System.out.print(ele + " ");
            }
            System.out.println();
        }
    }
}

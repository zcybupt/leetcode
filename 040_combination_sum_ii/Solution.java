import java.util.*;

class Solution {
    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        int len = candidates.length;
        List<List<Integer>> result = new ArrayList<>();
        if (len == 0) return result;

        Arrays.sort(candidates);
        Deque<Integer> path = new ArrayDeque<>();
        dfs(candidates, 0, len, target, path, result);
        return result;
    }

    public void dfs(int[] candidates, int begin, int len, int target, Deque<Integer> path, List<List<Integer>> result) {
        if (target == 0) {
            result.add(new ArrayList<>(path));
            return;
        }

        for (int i = begin; i < len; i++) {
            if (target < candidates[i]) break;
            if (i > begin && candidates[i] == candidates[i - 1]) continue;
            path.addLast(candidates[i]);
            dfs(candidates, i + 1, len, target - candidates[i], path, result);
            path.removeLast();
        }
    }

    public static void main(String[] args) {
        int[] candidates = new int[]{2, 5, 2, 1, 2};
        int target = 5;
        Solution solution = new Solution();
        List<List<Integer>> results = solution.combinationSum2(candidates, target);
        for (List<Integer> result : results) {
            for (int ele : result) {
                System.out.print(ele + " ");
            }
            System.out.println();
        }
    }
}

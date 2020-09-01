import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Solution {
    public List<List<Integer>> permuteUnique(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(nums);
        dfs(res, nums, new ArrayList<>(), new int[nums.length]);
        return res;
    }

    public void dfs(List<List<Integer>> res, int[] nums, List<Integer> tmp, int[] visited) {
        if (tmp.size() == nums.length) {
            res.add(new ArrayList<>(tmp));
            return;
        }

        for (int i = 0; i < nums.length; i++) {
            if (visited[i] == 1) continue;

            if (i > 0 && nums[i] == nums[i - 1] && visited[i - 1] != 1) continue;

            tmp.add(nums[i]);
            visited[i] = 1;
            dfs(res, nums, tmp, visited);
            visited[i] = 0;
            tmp.remove(tmp.size() - 1);
        }
    }

    public static void main(String[] args) {
        int[] nums = new int[]{3, 3, 0, 3};
        Solution solution = new Solution();
        List<List<Integer>> results = solution.permuteUnique(nums);
        for (List<Integer> result : results) {
            for (int ele : result) {
                System.out.print(ele + " ");
            }
            System.out.println();
        }
    }
}


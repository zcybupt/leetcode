import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class Solution {
    public List<List<Integer>> fourSum(int[] nums, int target) {
        int arrLen = nums.length;
        List<List<Integer>> results = new ArrayList<>();
        if (nums == null || arrLen < 4) return results;

        Arrays.sort(nums);
        if (isExceeded(nums, target, 0, arrLen - 1)) return results;

        int L, R;

        for (int i = 0; i < arrLen - 3; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) continue;
            for (int j = i + 1; j < arrLen - 2; j++) {
                if (j > i + 1 && nums[j] == nums[j - 1]) continue;
                L = j + 1;
                R = arrLen - 1;
                if (isExceeded(nums, target, i, R)) return results;

                while (L < R) {
                    int tmpSum = nums[i] + nums[j] + nums[L] + nums[R];
                    if (tmpSum == target) {
                        results.add(new ArrayList<>(Arrays.asList(nums[i], nums[j], nums[L], nums[R])));
                        while (L < R && nums[L] == nums[L + 1]) L++;
                        while (L < R && nums[R] == nums[R - 1]) R--;
                        L++; R--;
                    } else if (tmpSum > target) {
                        R--;
                    } else {
                        L++;
                    }
                }
            }
        }
        return results;
    }

    public boolean isExceeded(int[] nums, int target, int L, int R) {
        int minSum = nums[L] + nums[L + 1] + nums[L + 2] + nums[L + 3];
        int maxSum = nums[R] + nums[R - 1] + nums[R - 2] + nums[R - 3];
        return target > maxSum || target < minSum;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] nums = new int[]{1, 0, -1, 0, -2, 2};
        int target = 0;
        List<List<Integer>> results = solution.fourSum(nums, target);
        for (List<Integer> result : results) {
            for (int ele : result) {
                System.out.print(ele + " ");
            }
            System.out.println();
        }
    }
}

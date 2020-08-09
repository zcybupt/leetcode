import java.util.Arrays;

public class Solution {
    public int threeSumClosest(int[] nums, int target) {
        int arrLen = nums.length;
        if (arrLen < 3) return 0;

        Arrays.sort(nums);
        int L, R;
        int result = 0;
        int minDiff = Integer.MAX_VALUE;
        int curDiff, curSum;

        for (int i = 0; i < arrLen; i++) {
            L = i + 1;
            R = arrLen - 1;
            while (L < R) {
                curSum = nums[i] + nums[L] + nums[R];
                if (curSum == target) return curSum;
                curDiff = Math.abs(curSum - target);
                if (curDiff < minDiff) {
                    result = curSum;
                    minDiff = curDiff;
                }
                if (curSum > target) {
                    R--;
                } else {
                    L++;
                }
            }
        }

        return result;
    }

    public static void main(String[] args) {
        // int[] nums = new int[]{-1, 2, 1, -4};
        // int target = 1;
        int[] nums = new int[]{-1, 0, 1, 1, 55};
        int target = 3;
        Solution solution = new Solution();
        System.out.println(solution.threeSumClosest(nums, target));
    }
}

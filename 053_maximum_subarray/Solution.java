class Solution {
    public int maxSubArray(int[] nums) {
        if (nums == null || nums.length == 0) return 0;

        int current = 0;
        int maxSum = Integer.MIN_VALUE;

        for (int num : nums) {
            current += num;
            maxSum = Math.max(current, maxSum);
            current = Math.max(0, current);
        }

        return maxSum;
    }

    public static void main(String[] args) {
        int[] nums = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
        System.out.println(new Solution().maxSubArray(nums));
    }
}

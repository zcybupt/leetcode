class Solution {
    public int lengthOfLIS(int[] nums) {
        int arrLen = nums.length;
        if (arrLen < 2) return arrLen;

        int[] dp = new int[arrLen];
        dp[0] = 1;
        int result = 1;
        for (int i = 1; i < arrLen; i++) {
            dp[i] = 1;
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j])
                    dp[i] = Math.max(dp[i], dp[j] + 1);
            }
            result = Math.max(result, dp[i]);
        }

        return result;
    }

    public static void main(String[] args) {
        int[] nums = {10, 9, 2, 5, 3, 7, 101, 18};
        System.out.println(new Solution().lengthOfLIS(nums));
    }
}

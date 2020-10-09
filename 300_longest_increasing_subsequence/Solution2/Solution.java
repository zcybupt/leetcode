class Solution {
    public int lengthOfLIS(int[] nums) {
        int arrLen = nums.length;
        if (arrLen < 2) return arrLen;

        int[] tails = new int[arrLen];
        int result = 0;
        for (int num : nums) {
            int l = 0, r = result;
            while (l < r) {
                int mid = l + (r - l) / 2;
                if (tails[mid] < num) l = mid + 1;
                else r = mid;
            }
            tails[l] = num;
            if (result == r) result++;
        }

        return result;
    }

    public static void main(String[] args) {
        int[] nums = {4, 10, 4, 3, 8, 9};
        System.out.println(new Solution().lengthOfLIS(nums));
    }
}

class Solution {
    public boolean canJump(int[] nums) {
        int arrLen = nums.length;
        int rightMost = 0;
        for (int i = 0; i < arrLen; i++) {
            if (i <= rightMost) {
                rightMost = Math.max(rightMost, i + nums[i]);
                if (rightMost >= arrLen - 1) return true;
            }
        }

        return false;
    }

    public boolean canJump2(int[] nums) {
        int arrLen = nums.length;
        int lastPos = arrLen - 1;
        for (int i = arrLen - 2; i >= 0; i--) {
            if (i + nums[i] >= lastPos) lastPos = i;
        }

        return lastPos == 0;
    }

    public static void main(String[] args) {
        int[] nums = {2, 3, 1, 1, 4};
        System.out.println(new Solution().canJump2(nums));
    }
}


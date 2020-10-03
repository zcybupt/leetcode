class Solution {
    public int minArray(int[] nums) {
        int l = 0, r = nums.length - 1;
        while (l < r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] < nums[r]) r = mid;
            else if (nums[mid] > nums[r]) l = mid + 1;
            else r--;
        }

        return nums[l];
    }

    public static void main(String[] args) {
        int[] nums = {3, 4, 5, 1, 2};
        System.out.println(new Solution().minArray(nums));
    }
}

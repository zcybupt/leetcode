class Solution {
    public boolean search(int[] nums, int target) {
        int arrLen = nums.length;

        int l = 0, r = arrLen - 1;
        while (l <= r) {
            while (l < r && nums[l] == nums[l + 1]) l++;
            while (l < r && nums[r] == nums[r - 1]) r--;
            int mid = (l + r) / 2;
            if (nums[mid] == target) return true;
            // nums[l] <= nums[mid] nums[l]等于nums[mid]且有序的情况必在左侧，
            // 否则将违反只旋转一次的规定
            if (nums[mid] < nums[r]) {
                if (nums[mid] < target && target <= nums[r]) l = mid + 1;
                else r = mid - 1;
            } else {
                if (nums[l] <= target && target < nums[mid]) r = mid - 1;
                else l = mid + 1;
            }
        }

        return false;
    }

    public static void main(String[] args) {
        int[] nums = {2, 5, 6, 0, 0, 1, 2};
        System.out.println(new Solution().search(nums, 2));
    }
}

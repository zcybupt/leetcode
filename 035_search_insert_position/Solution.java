class Solution {
    public int searchInsert(int[] nums, int target) {
        if (nums == null || nums.length == 0) return 0;

        int left = 0, right = nums.length - 1;
        while (left <= right) {
            int mid = (left + right) / 2;
            if (nums[mid] == target) return mid;
            if (nums[mid] < target) left = mid + 1;
            else right = mid -1;
        }

        return left;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{1, 3, 5, 6};
        int target = 0;
        Solution solution = new Solution();
        System.out.println(solution.searchInsert(nums, target));
    }
}

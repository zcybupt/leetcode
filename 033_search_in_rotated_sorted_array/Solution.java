public class Solution {
    public int search(int[] nums, int target) {
        int arrLen = nums.length;
        int left = 0, right = arrLen - 1;
        while (left <= right) {
            int mid = (left + right) / 2;
            if (nums[mid] == target) return mid;
            if (nums[mid] < nums[right]) { // 右侧有序
                if (nums[mid] < target && target <= nums[right]) left = mid + 1;
                else right = mid - 1;
            } else { // 左侧有序
                if (nums[left] <= target && target < nums[mid]) right = mid - 1;
                else left = mid + 1;
            }
        }

        return -1;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{4, 5, 6, 7, 0, 1, 2};
        Solution solution = new Solution();
        System.out.println(solution.search(nums, 0));
    }
}

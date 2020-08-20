class Solution {
    public int[] searchRange(int[] nums, int target) {
        if (nums == null || nums.length == 0) return new int[]{-1, -1};
        int arrLen = nums.length;
        int left = 0, right = arrLen - 1;

        while (left <= right) {
            int mid = (left + right) / 2;
            if (nums[mid] == target) {
                while (nums[left] != target) left++;
                while (nums[right] != target) right--;
                return new int[]{left, right};
            }
            
            if (nums[mid] < target) left = mid + 1;
            else right = mid - 1;
        }

        return new int[]{-1, -1};
    }

    public static void main(String[] args) {
        int[] nums = new int[]{5, 7, 7, 8, 8, 10};
        int target = 8;
        Solution solution = new Solution();
        int[] result = solution.searchRange(nums, target);
        System.out.println(result[0] + " " + result[1]);
    }
}


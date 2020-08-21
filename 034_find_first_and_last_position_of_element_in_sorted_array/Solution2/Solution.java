class Solution {
    public int[] searchRange(int[] nums, int target) {
        return new int[] {
            binarySearch(nums, target, true),
            binarySearch(nums, target, false)
        };
    }

    public int binarySearch(int[] nums, int target, boolean isSearchFirst) {
        int arrLen = nums.length;
        if (arrLen == 0) return -1;

        int left = 0, right = arrLen - 1;
        while (left <= right) {
            int mid = (left + right) / 2;
            if (nums[mid] < target) left = mid + 1;
            else if (nums[mid] > target) right = mid - 1;
            else {
                if (isSearchFirst) {
                    if (mid > 0 && nums[mid] == nums[mid - 1]) right = mid - 1;
                    else return mid;
                } else {
                    if (mid < arrLen - 1 && nums[mid] == nums[mid + 1]) left = mid + 1;
                    else return mid;
                }
            }
        }

        return -1;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{5, 7, 7, 8, 8, 10};
        int target = 8;
        Solution solution = new Solution();
        int[] result = solution.searchRange(nums, target);
        System.out.println(result[0] + " " + result[1]);
    }
}

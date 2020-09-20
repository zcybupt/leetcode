class Solution {
    public void sortColors(int[] nums) {
        int L = 0, R = nums.length - 1, cur = 0;

        while (cur <= R) {
            if (nums[cur] == 0) {
                swap(nums, cur, L);
                L++; cur++;
            } else if (nums[cur] == 2) {
                swap(nums, cur, R);
                R--;
            } else {
                cur++;
            }
        }
    }

    public void swap(int[] nums, int i, int j) {
        int tmp = nums[i];
        nums[i] = nums[j];
        nums[j] = tmp;
    }

    public static void main(String[] args) {
        int[] nums = {2, 0, 2, 1, 1, 0};
        new Solution().sortColors(nums);
        for (int num : nums)
            System.out.print(num + " ");
    }
}

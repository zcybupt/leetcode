class Solution {
    public void nextPermutation(int[] nums) {
        int arrLen = nums.length;
        if (arrLen <= 1) return;

        int i = arrLen - 2, j = arrLen - 1, k = arrLen - 1;

        while (i >= 0 && nums[i] >= nums[j]) {
            i--; j--;
        }

        if (i >= 0) {
            while (nums[i] >= nums[k]) k--;
            swap(i, k, nums);
        }

        for (i = j, j = arrLen - 1; i < j; i++, j--) {
            swap(i, j, nums);
        }
    }

    public void swap(int i, int j, int[] nums) {
        int tmp = nums[i];
        nums[i] = nums[j];
        nums[j] = tmp;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{1, 2, 3};
        Solution solution = new Solution();
        solution.nextPermutation(nums);
        for (int num : nums)
            System.out.print(num + " ");
    }
}

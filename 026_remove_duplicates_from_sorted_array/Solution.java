class Solution {
    public int removeDuplicates(int[] nums) {
        if (nums == null || nums.length == 0) return 0;

        int j = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != nums[j])
                nums[++j] = nums[i];
        }

        return j + 1;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{1, 1, 2, 2};
        Solution solution = new Solution();
        int newLen = solution.removeDuplicates(nums);
        for (int i = 0; i < newLen; i++) {
            System.out.print(nums[i] + " ");
        }
    }
}

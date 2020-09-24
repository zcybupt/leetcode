class Solution {
    public int removeDuplicates(int[] nums) {
        if (nums.length <= 2) return nums.length;

        int count = 1, j = 0;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] == nums[j]) count++;
            else count = 1;

            if (count <= 2) nums[++j] = nums[i];
        }

        return j + 1;
    }

    public static void main(String[] args) {
        int[] nums = {1, 2, 2};
        System.out.println(new Solution().removeDuplicates(nums));
        for (int num : nums)
            System.out.print(num + " ");
    }
}

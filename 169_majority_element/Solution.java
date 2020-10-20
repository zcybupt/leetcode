class Solution {
    public int majorityElement(int[] nums) {
        int votes = 0, x = 0;
        for (int num : nums) {
            if (votes == 0) x = num;
            if (x == num) votes++;
            else votes--;
        }

        return x;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{1, 2, 2, 2, 3, 3};
        System.out.println(new Solution().majorityElement(nums));
    }
}

class Solution {
    public int partition(int[] nums, int L, int R) {
        int p = nums[L];
        while (L != R) {
            while (L < R && nums[R] >= p) R--;
            nums[L] = nums[R];
            while (L < R && nums[L] <= p) L++;
            nums[R] = nums[L];
        }
        nums[L] = p;
        return L;
    }

    public int[] quickSort(int[] nums, int L, int R) {
        if (L < R) {
            int pIndex = partition(nums, L, R);
            quickSort(nums, L, pIndex - 1);
            quickSort(nums, pIndex + 1, R);
        }

        return nums;
    }

    public static void main(String[] args) {
        int[] nums = {5, 3, 8, 5, 6, 7, 4, 9, 2};
        new Solution().quickSort(nums, 0, nums.length - 1);
        for (int num : nums)
            System.out.print(num + " ");
    }
}

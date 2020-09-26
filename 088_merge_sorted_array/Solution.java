class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int idx = m-- + n-- - 1;
        while (m >= 0 && n >= 0) {
            nums1[idx--] = nums1[m] > nums2[n] ? nums1[m--] : nums2[n--];
        }

        while (n >= 0) nums1[idx--] = nums2[n--];
    }

    public static void main(String[] args) {
        int[] nums1 = {1, 2, 3, 0, 0, 0};
        int[] nums2 = {2, 5, 6};
        int m = 3, n = 3;

        new Solution().merge(nums1, m, nums2, n);
        for (int num : nums1) {
            System.out.print(num + " ");
        }
    }
}

public class Solution {
    public int maxArea(int[] height) {
        if (height == null || height.length == 0) return 0;

        int arrLen = height.length;
        int l = 0, r = arrLen - 1;
        int ans = 0, area = 0;

        while (l < r) {
            area = Math.min(height[l], height[r]) * (r - l);
            ans = Math.max(ans, area);
            if (height[l] > height[r]) r--;
            else l++;
        }

        return ans;
    }

    public static void main(String[] args) {
        int[] heightArr = new int[]{1, 8, 6, 2, 5, 4, 8, 3, 7};
        Solution solution = new Solution();
        System.out.println(solution.maxArea(heightArr));      
    }
}

class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0) return false;
        int m = matrix.length, n = matrix[0].length;
        int t = 0, b = m - 1, l = 0, r = n - 1, mid1 = 0, mid2 = 0;

        if (target < matrix[0][0] || target > matrix[b][r]) return false;

        for (int i = 0; i < m && t <= b; i++) {
            mid1 = (t + b) / 2;
            if (matrix[mid1][0] > target) b = mid1 - 1;
            else if (matrix[mid1][0] <= target && target <= matrix[mid1][r]) break;
            else t = mid1 + 1;
        }

        for (int i = 0; i < n && l <= r; i++) {
            mid2 = (l + r) / 2;
            if (matrix[mid1][mid2] == target) return true;
            else if (matrix[mid1][mid2] > target) r = mid2 - 1;
            else l = mid2 + 1;
        }

        return false;
    }

    public static void main(String[] args) {
        int[][] matrix = {
            {1,   3,  5,  7},
            {10, 11, 16, 20},
            {23, 30, 34, 50}
        };

        System.out.println(new Solution().searchMatrix(matrix, 10));
    }
}

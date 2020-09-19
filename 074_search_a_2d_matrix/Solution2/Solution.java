class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0) return false;
        int m = matrix.length, n = matrix[0].length;
        int l = 0, r = m * n - 1;
        int mid, midVal;

        while (l <= r) {
            mid = (l + r) / 2;
            midVal = matrix[mid / n][mid % n];
            if (target == midVal) return true;
            else if (target > midVal) l = mid + 1;
            else r = mid - 1;
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

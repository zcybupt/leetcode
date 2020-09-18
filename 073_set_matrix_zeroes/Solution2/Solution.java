class Solution {
    public void setZeroes(int[][] matrix) {
        // 第一行是否含有 0 记录在 matrix[0][0] 中，故无需单独为其设置标记
        // 但需注意填 0 时需倒序，否则可能使矩阵全部为 0
        boolean columnFlag = false;
        int m = matrix.length, n = matrix[0].length;

        for (int i = 0; i < m; i++) {
            if (matrix[i][0] == 0) columnFlag = true;
            for (int j = 1; j < n; j++) {
                if (matrix[i][j] == 0)
                    matrix[i][0] = matrix[0][j] = 0;
            }
        }

        for (int i = m - 1; i >= 0; i--) {
            for (int j = n - 1; j >= 1; j--) {
                if (matrix[i][0] == 0 || matrix[0][j] == 0)
                    matrix[i][j] = 0;
            }
            if (columnFlag) matrix[i][0] = 0;
        }
    }

    public static void main(String[] args) {
        int[][] matrix = {
            {0, 1, 2, 0},
            {3, 4, 5, 2},
            {1, 3, 1, 5}
        };

        new Solution().setZeroes(matrix);

        for (int[] row : matrix) {
            for (int ele : row) {
                System.out.print(ele + " ");
            }
            System.out.println();
        }
    }
}

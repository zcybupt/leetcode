class Solution {
    public int[][] generateMatrix(int n) {
        int[][] result = new int[n][n];
        if (n == 0) return result;
        int row1 = 0, col1 = 0;
        int row2 = n - 1, col2 = n - 1;
        int i = 1;

        while (row1 <= row2 && col1 <= col2)
            i = process(result, i, row1++, col1++, row2--, col2--);

        return result;
    }

    public int process(int[][] result, int i, int row1, int col1, int row2, int col2) {
        if (row1 == row2) {
            for (int j = col1; j <= col2; j++) result[row1][j] = i++;
            return i;
        }
        if (col1 == col2) {
            for (int j = row1; j <= row2; j++) result[j][col1] = i++;
            return i;
        }

        int curR = row1, curC = col1;
        while (curC != col2) result[row1][curC++] = i++;
        while (curR != row2) result[curR++][col2] = i++;
        while (curC != col1) result[row2][curC--] = i++;
        while (curR != row1) result[curR--][col1] = i++;

        return i;
    }

    public int[][] generateMatrix2(int n) {
        int l = 0, r = n - 1, t = 0, b = n - 1;
        int[][] matrix = new int[n][n];
        int num = 1;

        while (num <= n * n) {
            for (int i = l; i <= r; i++) matrix[t][i] = num++;
            t++;
            for (int i = t; i <= b; i++) matrix[i][r] = num++;
            r--;
            for (int i = r; i >= l; i--) matrix[b][i] = num++;
            b--;
            for (int i = b; i >= t; i--) matrix[i][l] = num++;
            l++;
        }

        return matrix;
    }

    public static void main(String[] args) {
        int[][] result = new Solution().generateMatrix2(3);
        for (int[] row : result) {
            for (int ele : row) {
                System.out.print(ele + "\t");
            }
            System.out.println();
        }
    }
}

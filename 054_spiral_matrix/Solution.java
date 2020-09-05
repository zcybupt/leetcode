import java.util.*;

class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> result = new ArrayList<>();
        if (matrix == null || matrix.length == 0) return result;
        int row1 = 0, col1 = 0;
        int row2 = matrix.length - 1, col2 = matrix[0].length - 1;

        while (row1 <= row2 && col1 <= col2)
            result.addAll(getEdge(matrix, row1++, col1++, row2--, col2--));

        return result;
    }

    public List<Integer> getEdge(int[][] m, int row1, int col1, int row2, int col2) {
        List<Integer> tmpResult = new ArrayList<>();

        if (row1 == row2) {
            for (int i = col1; i <= col2; i++)
                tmpResult.add(m[row1][i]);
            return tmpResult;
        }

        if (col1 == col2) {
            for (int i = row1; i <= row2; i++)
                tmpResult.add(m[i][col1]);
            return tmpResult;
        }

        int curR = row1;
        int curC = col1;
        while (curC != col2) tmpResult.add(m[row1][curC++]);
        while (curR != row2) tmpResult.add(m[curR++][col2]);
        while (curC != col1) tmpResult.add(m[row2][curC--]);
        while (curR != row1) tmpResult.add(m[curR--][col1]);

        return tmpResult;
    }

    public static void main(String[] args) {
        int[][] nums = {
            {1, 2, 3},
            {4, 5, 6},
            {7, 8, 9}
        };

        List<Integer> result = new Solution().spiralOrder(nums);
        for (int num : result)
            System.out.print(num + " ");
    }
}


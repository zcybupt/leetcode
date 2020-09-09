import java.util.*;

class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> result = new ArrayList<>();
        if (matrix == null || matrix.length == 0) return result;
        int l = 0, r = matrix[0].length - 1;
        int t = 0, b = matrix.length - 1;

        while (l <= r && t <= b) {
            for (int i = l; i <= r; i++) result.add(matrix[t][i]);
            t++;
            for (int i = t; i <= b; i++) result.add(matrix[i][r]);
            r--;
            for (int i = r; i >= l && t <= b; i--) result.add(matrix[b][i]);
            b--;
            for (int i = b; i >= t && l <= r; i--) result.add(matrix[i][l]);
            l++;
        }

        return result;
    }

    public static void main(String[] args) {
        int[][] nums = {
            {1,  2,  3,  4},
            {5,  6,  7,  8},
            {9, 10, 11, 12}
        };

        List<Integer> result = new Solution().spiralOrder(nums);
        for (int num : result)
            System.out.print(num + " ");
    }
}

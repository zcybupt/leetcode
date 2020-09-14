class Solution {
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        int m = obstacleGrid.length, n = obstacleGrid[0].length;
        int[] dp = new int[n];
        dp[0] = obstacleGrid[0][0] == 1 ? 0 : 1;

        for (int i = 0; i < m; i++) {
            for (int j = 1; j < n; j++) {
                if (obstacleGrid[i][j] == 1) dp[j] = 0;
                else dp[j] = j > 0 ? dp[j] + dp[j - 1] : dp[j];
            }
        }

        return dp[n - 1];
    }

    public static void main(String[] args) {
        int[][] obstacleGrid = {
            {0, 0, 0},
            {0, 1, 0},
            {0, 0, 0}
        };
        System.out.println(new Solution().uniquePathsWithObstacles(obstacleGrid));
    }
}

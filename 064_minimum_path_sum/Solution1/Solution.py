class Solution:
    def minPathSum(self, grid):
        if not grid or len(grid) == 0 or not grid[0] or len(grid[0]) == 0:
            return 0

        m, n = len(grid), len(grid[0])
        dp = [[0] * n for _ in range(m)]

        for i in range(m):
            for j in range(n):
                if i == 0 and j == 0: dp[0][0] = grid[0][0]
                elif i == 0: dp[0][j] = dp[0][j - 1] + grid[0][j]
                elif j == 0: dp[i][0] = dp[i - 1][0] + grid[i][0]
                else: dp[i][j] = grid[i][j] + min(dp[i - 1][j], dp[i][j - 1])

        return dp[-1][-1]


if __name__ == '__main__':
    grid = [
        [1, 3, 1],
        [1, 5, 1],
        [4, 2, 1]
    ]
    print(Solution().minPathSum(grid))


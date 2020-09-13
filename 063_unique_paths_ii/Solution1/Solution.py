class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: list) -> int:
        if not obstacleGrid or not obstacleGrid[0]: return 0
        m, n = len(obstacleGrid), len(obstacleGrid[0])
        dp = [[0] * n for _ in range(m)]

        for i in range(0, m):
            for j in range(0, n):
                if obstacleGrid[i][j] == 1: dp[i][j] = 0
                else:
                    if i == 0 and j == 0: dp[0][0] = 1
                    elif i == 0: dp[0][j] = dp[0][j - 1]
                    elif j == 0: dp[i][0] = dp[i - 1][0]
                    else: dp[i][j] = dp[i - 1][j] + dp[i][j - 1]

        return dp[-1][-1]


if __name__ == '__main__':
    obstacleGrid = [
        [0,0,0],
        [0,1,0],
        [0,0,0]
    ]
    print(Solution().uniquePathsWithObstacles(obstacleGrid))

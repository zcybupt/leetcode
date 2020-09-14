class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: list) -> int:
        m, n = len(obstacleGrid), len(obstacleGrid[0])
        dp = [0] * n
        if obstacleGrid[0][0] == 0: dp[0] = 1

        for i in range(m):
            for j in range(n):
                if obstacleGrid[i][j] == 1: dp[j] = 0
                else: dp[j] = dp[j] + dp[j - 1] if j > 0 else dp[j]

        return dp[-1]


if __name__ == '__main__':
    obstacleGrid = [
        [0,0,0],
        [0,1,0],
        [0,0,0]
    ]
    print(Solution().uniquePathsWithObstacles(obstacleGrid))


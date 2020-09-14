class Solution:
    def minPathSum(self, grid):
        m, n = len(grid), len(grid[0])

        for i in range(m):
            for j in range(n):
                if i == j == 0: continue
                elif i == 0: grid[0][j] += grid[0][j - 1]
                elif j == 0: grid[i][0] += grid[i - 1][0]
                else: grid[i][j] += min(grid[i - 1][j], grid[i][j - 1])

        return grid[-1][-1]


if __name__ == '__main__':
    grid = [
        [1, 3, 1],
        [1, 5, 1],
        [4, 2, 1]
    ]
    print(Solution().minPathSum(grid))

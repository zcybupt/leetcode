class Solution:
    def generateMatrix(self, n: int) -> [[int]]:
        matrix = [[0 for _ in range(n)] for _ in range(n)]
        l, r, t, b = 0, n - 1, 0, n - 1
        num = 1

        while num <= n * n:
            for i in range(l, r + 1):
                matrix[t][i] = num
                num += 1
            t += 1
            for i in range(t, b + 1):
                matrix[i][r] = num
                num += 1
            r -= 1
            for i in range(r, l - 1, -1):
                matrix[b][i] = num
                num += 1
            b -= 1
            for i in range(b, t - 1, -1):
                matrix[i][l] = num
                num += 1
            l += 1

        return matrix


if __name__ == '__main__':
    matrix = Solution().generateMatrix(3)
    for row in matrix:
        print(row)


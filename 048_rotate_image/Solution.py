class Solution:
    def rotate(self, matrix: list) -> None:
        matrix[::] = zip(*matrix[::-1])

    def rotate2(self, matrix: list) -> None:
        n = len(matrix)
        for i in range(n):
            for j in range(i, n):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

        for i in range(n):
            for j in range(0, n // 2):
                tmp = matrix[i][j]
                matrix[i][j] = matrix[i][n - j - 1]
                matrix[i][n - j - 1] = tmp

if __name__ == '__main__':
    matrix = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]
    solution = Solution()
    solution.rotate2(matrix)
    for row in matrix:
        print(row)


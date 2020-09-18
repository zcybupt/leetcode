class Solution:
    def setZeroes(self, matrix: list) -> None:
        col_flag = False
        m, n = len(matrix), len(matrix[0])

        for i in range(m):
            if matrix[i][0] == 0: col_flag = True
            for j in range(1, n):
                if matrix[i][j] == 0:
                    matrix[i][0] = matrix[0][j] = 0

        for i in range(m - 1, -1, -1):
            for j in range(n - 1, 0, -1):
                if (matrix[i][0] == 0 or matrix[0][j] == 0):
                    matrix[i][j] = 0
            if col_flag: matrix[i][0] = 0



if __name__ == '__main__':
    matrix = [
        [0,1,2,0],
        [3,4,5,2],
        [1,3,1,5]
    ]
    Solution().setZeroes(matrix)

    for row in matrix:
        print(row)


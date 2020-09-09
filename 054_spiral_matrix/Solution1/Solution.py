class Solution:
    def spiralOrder(self, matrix: list) -> list:
        result = []
        if not matrix or len(matrix) == 0: return result
        row1, col1 = 0, 0
        row2, col2 = len(matrix) - 1, len(matrix[0]) - 1

        def get_edge(matrix: list) -> list:
            if row1 == row2:
                for i in range(col1, col2 + 1):
                    result.append(matrix[row1][i])
                return

            if col1 == col2:
                for i in range(row1, row2 + 1):
                    result.append(matrix[i][col1])
                return

            for i in range(col1, col2): result.append(matrix[row1][i])
            for i in range(row1, row2): result.append(matrix[i][col2])
            for i in range(col2, col1, -1): result.append(matrix[row2][i])
            for i in range(row2, row1, -1): result.append(matrix[i][col1])

        while (row1 <= row2 and col1 <= col2):
            get_edge(matrix)
            row1, col1, row2, col2 = row1 + 1, col1 + 1, row2 - 1, col2 - 1

        return result


if __name__ == '__main__':
    nums = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]
    print(Solution().spiralOrder(nums))


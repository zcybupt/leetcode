class Solution:
    def searchMatrix(self, matrix: list, target: int) -> bool:
        if not matrix or len(matrix) == 0: return False
        m, n = len(matrix) - 1, len(matrix[0]) - 1
        if target < matrix[0][0] or matrix[m][n] < target: return False

        t, b, mid = 0, m, 0

        while t <= b:
            mid = (t + b) // 2
            if matrix[mid][0] <= target <= matrix[mid][n]: break
            elif matrix[mid][0] > target: b = mid - 1
            else: t = mid + 1

        return target in matrix[mid]


if __name__ == '__main__':
    matrix = [
        [1,   3,  5,  7],
        [10, 11, 16, 20],
        [23, 30, 34, 50]
    ]

    print(Solution().searchMatrix(matrix, 10))


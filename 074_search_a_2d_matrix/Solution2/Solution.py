class Solution:
    def searchMatrix(self, matrix: list, target: int) -> bool:
        if not matrix or len(matrix) == 0 or len(matrix[0]) == 0: return False
        m, n = len(matrix), len(matrix[0])
        l, r = 0, m * n - 1

        while l <= r:
            mid = (l + r) // 2
            mid_val = matrix[mid // n][mid % n]
            if mid_val == target: return True
            elif mid_val < target: l = mid + 1
            else: r = mid - 1

        return False


if __name__ == '__main__':
    matrix = [
        [1,   3,  5,  7],
        [10, 11, 16, 20],
        [23, 30, 34, 50]
    ]

    print(Solution().searchMatrix(matrix, 10))


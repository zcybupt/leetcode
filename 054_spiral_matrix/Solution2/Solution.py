class Solution:
    def spiralOrder(self, matrix: list) -> list:
        result = []
        if not matrix or len(matrix) == 0: return result
        l, r, t, b = 0, len(matrix[0]) - 1, 0, len(matrix) - 1

        while l <= r and t <= b:
            for i in range(l, r + 1): result.append(matrix[t][i])
            t += 1
            for i in range(t, b + 1): result.append(matrix[i][r])
            r -= 1
            if (l <= r and t <= b):
                for i in range(r, l - 1, -1): result.append(matrix[b][i])
                b -= 1
                for i in range(b, t - 1, -1): result.append(matrix[i][l])
                l += 1

        return result


if __name__ == '__main__':
    matrix = [
        [1,  2,  3,  4],
        [5,  6,  7,  8],
        [9, 10, 11, 12]
    ]
    print(Solution().spiralOrder(matrix))

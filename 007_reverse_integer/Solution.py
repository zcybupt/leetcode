class Solution:
    def reverse(self, x: int) -> int:
        limit = 2 ** 31 - 1 if x >= 0 else 2 ** 31
        y, result = abs(x), 0

        while y != 0:
            result = result * 10 + y % 10
            if result > limit: return 0
            y //= 10

        return result if x >= 0 else -result


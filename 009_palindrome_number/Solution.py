class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0: return False

        result = 0
        y = x
        max_int = 2 ** 31 - 1

        while (y != 0):
            result = result * 10 + y % 10
            if result > max_int: return False
            y //= 10

        return x == result


if __name__ == '__main__':
    solution = Solution()
    print(solution.isPalindrome(12321))

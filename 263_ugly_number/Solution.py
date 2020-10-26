class Solution:
    def isUgly(self, num: int) -> bool:
        if not num: return False

        while num % 2 == 0: num //= 2
        while num % 3 == 0: num //= 3
        while num % 5 == 0: num //= 5

        return num == 1


if __name__ == '__main__':
    print(Solution().isUgly(14))


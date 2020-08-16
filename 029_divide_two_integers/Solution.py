class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        res = 0
        sign = -1 if (dividend > 0) ^ (divisor > 0) else 1
        dividend, divisor = abs(dividend), abs(divisor)
        while dividend >= divisor:
            t, p = divisor, 1
            while dividend > (t << 1):
                t <<= 1
                p <<= 1
            dividend -= t
            res += p

        if sign == 1: return min(res, (1 << 31) - 1)
        else: return max(-res, - 1 << 31)


if __name__ == '__main__':
    solution = Solution()
    print(solution.divide(-2147483648, -1))

class Solution:
    def myPow(self, x: float, n: int) -> float:
        def quick_mul(x: float, n: int):
            if n == 0: return 1
            y = quick_mul(x, n // 2)
            return y * y if n % 2 == 0 else y * y * x

        return quick_mul(x, n) if n >= 0 else 1 / quick_mul(x, -n)


if __name__ == '__main__':
    solution = Solution()
    print(solution.myPow(2, 10))


class Solution:
    def myPow(self, x: float, n: int) -> float:
        def quick_mul(x: float, n: int) -> float:
            result = 1
            while n > 0:
                if n % 2 == 1: result *= x
                x *= x
                n //= 2
            return result

        return quick_mul(x, n) if n >= 0 else 1 / quick_mul(x, -n)


if __name__ == '__main__':
    print(Solution().myPow(2, 10))


class Solution:
    def mySqrt(self, x: int) -> int:
        l, r, res = 0, x, -1

        while l <= r:
            mid = (l + r) // 2
            if mid * mid <= x:
                res = mid
                l = mid + 1
            else:
                r = mid - 1

        return res


if __name__ == '__main__':
    print(Solution().mySqrt(5))


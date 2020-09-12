class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        def get_fact(x: int, y: int) -> int:
            if x < y: return 1
            return x * get_fact(x - 1, y)

        return int(get_fact(m + n - 2, n) / get_fact(m - 1, 1))


if __name__ == '__main__':
    print(Solution().uniquePaths(10, 10))


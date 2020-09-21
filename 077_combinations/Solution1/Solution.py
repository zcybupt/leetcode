class Solution:
    def combine(self, n: int, k: int) -> list:
        if n == 0 or k == 0 or n < k: return []

        results = []
        path = []

        def dfs(start: int, n: int, k: int):
            if k == 0:
                results.append(path[:])
                return

            for i in range(start, n + 1):
                path.append(i)
                dfs(i + 1, n, k - 1)
                path.pop()

        dfs(1, n, k)

        return results


if __name__ == '__main__':
    print(Solution().combine(4, 2))


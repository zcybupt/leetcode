class Solution:
    def __init__(self):
        self.facts = [1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880]

    def get_fact(self, n: int) -> int:
        if n < 10: return self.facts[n]
        return n * self.get_fact(n - 1)

    def getPermutation(self, n: int, k: int) -> str:
        result = []
        is_visited = [False] * n
        k -= 1
        for i in range(n):
            tmp_fact = self.get_fact(n - 1 - i)
            cnt = k // tmp_fact
            k %= tmp_fact
            for j in range(n):
                if is_visited[j]: continue
                if cnt == 0:
                    is_visited[j] = True
                    result.append(str(j + 1))
                    break
                cnt -= 1

        return ''.join(result)


if __name__ == '__main__':
    print(Solution().getPermutation(4, 9))


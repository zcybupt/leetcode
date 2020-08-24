class Solution:
    def __init__(self):
        self.path = []
        self.results = []

    def combinationSum(self, candidates: list, target: int) -> list:
        arr_len = len(candidates)
        if arr_len == 0: return self.results

        candidates.sort()
        self.dfs(candidates, 0, arr_len, target)

        return self.results

    def dfs(self, candidates: list, begin: int, arr_len: int, target: int) -> None:
        if target == 0:
            self.results.append(self.path[:]) # 注意此处 path 需要深拷贝
            return

        for i in range(begin, arr_len):
            if target - candidates[i] < 0: break
            self.path.append(candidates[i])
            self.dfs(candidates, i, arr_len, target - candidates[i])
            self.path.pop()


if __name__ == '__main__':
    solution = Solution()
    results = solution.combinationSum([2, 3, 6, 7], 7)
    for result in results:
        print(result)


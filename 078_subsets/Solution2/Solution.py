class Solution:
    def subsets(self, nums: list) -> list:
        results = []
        arr_len = len(nums)

        def dfs(start: int, steps: int, used: list):
            if steps == 0:
                tmp_res = []
                for i in range(0, arr_len):
                    if used[i]: tmp_res.append(nums[i])
                results.append(tmp_res)
                return

            for i in range(start, arr_len):
                used[i] = True
                dfs(i + 1, steps - 1, used)
                used[i] = False

        for i in range(arr_len + 1):
            dfs(0, i, [False] * arr_len)

        return results


if __name__ == '__main__':
    nums = [1, 2, 3]
    results = Solution().subsets(nums)
    for result in results:
        print(result)


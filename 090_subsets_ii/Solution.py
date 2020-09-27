class Solution:
    def subsetsWithDup(self, nums: list) -> list:
        nums.sort()
        arr_len = len(nums)
        results = []
        def dfs(start: int, k: int, used: list):
            if k == 0:
                tmp_res = []
                for i in range(arr_len):
                    if used[i]: tmp_res.append(nums[i])
                results.append(tmp_res)
                return

            for i in range(start, arr_len):
                if i > 0 and not used[i - 1] and nums[i] == nums[i - 1]: continue
                used[i] = True
                dfs(i + 1, k - 1, used)
                used[i] = False

        for i in range(arr_len + 1):
            dfs(0, i, [False] * arr_len)

        return results


if __name__ == '__main__':
    nums = [4, 4, 1, 4]
    results = Solution().subsetsWithDup(nums)
    for result in results:
        print(result)


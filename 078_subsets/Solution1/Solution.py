class Solution:
    def subsets(self, nums: list) -> list:
        results = []
        arr_len = len(nums)

        def dfs(start: int, steps: int, path: list):
            if steps == 0:
                results.append(path[:])
                return

            for i in range(start, arr_len):
                path.append(nums[i])
                dfs(i + 1, steps - 1, path)
                path.pop()

        for i in range(arr_len + 1):
            dfs(0, i, [])

        return results


    def subsets2(self, nums: list) -> list:
        results = []
        arr_len = len(nums)

        def dfs(idx: int, tmp_list: list):
            results.append(tmp_list)
            for i in range(idx, arr_len):
                dfs(i + 1, tmp_list + [nums[i]])

        dfs(0, [])
        return results


if __name__ == '__main__':
    nums = [1, 2, 3]
    print(Solution().subsets(nums))


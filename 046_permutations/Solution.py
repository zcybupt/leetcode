import itertools

class Solution:
    def permute(self, nums: list) -> list:
        results = []
        def dfs(nums, tmp):
            if not nums:
                results.append(tmp)
                return
            for i in range(len(nums)):
                dfs(nums[:i] + nums[i + 1:], tmp + [nums[i]])
        dfs(nums, [])
        return results

    def permute1(self, nums: list) -> list:
        return list(itertools.permutations(nums))


if __name__ == '__main__':
    nums = [1, 2, 3]
    solution = Solution()
    print(solution.permute(nums))


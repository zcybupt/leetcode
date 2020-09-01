class Solution:
    def permuteUnique(self, nums: list) -> list:
        results = []
        arr_len = len(nums)
        def dfs(tmp: list, visited: list):
            if len(tmp) == arr_len:
                results.append(tmp[:])
                return

            for i in range(arr_len):
                if visited[i] or (i > 0 and nums[i] == nums[i - 1] and not visited[i - 1]): continue
                tmp.append(nums[i])
                visited[i] = 1
                dfs(tmp, visited)
                visited[i] = 0
                tmp.pop()

        nums.sort()
        dfs([], [0] * len(nums))
        return results


if __name__ == '__main__':
    nums = [3, 3, 0, 3]
    solution = Solution()
    print(solution.permuteUnique(nums))


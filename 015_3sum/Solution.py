class Solution:
    def threeSum(self, nums: list) -> list:
        list_len = len(nums)
        results = []
        if not nums or list_len < 3: return results

        nums.sort()
        for i in range(list_len):
            if nums[i] > 0: return results
            if (i > 0 and nums[i] == nums[i - 1]): continue
            L = i + 1
            R = list_len - 1
            while L < R:
                if nums[i] + nums[L] + nums[R] == 0:
                    results.append([nums[i], nums[L], nums[R]])
                    while L < R and nums[L] == nums[L + 1]: L += 1
                    while L < R and nums[R] == nums[R - 1]: R -= 1
                    L += 1
                    R -= 1
                elif nums[i] + nums[L] + nums[R] > 0:
                    R -= 1
                else:
                    L += 1
        return results


if __name__ == '__main__':
    nums = [-1, 0, 1, 2, -1, -4]
    solution = Solution()
    print(solution.threeSum(nums))

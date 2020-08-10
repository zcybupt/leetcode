class Solution:
    def fourSum(self, nums: list, target: int) -> list:
        results = []
        arr_len = len(nums)
        if not nums or arr_len < 4: return results

        nums.sort()
        if self.is_exceeded(nums, target, 0, arr_len - 1): return results

        for i in range(arr_len - 3):
            if i > 0 and nums[i] == nums[i - 1]: continue
            for j in range(i + 1, arr_len - 2):
                if (j > i + 1 and nums[j] == nums[j - 1]): continue;
                l = j + 1
                r = arr_len - 1

                while l < r:
                    tmp_sum = nums[i] + nums[j] + nums[l] + nums[r]
                    if tmp_sum == target:
                        results.append([nums[i], nums[j], nums[l], nums[r]])
                        while l < r and nums[l] == nums[l + 1]: l += 1
                        while l < r and nums[r] == nums[r - 1]: r -= 1
                        l += 1
                        r -= 1
                        if self.is_exceeded(nums, target, i, r): continue
                    elif tmp_sum > target: r -= 1
                    else: l += 1

        return results

    def is_exceeded(self, nums: list, target: int, l: int, r: int) -> bool:
        min_sum = nums[l] + nums[l + 1] + nums[l + 2] + nums[l + 3]
        max_sum = nums[r] + nums[r - 1] + nums[r - 2] + nums[r - 3]
        return target > max_sum or target < min_sum


if __name__ == '__main__':
    nums = [1, 0, -1, 0, -2, 2]
    target = 0
    solution = Solution()
    print(solution.fourSum(nums, target))

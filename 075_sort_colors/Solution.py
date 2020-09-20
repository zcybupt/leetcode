class Solution:
    def sortColors(self, nums: list) -> None:
        l, r, cur = 0, len(nums) - 1, 0

        while cur <= r:
            if nums[cur] == 0:
                nums[cur], nums[l] = nums[l], nums[cur]
                cur += 1
                l += 1
            elif nums[cur] == 2:
                nums[cur], nums[r] = nums[r], nums[cur]
                r -= 1
            else:
                cur += 1


if __name__ == '__main__':
    nums = [2, 0, 2, 1, 1, 0]
    Solution().sortColors(nums)
    print(nums)


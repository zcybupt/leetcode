class Solution:
    def removeDuplicates(self, nums: list) -> int:
        if not nums or len(nums) == 0: return 0

        j = 0
        for i in range(len(nums)):
            if nums[i] != nums[j]:
                j += 1
                nums[j] = nums[i]

        return j + 1


if __name__ == '__main__':
    nums = [1, 1, 2, 2]
    solution = Solution()
    new_len = solution.removeDuplicates(nums)
    for i in range(new_len):
        print(nums[i], end = ' ')

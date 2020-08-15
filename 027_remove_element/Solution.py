class Solution:
    def removeElement(self, nums: list, val: int) -> int:
        j = 0
        for i in range(len(nums)):
            if nums[i] != val:
                nums[j] = nums[i]
                j += 1

        return j


if __name__ == '__main__':
    nums = [0, 1, 2, 2, 3, 0, 4, 2]
    val = 2
    solution = Solution()
    new_len = solution.removeElement(nums, val)
    for i in range(new_len):
        print(nums[i], end = ' ')

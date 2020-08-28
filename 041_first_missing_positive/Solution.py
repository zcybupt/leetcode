class Solution:
    def firstMissingPositive(self, nums: list) -> int:
        arr_len = len(nums)

        def swap(i, j, nums):
            tmp = nums[i]
            nums[i] = nums[j]
            nums[j] = tmp

        for i in range(arr_len):
            while 0 < nums[i] <= arr_len and nums[i] != nums[nums[i] - 1]:
                swap(i, nums[i] - 1, nums)

        for i in range(arr_len):
            if nums[i] != i + 1: return i + 1

        return arr_len + 1


if __name__ == '__main__':
    nums = [3, 4, -1, 1]
    solution = Solution()
    print(solution.firstMissingPositive(nums))


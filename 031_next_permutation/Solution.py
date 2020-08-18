class Solution:
    def nextPermutation(self, nums: list) -> None:
        arr_len = len(nums)
        if arr_len <= 1: return

        i, j, k = arr_len - 2, arr_len - 1, arr_len - 1

        while i >= 0 and nums[i] >= nums[j]:
            i -= 1
            j -= 1

        if i >= 0:
            while nums[i] >= nums[k]: k -= 1
            nums[i], nums[k] = nums[k], nums[i]

        i, j = j, arr_len - 1
        while i < j:
            nums[i], nums[j] = nums[j], nums[i]
            i += 1
            j -= 1


if __name__ == '__main__':
    # nums = [1, 2, 3, 4, 5, 6]
    nums = [1, 1]
    solution = Solution()
    solution.nextPermutation(nums)
    print(nums)


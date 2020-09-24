class Solution:
    def removeDuplicates(self, nums: list) -> int:
        arr_len = len(nums)
        if arr_len <= 2: return arr_len

        i, count = 0, 1
        for j in range(1, arr_len):
            if nums[i] == nums[j]: count += 1
            else: count = 1

            if count <= 2:
                i += 1
                nums[i] = nums[j]

        return i + 1


if __name__ == '__main__':
    nums = [1, 1, 1, 2, 2, 3]
    print(Solution().removeDuplicates(nums))
    print(nums)


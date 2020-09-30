class Solution:
    def findDisappearedNumbers(self, nums: list) -> list:
        if not nums: return []

        result = []
        for num in nums: nums[abs(num) - 1] = -abs(nums[abs(num) - 1])
        for i, num in enumerate(nums):
            if num > 0: result.append(i + 1)

        return result


if __name__ == '__main__':
    nums = [4, 3, 2, 7, 8, 2, 3, 1]
    print(Solution().findDisappearedNumbers(nums))


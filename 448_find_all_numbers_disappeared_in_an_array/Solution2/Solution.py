class Solution:
    def findDisappearedNumbers(self, nums: list) -> list:
        if not nums: return []
        for num in nums: nums[abs(num) - 1] = -abs(nums[abs(num) - 1])
        return [i + 1 for i, num in enumerate(nums) if num > 0]


if __name__ == '__main__':
    nums = [4, 3, 2, 7, 8, 2, 3, 1]
    print(Solution().findDisappearedNumbers(nums))

